package simulate

import (
	"admin-gin/global"
	"admin-gin/model/product"
	"admin-gin/process"
	"admin-gin/utils"
	"fmt"
	"time"
)

func (*SimService) SimulateLoadOfGoods() {
	// 目前数据库中有10万个数据, 以这10万个为基础 遍历每个数据 复制100份
	var pageSize = 10
	var page = 1
	for {
		var goods []product.Goods
		offset := (page - 1) * pageSize
		err := global.XTK_DB.
			Order("id desc").
			Where("data_from = 1").
			Limit(pageSize).
			Offset(offset).
			Find(&goods).
			Error
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			continue
		}
		// 拿到了10条真实数据
		for _, g := range goods {
			// 放大100陪
			for i := 0; i < 100; i++ {
				g.ID = 0
				g.IsToEs = 0
				g.CreatedAt = time.Now()
				g = *GoodsBindCoupon(&g)
				g.MediaUID = MakeRandomMediaUID()
				g.Tags = MakeRandomTags()
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
				// 放入队列, 刷入数据库
				process.ToMysqlJobCh <- g
			}
		}
		if len(goods) == 0 {
			break
		}
		page++
		fmt.Println("模拟了1000 条数据")
	}
	// 清理模拟的基础数据
	CleanSimulateBaseData()
}
