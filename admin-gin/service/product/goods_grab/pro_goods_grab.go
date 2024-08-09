package goods_grab

import (
	"admin-gin/global"
	"admin-gin/model/product"
	"admin-gin/model/product/request"
	"admin-gin/process"
	"admin-gin/service/product/goods_grab/internal"
	"admin-gin/service/product/simulate"
	"admin-gin/utils"
	"admin-gin/websocket/common"
	"encoding/json"
	"github.com/google/uuid"
	"strconv"
	"sync"
	"time"
)

type GoodsGrabService struct{}

type CategoryInfo struct {
	ID   uint
	Name string
}

var (
	goNum          int
	workerSignalCh chan struct{}

	pageNum  int
	wg       sync.WaitGroup
	wgMain   sync.WaitGroup
	IsCancel = false

	categoryInfos []CategoryInfo

	UserID uint

	msgListenerCode     int
	msgListenerMsg      string
	msgListenerData     interface{} = nil
	msgListenerToUsers  []uint
	msgListenerForWhere = common.FORWHERE_GRABACTIONLOG
	msgListenerUuid     string

	count int64

	elapsed time.Duration
)

func (*GoodsGrabService) GoodsGrab(uID uint, grabR request.GrabRequest) {
	// 初始化用户ID
	UserID = uID
	msgListenerToUsers = []uint{uID}
	// 初始化携程数量
	goNum = 5
	if grabR.GoNum > 0 {
		goNum = grabR.GoNum
	}
	workerSignalCh = make(chan struct{}, goNum)
	// 初始化分页深度
	pageNum = 100
	if pageNum > 0 {
		pageNum = grabR.PageNum
	}
	// 初始化UUID
	msgListenerUuid = uuid.New().String()
	// 初始化计数
	count = 0

	// 发送开始信息
	msgListenerCode = common.MESSAGE_LEVEL_INFO
	msgListenerMsg = "开始抓取..."
	common.WsSendMsg(msgListenerCode, msgListenerMsg, msgListenerData, msgListenerForWhere, msgListenerToUsers, msgListenerUuid)

	begin := time.Now()
	// 当workerLeader全部结束后， 释放等待
	wgMain.Add(1)
	go workerLeader(grabR.CategoryIDs)

	wgMain.Wait()
	elapsed = time.Since(begin)

	msgListenerCode = common.MESSAGE_LEVEL_INFO
	msgListenerMsg = "采集结束！， 共耗时：" + elapsed.String() + ". 共采集数据" + strconv.FormatInt(count, 10) + "个。"
	common.WsSendMsg(msgListenerCode, msgListenerMsg, msgListenerData, msgListenerForWhere, msgListenerToUsers, msgListenerUuid)

	// 清理模拟的基础数据
	simulate.CleanSimulateBaseData()
	GrabRecordAdd()

}

func GrabRecordAdd() {
	// 将本次采集记录保存到数据库
	var record product.GrabRecord
	record.UserID = UserID
	record.CategoryNum = len(categoryInfos)
	record.GoNum = goNum
	record.PageNum = pageNum
	record.RunTime = elapsed.String()
	record.GrabNum = count
	record.Uuid = msgListenerUuid

	err := global.XTK_DB.Create(&record).Error
	if err != nil {
		global.GVA_LOG.Error(err.Error())
	}

}

func workerLeader(categoryIDs []uint) {
	//defer close(toMysqlJobCh) // 任务全部执行完成后， 关闭发送mysql信息的通道
	defer wgMain.Done()

	query := global.XTK_DB.Model(product.Category{}).Where("level = 3").Select("id, name")
	if len(categoryIDs) > 0 {
		query = query.Where("id IN (?)", categoryIDs)
	}

	err := query.Scan(&categoryInfos).Error
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}

	for _, cate := range categoryInfos {
		if IsCancel {
			break
		}

		workerSignalCh <- struct{}{}
		wg.Add(1)
		go work(cate)
	}
	// 等所以的任务全部执行完成后在return
	wg.Wait()
	return
}

func work(categoryInfo CategoryInfo) {
	defer func() {
		<-workerSignalCh // 释放一个工人
		wg.Done()
	}()

	// 拼接Url
	for i := 1; i < pageNum; i++ {
		url, err := internal.UrlInitByCategoryID(i, categoryInfo.ID)
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			continue
		}

		responseJson, err := internal.SendHttp(url)
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			continue
		}
		// 检查是否已经抓取完
		if responseJson == "[]" {
			msgListenerCode = common.MESSAGE_LEVEL_INFO
			msgListenerMsg = categoryInfo.Name + "采集完成！"
			common.WsSendMsg(msgListenerCode, msgListenerMsg, msgListenerData, msgListenerForWhere, msgListenerToUsers, msgListenerUuid)

			global.GVA_LOG.Info(categoryInfo.Name + "采集完成！")
			return
		}

		// 将json数据转为结构体
		var goodsRes []internal.GoodsReceive
		err = json.Unmarshal([]byte(responseJson), &goodsRes)
		if err != nil {
			if err.Error() == "invalid character '<' looking for beginning of value" {
				global.GVA_LOG.Error("服务器错误,id: " + strconv.Itoa(int(categoryInfo.ID)) + "分类名称：" + categoryInfo.Name)

				msgListenerCode = common.MESSAGE_LEVEL_WARNING
				msgListenerMsg = "服务器错误,id: " + strconv.Itoa(int(categoryInfo.ID)) + "分类名称：" + categoryInfo.Name
				common.WsSendMsg(msgListenerCode, msgListenerMsg, msgListenerData, msgListenerForWhere, msgListenerToUsers, msgListenerUuid)

				return
			}
			global.GVA_LOG.Error(err.Error())
			continue
		}

		for _, re := range goodsRes {
			// 检查是否被取消
			if IsCancel {
				return
			}

			g := re.GoodsReceive2model()
			/// 将分三级分类的信息装入模型
			g.CategoryID = categoryInfo.ID
			// 模拟用户评分
			g.ExperienceScore = utils.RandomFloat64With1Decimal(1.0, 10.0)
			// 计算佣金
			cv := float64(g.CommissionRate) * g.Price / 100
			g.CommissionValue = utils.RoundToOneDecimal(cv)
			// 模拟销量
			g.SalesAll = int32(utils.RandomInt(0, 100000))
			g.SalesMonth = int32(utils.RandomInt(0, 10000))
			g.SalesDay = int32(utils.RandomInt(0, 1000))
			g.Sales2Hour = int32(utils.RandomInt(0, 100))
			// 模拟优惠券
			g = *simulate.GoodsBindCoupon(&g)
			// 模拟团长
			g.MediaUID = simulate.MakeRandomMediaUID()
			// 模拟标签 随机加5个不同的标签
			g.Tags = simulate.MakeRandomTags()
			g.DataFrom = 1
			process.ToMysqlJobCh <- g

			// 由于爬取请求 速度比较慢 所以直接将生成的数据放大100倍
			for i := 0; i < 50; i++ {
				g.ID = 0
				g.CreatedAt = time.Now()
				g = *simulate.GoodsBindCoupon(&g)
				g.MediaUID = simulate.MakeRandomMediaUID()
				g.Tags = simulate.MakeRandomTags()
				g.DataFrom = 2
				// 佣金, 库存
				g.CommissionRate = int32(utils.RandomInt(1, 60))
				g.CommissionValue = utils.RoundToOneDecimal(float64(g.CommissionRate) * g.PostCouponPrice / 100)

				g.Inventory = int32(utils.RandomInt(1000, 60000))

				// 销量, 评分
				g.SalesAll = int32(utils.RandomInt(0, 100000))
				g.SalesMonth = int32(utils.RandomInt(0, 10000))
				g.SalesDay = int32(utils.RandomInt(0, 1000))
				g.Sales2Hour = int32(utils.RandomInt(0, 100))
				g.ExperienceScore = utils.RandomFloat64With1Decimal(1.0, 10.0)
				// 由于查询出来的关联图片 对应的是原理的 商品id 直接往里放的话 复合组件冲突,不会新增新的关联, 所以要将商品id去除, 才好创建新的关联
				process.ToMysqlJobCh <- g

			}

		}
	}
}
