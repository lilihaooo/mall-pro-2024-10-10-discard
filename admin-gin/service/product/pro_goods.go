package product

import (
	"admin-gin/global"
	"admin-gin/kafka/goods"
	"admin-gin/model/product"
	"admin-gin/model/product/request"
	"admin-gin/model/product/response"
	"admin-gin/model/system"
	"admin-gin/utils"
	"admin-gin/utils/captcha"
	"admin-gin/utils/redis"
	"admin-gin/utils/upload"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
	"mime/multipart"
	"strconv"
	"strings"
	"sync"
	"time"
)

type GoodsService struct{}

type categoryAPart struct {
	ID       uint
	ParentID uint
	Name     string
}

var (
	categoryAParts []categoryAPart
	once           sync.Once
)

func InitCategory() {
	fmt.Println("初始化分类信息")
	err := global.XTK_DB.Model(product.Category{}).Select("id, parent_id, name").Scan(&categoryAParts).Error
	if err != nil {
		global.GVA_LOG.Error(err.Error())
	}
}

func collectTags(req request.ProGoodsSearch) (tags []int) {
	if req.IsChoice != 0 {
		tags = append(tags, 36)
	}
	if req.AbsoluteLowPrice != 0 {
		tags = append(tags, 1)
	}
	if req.BuyerReviews != 0 {
		tags = append(tags, 3)
	}
	var pushSucai []string
	if req.PushSuCai != "" {
		pushSucai = strings.Split(req.PushSuCai, ",")
		if len(pushSucai) > 0 {
			for _, one := range pushSucai {
				if one == "1" {
					tags = append(tags, 5)
				}
				if one == "2" {
					tags = append(tags, 6)
				}
				if one == "3" {
					tags = append(tags, 7)
				}
			}
		}

	}

	if req.YunFeiXian != 0 {
		tags = append(tags, 8)
	}
	if req.HuaBeiMX != 0 {
		tags = append(tags, 9)
	}
	if req.PoSunBZ != 0 {
		tags = append(tags, 10)
	}

	var areas []string
	if req.Areas != "" {
		areas = strings.Split(req.Areas, ",")
		if len(areas) > 0 {
			for _, one := range areas {
				if one == "1" {
					tags = append(tags, 12)
				}
				if one == "2" {
					tags = append(tags, 11)
				}
				if one == "3" {
					tags = append(tags, 13)
				}
				if one == "4" {
					tags = append(tags, 16)
				}
				if one == "5" {
					tags = append(tags, 17)
				}
				if one == "6" {
					tags = append(tags, 14)
				}
				if one == "7" {
					tags = append(tags, 15)
				}
			}
		}

	}

	if req.ShopType != 0 {
		switch req.ShopType {
		case 1:
			tags = append(tags, 18)
		case 2:
			tags = append(tags, 19)
		case 3:
			tags = append(tags, 20)
		case 4:
			tags = append(tags, 21)
		case 5:
			tags = append(tags, 22)
		case 6:
			tags = append(tags, 23)
		}
	}
	//
	var ag []string
	if req.ActiveGroup != "" {
		ag = strings.Split(req.ActiveGroup, ",")
		if len(ag) > 0 {
			for _, one := range ag {
				if one == "1" {
					tags = append(tags, 24)
				}
				if one == "2" {
					tags = append(tags, 25)
				}
				if one == "3" {
					tags = append(tags, 26)
				}
				if one == "4" {
					tags = append(tags, 27)
				}
				if one == "5" {
					tags = append(tags, 28)
				}
				if one == "6" {
					tags = append(tags, 29)
				}
				if one == "7" {
					tags = append(tags, 30)
				}
			}
		}
	}

	if req.BiaoXianBD == 1 {
		tags = append(tags, 31)
	}
	if req.BiaoXianBD == 2 {
		tags = append(tags, 32)
	}

	if req.BiaoXianType == 1 {
		tags = append(tags, 33)
	}
	if req.BiaoXianType == 2 {
		tags = append(tags, 34)
	}
	if req.BiaoXianType == 3 {
		tags = append(tags, 35)
	}
	return
}
func (*GoodsService) GoodsSearchByMysql(req request.ProGoodsSearch) (respGoods []response.GoodsSearch, total int64, err error) {
	once.Do(func() {
		InitCategory()
	})

	query := global.XTK_DB.Model(product.Goods{})
	query = query.Model(product.Goods{}).
		Preload("Images").
		Preload("Shop").
		Preload("Coupon").
		Preload("Images").
		Preload("Image").
		Preload("Media").
		Preload("Platform")

	// 条件筛选
	if req.ID != 0 {
		query = query.Where("goods.id=?", req.ID)
		query.Count(&total)
	} else {
		// 分类
		if req.CategoryID != 0 {
			// 递归获取全部最底层子id
			categoryIDs := findLeafCategoryIDs(categoryAParts, req.CategoryID)
			query = query.Where("category_id IN (?)", categoryIDs)
		}

		// 平台
		if req.Platform != 0 {
			query.Joins("join platform pl on pl.id = goods.platform_id")
			query = query.Where("pl.id =?", req.Platform)
		}

		// 品牌库
		if req.IsBrand != 0 {
			//query.Joins("join brands br on br.id = goods.brand_id")
			query.Where("goods.brand_id <>  0")
		}

		// 今日上新
		if req.TodayUp != 0 {
			td := time.Now().Format("2006-01-02")
			startDate := td + " 00:00:00"
			endDate := td + " 23:59:59.999999"
			query.Where("goods.created_at BETWEEN ? AND ?", startDate, endDate)
		}
		// 数据来源
		if req.DataFrom == 1 {
			query = query.Where("data_from =?", 1)
		}
		if req.DataFrom == 2 {
			query = query.Where("data_from =?", 2)
		}
		if req.DataFrom == 3 {
			query = query.Where("data_from =?", 3)
		}

		// 券后价
		if req.PostCouponPriceFrom != 0 {
			query = query.Where("post_coupon_price >= ?", req.PostCouponPriceFrom)
		}
		if req.PostCouponPriceTo != 0 {
			query = query.Where("post_coupon_price  <= ?", req.PostCouponPriceTo)
		}

		if req.CouponValueTo != 0 || req.CouponValueForm != 0 || req.CouponExpirationDay != 0 || req.OrderBy == "qnum" {
			query = query.Joins("join coupons co on co.id = goods.coupon_id")
		}
		// 券过期天数
		if req.CouponExpirationDay != 0 {
			// 获取现在的日期
			p := time.Now().Add(time.Hour * 24 * time.Duration(req.CouponExpirationDay)).Format("2006-01-02 15:04:05")
			query = query.Where("co.end_time > ?", p)
		}
		if req.CouponValueTo != 0 {
			query = query.Where("co.amount <= ?", req.CouponValueTo)
		}
		if req.CouponValueForm != 0 {
			query = query.Where("co.amount >= ?", req.CouponValueForm)
		}
		if req.OrderBy == "qnum" {
			query = query.Order("co.coupon_total DESC")
		}

		// 筛选标签
		tags := collectTags(req) // 收集标签
		if len(tags) > 0 {
			query = query.Joins("join goods_tag gt on gt.goods_id = goods.id").
				Where("gt.tag_id in (?)", tags).
				Group("gt.goods_id").
				Having("COUNT(DISTINCT gt.tag_id) = ?", len(tags))
		}

		// 销量
		if req.SaleNun != 0 {
			query = query.Where("sales_all>=?", req.SaleNun)
		}

		// 佣金率
		if req.CommissionRateFrom != 0 {
			query = query.Where("commission_rate>=?", req.CommissionRateFrom)
		}
		if req.CommissionRateTo != 0 {
			query = query.Where("commission_rate<=?", req.CommissionRateTo)
		}
		// 佣金
		if req.CommissionValue != 0 {
			query = query.Where("commission_value >= ?", req.CommissionValue)
		}

		// 体验分
		if req.ExperienceScore != 0 {
			query = query.Where("experience_score>=?", req.ExperienceScore)
		}

		// 排序
		if req.OrderBy == "cr" {
			query = query.Order("commission_rate DESC")
		}
		if req.OrderBy == "sales_month" {
			query = query.Order("sales_month DESC")
		}
		if req.OrderBy == "sales_day" {
			query = query.Order("sales_day DESC")
		}
		if req.OrderBy == "sales_2_hour" {
			query = query.Order("sales_2_hour DESC")
		}

		if req.OrderBy == "new" {
			query = query.Order("goods.created_at DESC")
		}
		if req.OrderBy == "hot" {
			query = query.Order("push_hot DESC")
		}

		if req.OrderBy == "qhj" {
			if req.OrderType == "ASC" {
				query = query.Order("post_coupon_price ASC")
			} else {
				query = query.Order("post_coupon_price DESC")
			}
		}

		// 筛选条件结束. 我先查询第1000条数据, 如果查到了, 证明数据集超过1000条. 如果没有查到说明数据集没有1000条

		//query.Count(&total)
		total = 1000

		// 分页
		offset := (req.Page - 1) * req.PageSize
		query = query.Limit(req.PageSize).Offset(offset)
	}
	var goods []product.Goods

	//fmt.Println("-------------------------------开始查询")
	err = query.Find(&goods).Error
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, 0, err
	}
	//fmt.Println("-------------------------------查询到了")

	/*序列化*/
	// 将全部分类信息组装为一个map 用于查找long(path)分类名
	categoryMap := make(map[uint]categoryAPart)
	for _, one := range categoryAParts {
		categoryMap[one.ID] = one
	}

	serializeOne := func(model *product.Goods) response.GoodsSearch {
		p := response.GoodsSearch{}
		err = copier.Copy(&p, model)
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			return p
		}
		// 不能复制的手动赋值
		p.ShopName = model.Shop.Name
		p.PlatformName = model.Platform.Name
		p.MediaName = model.Media.MediaName
		p.CouponID = model.CouponID
		p.CouponAmount = model.Coupon.Amount
		p.CouponEndTime = model.Coupon.EndTime.Format("2006-01-02 15:04:05")
		p.CouponCover = model.Coupon.CouponCover
		p.CouponTotal = model.Coupon.CouponTotal
		p.PushHot = model.PushHot

		// 返回封面图片   ---  默认图片集合第一张为封面, 如果设置了封面则替换为设置的
		coverImg := model.Image.Url
		if coverImg == "" {
			coverImg = model.Images[0].Url
		}
		p.CoverImage = coverImg
		// 获取分类信息
		p.CategoryName = findLongCategoryName(model.CategoryID, categoryMap)
		return p
	}

	for _, model := range goods {
		respGoods = append(respGoods, serializeOne(&model))
	}
	return
}

// 查找某个分类 ID 的所有最底层子分类 ID
func findLeafCategoryIDs(categories []categoryAPart, targetID uint) []uint {
	// 建立一个父子关系的映射
	childMap := make(map[uint][]uint)
	for _, category := range categories {
		childMap[category.ParentID] = append(childMap[category.ParentID], category.ID)
	}

	var leafIDs []uint
	// 辅助函数，用于递归查找叶子节点
	var findLeaves func(uint)
	findLeaves = func(id uint) {
		if children, exists := childMap[id]; exists {
			for _, childID := range children {
				findLeaves(childID)
			}
		} else {
			leafIDs = append(leafIDs, id)
		}
	}

	// 从目标 ID 开始查找
	findLeaves(targetID)
	return leafIDs
}

// 通过一个id找到该分类的路径
func findLongCategoryName(targetID uint, m map[uint]categoryAPart) (res string) {
	var findParentName func(uint)
	findParentName = func(id uint) {
		if self, exists := m[id]; exists {
			res = "/" + self.Name + res
			if self.ParentID == 0 {
				return
			}
			findParentName(self.ParentID)
		} else {
			return
		}
	}
	findParentName(targetID)
	return
}

func (*GoodsService) GoodsInfo(req request.GoodsInfo) (goods product.Goods, err error) {
	err = global.XTK_DB.
		Where("id=?", req.ID).
		Preload("Shop", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name, logo, logistics_score, product_score, service_score") // 这里选择 Shop 表中的 id 和 name 字段
		}).
		Preload("Media").
		Preload("Images").
		Preload("Image").
		Preload("Tags").
		Preload("Brand").
		Preload("Coupon").
		Find(&goods).
		Error

	// 再写一条sql SELECT * FROM `goods` WHERE `goods`.`id` = 1213;  作为对比
	global.XTK_DB.Find(&goods, req.ID)

	return
}

// 抓取记录
func (*GoodsService) GrabRecord(req request.GrabRecord) (list []response.ProGrabRecord, total int64, err error) {
	query := global.XTK_DB
	err = query.Model(product.GrabRecord{}).Count(&total).Error
	if err != nil {
		return
	}

	var gr []product.GrabRecord
	offset := req.PageSize * (req.Page - 1)
	err = query.Order("created_at desc").Limit(req.PageSize).Offset(offset).Find(&gr).Error
	if err != nil {
		return
	}

	// 获取用户数据
	var Users []system.SysUser
	err = global.GVA_DB.Select("id, username").Find(&Users).Error
	if err != nil {
		return
	}
	userIDNameMap := make(map[uint]string)
	for _, user := range Users {
		userIDNameMap[user.ID] = user.Username
	}
	// 序列化数据
	for _, record := range gr {
		list = append(list, response.ProGrabRecord{
			CreatedTime: record.GVA_MODEL.CreatedAt.Format("2006-01-02 15:04:05"),
			UserName:    userIDNameMap[record.UserID],
			CategoryNum: record.CategoryNum,
			GoNum:       record.GoNum,
			PageNum:     record.PageNum,
			RunTime:     record.RunTime,
			GrabNum:     record.GrabNum,
			Uuid:        record.Uuid,
		})
	}
	return
}

// 收集标签
func collectV2Tags(req request.ProGoodsSearchV2) (tags []int) {
	if req.IsChoice != 0 {
		tags = append(tags, 36)
	}
	if req.AbsoluteLowPrice != 0 {
		tags = append(tags, 1)
	}
	if req.BuyerReviews != 0 {
		tags = append(tags, 3)
	}
	var pushSucai []string
	if req.PushSuCai != "" {
		pushSucai = strings.Split(req.PushSuCai, ",")
		if len(pushSucai) > 0 {
			for _, one := range pushSucai {
				if one == "1" {
					tags = append(tags, 5)
				}
				if one == "2" {
					tags = append(tags, 6)
				}
				if one == "3" {
					tags = append(tags, 7)
				}
			}
		}

	}

	if req.YunFeiXian != 0 {
		tags = append(tags, 8)
	}
	if req.HuaBeiMX != 0 {
		tags = append(tags, 9)
	}
	if req.PoSunBZ != 0 {
		tags = append(tags, 10)
	}

	var areas []string
	if req.Areas != "" {
		areas = strings.Split(req.Areas, ",")
		if len(areas) > 0 {
			for _, one := range areas {
				if one == "1" {
					tags = append(tags, 12)
				}
				if one == "2" {
					tags = append(tags, 11)
				}
				if one == "3" {
					tags = append(tags, 13)
				}
				if one == "4" {
					tags = append(tags, 16)
				}
				if one == "5" {
					tags = append(tags, 17)
				}
				if one == "6" {
					tags = append(tags, 14)
				}
				if one == "7" {
					tags = append(tags, 15)
				}
			}
		}

	}

	if req.ShopType != 0 {
		switch req.ShopType {
		case 1:
			tags = append(tags, 18)
		case 2:
			tags = append(tags, 19)
		case 3:
			tags = append(tags, 20)
		case 4:
			tags = append(tags, 21)
		case 5:
			tags = append(tags, 22)
		case 6:
			tags = append(tags, 23)
		}
	}
	//
	var ag []string
	if req.ActiveGroup != "" {
		ag = strings.Split(req.ActiveGroup, ",")
		if len(ag) > 0 {
			for _, one := range ag {
				if one == "1" {
					tags = append(tags, 24)
				}
				if one == "2" {
					tags = append(tags, 25)
				}
				if one == "3" {
					tags = append(tags, 26)
				}
				if one == "4" {
					tags = append(tags, 27)
				}
				if one == "5" {
					tags = append(tags, 28)
				}
				if one == "6" {
					tags = append(tags, 29)
				}
				if one == "7" {
					tags = append(tags, 30)
				}
			}
		}
	}

	if req.BiaoXianBD == 1 {
		tags = append(tags, 31)
	}
	if req.BiaoXianBD == 2 {
		tags = append(tags, 32)
	}

	if req.BiaoXianType == 1 {
		tags = append(tags, 33)
	}
	if req.BiaoXianType == 2 {
		tags = append(tags, 34)
	}
	if req.BiaoXianType == 3 {
		tags = append(tags, 35)
	}
	return
}
func (*GoodsService) GoodsSearchByEs(req request.ProGoodsSearchV2, userID uint) (respGoods []response.GoodsSearchV2, total int64, err error) {
	respGoods = []response.GoodsSearchV2{}
	// 获取收藏set
	myCollectMap := make(map[string]struct{})
	// 查看是否存在该缓存
	key := "collect:" + strconv.Itoa(int(userID))
	exists, err := global.GVA_REDIS.Exists(context.Background(), key).Result()
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}
	// 如果缓存不存在
	if exists == 0 {
		// 不存在就查数据库并保存到缓存
		var myCollectGoodsIDs []string
		err = global.XTK_DB.Model(product.Collect{}).Where("user_id = ?", userID).Pluck("goods_id", &myCollectGoodsIDs).Error
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			return
		}
		// 查询后无论有没有记录 都加入缓存中
		if len(myCollectGoodsIDs) > 0 {
			_, err = global.GVA_REDIS.SAdd(context.Background(), key, myCollectGoodsIDs).Result()
			if err != nil {
				global.GVA_LOG.Error(err.Error())
				return
			}
		}

		// 设置过期时间，例如设置为 60 秒
		expiration := 60 * time.Second
		err = global.GVA_REDIS.Expire(context.Background(), key, expiration).Err()
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			return
		}

		if len(myCollectGoodsIDs) > 0 {
			for _, id := range myCollectGoodsIDs {
				myCollectMap[id] = struct{}{}
			}
		}
	} else {
		members, err := global.GVA_REDIS.SMembers(context.Background(), key).Result()
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			return respGoods, 0, nil
		}
		if len(members) > 0 {
			for _, member := range members {
				myCollectMap[member] = struct{}{}
			}
		}
	}

	once.Do(func() {
		InitCategory()
	})
	// 构建 复合 查询条件
	boolQuery := elastic.NewBoolQuery()

	// 关键词搜索 - 模糊匹配
	if strings.TrimSpace(req.Keyword) != "" {
		boolQuery.Must(elastic.NewMatchQuery("all", req.Keyword))
		// 处理关键词
		manageKeyword(req.Keyword)
	}
	// 分类查询
	if req.CategoryID != 0 {
		categoryIDs := findLeafCategoryIDs(categoryAParts, req.CategoryID)
		if len(categoryIDs) > 0 {
			cI := make([]interface{}, len(categoryIDs))
			for i, one := range categoryIDs {
				cI[i] = one
			}
			boolQuery.Filter(elastic.NewTermsQuery("category_id", cI...)) // NewTermsQuery 接口数组  而且还必须展开,  NewTermQuery int
		}
	}
	// 筛选
	// 券后价
	if req.PostCouponPriceFrom != 0 {
		boolQuery.Filter(elastic.NewRangeQuery("post_coupon_price").Gte(req.PostCouponPriceFrom))
	}
	if req.PostCouponPriceTo != 0 {
		boolQuery.Filter(elastic.NewRangeQuery("post_coupon_price").Lte(req.PostCouponPriceTo))
	}
	// 佣金比例
	if req.CommissionRateFrom != 0 {
		boolQuery.Filter(elastic.NewRangeQuery("commission_rate").Gte(req.CommissionRateFrom))
	}
	if req.CommissionRateTo != 0 {
		boolQuery.Filter(elastic.NewRangeQuery("commission_rate").Lte(req.CommissionRateTo))
	}
	// 佣金
	if req.CommissionValue != 0 {
		boolQuery.Filter(elastic.NewRangeQuery("commission_value").Gte(req.CommissionValue))
	}
	// 销量
	if req.SaleNun != 0 {
		boolQuery.Filter(elastic.NewRangeQuery("sales_all").Gte(req.SaleNun))
	}
	// 券面额
	if req.CouponValueForm != 0 {
		boolQuery.Filter(elastic.NewRangeQuery("coupon_amount").Gte(req.CouponValueForm))
	}
	if req.CouponValueTo != 0 {
		boolQuery.Filter(elastic.NewRangeQuery("coupon_amount").Lte(req.CouponValueTo))
	}
	// 到期时间
	if req.CouponExpirationDay != 0 {
		p := time.Now().Add(time.Hour * 24 * time.Duration(req.CouponExpirationDay)).Format("2006-01-02 15:04:05")
		boolQuery.Filter(elastic.NewRangeQuery("coupon_end_time").Gte(p))
	}
	// 体验分
	if req.ExperienceScore != 0 {
		boolQuery.Filter(elastic.NewRangeQuery("experience_score").Gte(req.CouponValueForm))
	}
	// 今日上新
	if req.TodayUp != 0 {
		td := time.Now().Format("2006-01-02")
		startDate := td + " 00:00:00"
		endDate := td + " 23:59:59"
		//endDate := td + " 23:59:59.999999"
		// 中国时间
		loc, _ := time.LoadLocation("Asia/Shanghai") // 加载中国时区
		st, _ := time.ParseInLocation("2006-01-02 15:04:05", startDate, loc)
		et, _ := time.ParseInLocation("2006-01-02 15:04:05", endDate, loc)
		boolQuery.Filter(elastic.NewRangeQuery("created_at").Gte(st.Format("2006-01-02 15:04:05")))
		boolQuery.Filter(elastic.NewRangeQuery("created_at").Lte(et.Format("2006-01-02 15:04:05")))
	}

	// 标签搜索
	tags := collectV2Tags(req)
	if len(tags) > 0 {
		for _, one := range tags {
			boolQuery.Filter(elastic.NewTermQuery("tags", one))
		}

	}
	// 数据来源
	if req.DataFrom == 1 {
		boolQuery.Filter(elastic.NewTermQuery("tags", req.DataFrom))
	}
	if req.DataFrom == 2 {
		boolQuery.Filter(elastic.NewTermQuery("tags", req.DataFrom))
	}
	if req.DataFrom == 3 {
		boolQuery.Filter(elastic.NewTermQuery("tags", req.DataFrom))
	}
	// 品牌库   查看品牌id不为0的商品
	if req.IsBrand == 1 {
		boolQuery.MustNot(elastic.NewTermQuery("brand_id", 0))
	}
	// 平台
	if req.Platform == 1 {
		boolQuery.Filter(elastic.NewTermQuery("platform_id", req.Platform))
	}
	if req.Platform == 2 {
		boolQuery.Filter(elastic.NewTermQuery("platform_id", req.Platform))
	}
	if req.Platform == 3 {
		boolQuery.Filter(elastic.NewTermQuery("platform_id", req.Platform))
	}

	query := global.ESClient.Search().Index("goods_index").Query(boolQuery)

	// 查询出总数量
	total, err = global.ESClient.Count("goods_index").Query(boolQuery).Do(context.Background())
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return respGoods, 0, err
	}
	// 设置了最大返回数量
	if total > 1000 {
		total = 1000
	}

	// 券后价排序
	if req.OrderBy == "qhj" {
		if req.OrderType == "DESC" {
			query = query.Sort("post_coupon_price", false)
		} else {
			query = query.Sort("post_coupon_price", true)
		}
	}
	// 今日热销排序
	if req.OrderBy == "sales_day" {
		query = query.Sort("sales_day", false)
	}
	// 2小时热销排序
	if req.OrderBy == "sales_2_hour" {
		query = query.Sort("sales_2_hour", false)
	}
	// 月销量排序
	if req.OrderBy == "sales_month" {
		query = query.Sort("sales_month", false)
	}
	// 佣金比例排序
	if req.OrderBy == "cr" {
		query = query.Sort("commission_rate", false)
	}
	// 最新上架排序
	if req.OrderBy == "new" {
		query = query.Sort("created_at", false)
	}
	// 推广热度排序
	if req.OrderBy == "hot" {
		query = query.Sort("push_hot", false)
	}
	// 领券量
	if req.OrderBy == "qnum" {
		query = query.Sort("coupon_total", false)
	}

	offset := req.PageSize * (req.Page - 1)
	query = query.From(offset).Size(req.PageSize)
	// 执行搜索请求
	esRes, err := query.Do(context.Background())
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return respGoods, 0, err
	}

	now := time.Now()
	location, err := time.LoadLocation("Asia/Shanghai")
	layout := "2006-01-02 15:04:05"
	if err != nil {
		return respGoods, 0, err
	}

	// 解析结果
	for _, hit := range esRes.Hits.Hits {
		var g product.EsGoods
		err = json.Unmarshal(hit.Source, &g)
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			return respGoods, 0, err
		}

		var gResp response.GoodsSearchV2
		err = copier.Copy(&gResp, g)
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			return respGoods, 0, err
		}
		// 处理时间
		btime, err := time.ParseInLocation(layout, g.CouponBeginTime, location)
		if err != nil {
			return respGoods, 0, err
		}
		etime, err := time.ParseInLocation(layout, g.CouponEndTime, location)
		if err != nil {
			return respGoods, 0, err
		}
		if now.After(btime) && now.Before(etime) {
			gResp.IsExpire = false
		} else {
			gResp.IsExpire = true
		}
		if g.CouponID == 0 {
			gResp.IsExpire = false
		}
		// 查看是否收藏
		mapKey := strconv.Itoa(int(g.ID))
		if _, ex := myCollectMap[mapKey]; ex {
			gResp.IsCollect = true
		}
		respGoods = append(respGoods, gResp)
	}
	return
}

// 处理关键词
func manageKeyword(keyword string) {
	// 将关键词保存到热搜中
	zset := redis.NewResZSet()
	// 检查这个关键词的得分
	score, err := zset.GetMemberScore(context.Background(), "HotSearch", keyword)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
	}

	if score > 10 {
		c := captcha.NewDefaultRedisStore()
		searchKeyword := c.Get(c.PreKey+keyword, false)
		if searchKeyword == "" {
			// 发送kafka消息, 添加这个词语到检索库中
			sender := goods.NewCreateEsSearchKeywordSender(keyword)
			err = sender.Send()
			if err != nil {
				global.GVA_LOG.Error(err.Error())
			}
			// 将该关键字保存到缓存中防止高并发发送信息, 频繁操作es
			err = c.Set(keyword, "1")
			if err != nil {
				global.GVA_LOG.Error(err.Error())
			}
		}
	}

	_, err = zset.IncrementMemberScore(context.Background(), "HotSearch", keyword, 1)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
	}
	// 只保留前100名
	err = global.GVA_REDIS.ZRemRangeByRank(context.Background(), "HotSearch", 100, -1).Err()
	if err != nil {
		global.GVA_LOG.Error(err.Error())
	}
}

// 商品图片上传
func (*GoodsService) GoodsUploadImage(file *multipart.FileHeader, goodsID uint) error {
	oss := upload.NewOss()
	filePath, _, uploadErr := oss.UploadFile(file)
	if uploadErr != nil {
		return uploadErr
	}
	// 图片上传成功 将图片地址保存到数据库
	err := global.XTK_DB.Transaction(func(tx *gorm.DB) error {
		var i product.Image
		i.Url = filePath
		if err := global.XTK_DB.Create(&i).Error; err != nil {
			return err
		}
		var gi product.GoodsImage
		gi.ImageID = i.ID
		gi.GoodsID = goodsID
		if err := global.XTK_DB.Create(&gi).Error; err != nil {
			return err
		}
		return nil
	})

	return err
}

// 删除商品图片
func (*GoodsService) DeleteGoodsImage(id uint) error {
	// 获取图片信息
	var i product.Image
	err := global.XTK_DB.Where("id = ?", id).First(&i).Error
	if err != nil {
		return err
	}

	err = global.XTK_DB.Transaction(func(tx *gorm.DB) error {
		// 查询出哪些商品该图片作为封面图
		var gs []product.Goods
		err = tx.Where("cover_image_id = ?", i.ID).Preload("Images").Find(&gs).Error
		if err != nil {
			return err
		}
		if len(gs) != 0 {
			// 更改商品封面应用为null
			err = tx.Model(product.Goods{}).Where("cover_image_id = ?", i.ID).Update("cover_image_id", nil).Error
			if err != nil {
				return err
			}
			for _, g := range gs {
				ci := ""
				if len(g.Images) > 0 {
					ci = g.Images[0].Url
				}
				m := goods.Message{
					ID:         g.ID,
					CoverImage: ci,
					CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
				}
				// 发送消息通知es
				sender := goods.NewUpdateEsGoodsInfoSender(m)
				err = sender.Send()
				if err != nil {
					return err
				}
			}
		}

		// 删除图片表
		err = tx.Where("id = ?", id).Delete(&product.Image{}).Error
		if err != nil {
			return err
		}
		// 删除关联表
		err = tx.Where("image_id = ?", id).Delete(&product.GoodsImage{}).Error
		if err != nil {
			return err
		}
		// 删除文件
		// 根据图片url来判断判读删除位置
		var oss upload.OSS
		parts := strings.Split(i.Url, "/")
		host := parts[2]
		var key string
		if host == "127.0.0.1:8887" {
			key = parts[len(parts)-1]
			oss = &upload.Local{}
		}
		// 判断oss是否已经被实例化
		if oss == nil {
			return errors.New("oss类型未被实例化")
		}
		if err = oss.DeleteFile(key); err != nil {
			return errors.New("文件删除失败")
		}
		return nil
	})
	return err
}

// 设置商品的封面图片
func (*GoodsService) SetCoverImage(goodsID uint, imageID uint) error {
	// 获得图片信息
	var i product.Image
	err := global.XTK_DB.First(&i, imageID).Error
	if err != nil {
		return err
	}
	err = global.XTK_DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Model(product.Goods{}).Where("id = ?", goodsID).Update("cover_image_id", imageID).Error
		if err != nil {
			return err
		}

		m := goods.Message{
			ID:         goodsID,
			CoverImage: i.Url,
			CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
		}
		sender := goods.NewUpdateEsGoodsInfoSender(m)
		return sender.Send()
	})
	return err
}

// 基本信息修改
func (*GoodsService) UpdateBaseInfo(baseInfo request.UpdateGoodsBaseInfo) error {
	// 查询要修改的商品对象
	var g product.Goods
	err := global.XTK_DB.Model(&g).Select("id, coupon_id").Preload("Coupon", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, amount")
	}).Where("id = ?", baseInfo.GoodsID).First(&g).Error
	if err != nil {
		return err
	}

	if baseInfo.Price < g.Coupon.Amount {
		return errors.New("价格设置低于优惠券金额")
	}
	// 券后价
	postCouponPrice := utils.RoundToOneDecimal(baseInfo.Price - g.Coupon.Amount)
	commissionValue := utils.ComputeCommissionValue(postCouponPrice, baseInfo.CommissionRate)

	// 添加时间
	createAt := time.Now().Format("2006-01-02 15:04:05")
	// 更改商品基础信息
	updateMap := make(map[string]interface{}, 6)
	updateMap["description"] = baseInfo.Description
	updateMap["price"] = baseInfo.Price
	updateMap["commission_rate"] = baseInfo.CommissionRate
	updateMap["commission_value"] = commissionValue
	updateMap["post_coupon_price"] = postCouponPrice
	updateMap["created_at"] = createAt

	// 获取标签信息
	var ts []product.Tag
	if len(baseInfo.Tags) > 0 {
		err = global.XTK_DB.Select("id").Where("id in ?", baseInfo.Tags).Find(&ts).Error
		if err != nil {
			return err
		}
	}

	err = global.XTK_DB.Transaction(func(tx *gorm.DB) error {
		// 更改商品表信息
		err = tx.Model(product.Goods{}).Where("id = ?", baseInfo.GoodsID).Updates(updateMap).Error
		if err != nil {
			return err
		}
		// 更改商品-标签表关系关系
		err = tx.Model(&g).Association("Tags").Replace(ts)
		if err != nil {
			return err
		}
		// 发送消息
		m := goods.Message{
			ID:              baseInfo.GoodsID,
			Description:     baseInfo.Description,
			Price:           baseInfo.Price,
			CommissionRate:  baseInfo.CommissionRate,
			CommissionValue: commissionValue,
			Tags:            baseInfo.Tags,
			CreatedAt:       createAt,
			PostCouponPrice: postCouponPrice,
		}

		sender := goods.NewUpdateEsGoodsInfoSender(m)
		return sender.Send()
	})
	return err
}

// 优惠券信息修改
func (*GoodsService) UpdateCouponInfo(couponInfo request.UpdateGoodsCouponInfo) error {
	// 中国时区
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return err
	}

	// 转换为中国时间
	beginDTimeStr := couponInfo.BeginTime.In(location).Format("2006-01-02") + " 00:00:00"
	endDTimeStr := couponInfo.EndTime.In(location).Format("2006-01-02") + " 23:59:59"

	// 将字符串解析为时间格式一定要指定时区
	dayE, err := time.ParseInLocation("2006-01-02 15:04:05", endDTimeStr, location) // 中国时间
	if err != nil {
		return err
	}
	dayB, err := time.ParseInLocation("2006-01-02 15:04:05", beginDTimeStr, location) // 中国时间
	if err != nil {
		return err
	}

	// mysql 新增优惠券
	c := product.Coupon{
		Title:       couponInfo.Title,
		Amount:      couponInfo.Amount,
		MinAmount:   couponInfo.MinAmount,
		BeginTime:   dayB,
		EndTime:     dayE,
		CouponTotal: couponInfo.CouponTotal,
		CouponCover: couponInfo.CouponCover,
		Status:      1,
	}
	// 获取商品原价
	var g product.Goods
	err = global.XTK_DB.Select("price, commission_rate").Where("id = ?", couponInfo.GoodsID).Take(&g).Error
	if err != nil {
		return err
	}
	// 更改商品优惠券id和券后价
	var postCouponPrice float64
	if g.Price > couponInfo.Amount {
		postCouponPrice = utils.RoundToOneDecimal(g.Price - couponInfo.Amount)
	} else {
		return errors.New("优惠券额度大于商品价格")
	}

	// 佣金值
	commissionValue := utils.ComputeCommissionValue(postCouponPrice, g.CommissionRate)

	createdAt := time.Now().Format("2006-01-02 15:04:05")

	return global.XTK_DB.Transaction(func(tx *gorm.DB) error {
		// 创建新的优惠券
		err = tx.Create(&c).Error
		if err != nil {
			return err
		}

		// mysql 更新商品
		updateMap := make(map[string]interface{}, 4)
		updateMap["post_coupon_price"] = postCouponPrice
		updateMap["commission_value"] = commissionValue
		updateMap["created_at"] = createdAt
		updateMap["coupon_id"] = c.ID
		err = tx.Model(product.Goods{}).Where("id = ?", couponInfo.GoodsID).Updates(updateMap).Error
		if err != nil {
			return err
		}

		// es 更新数据
		m := goods.Message{
			ID:              couponInfo.GoodsID,
			CouponID:        c.ID,
			CouponAmount:    couponInfo.Amount,
			CouponBeginTime: beginDTimeStr,
			CouponEndTime:   endDTimeStr,
			CouponTotal:     couponInfo.CouponTotal,
			CouponCover:     couponInfo.CouponCover,
			PostCouponPrice: postCouponPrice,
			CommissionValue: commissionValue,
			CreatedAt:       createdAt,
		}
		sender := goods.NewUpdateEsGoodsInfoSender(m)
		err = sender.Send()
		if err != nil {
			return err
		}
		return nil
	})
}

// 商品收藏
func (*GoodsService) Collect(req request.GoodsCollect) error {
	// 获取商品信息 判断是否过期
	var g product.Goods
	err := global.XTK_DB.Preload("Coupon", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "begin_time", "end_time", "amount")
	}).Select("id", "cover_image_id", "commission_rate", "commission_value", "post_coupon_price", "title", "coupon_id", "price").Take(&g, "id = ?", req.GoodsID).Error
	if err != nil {
		return err
	}
	now := time.Now()
	if g.CouponID != 0 {
		if g.Coupon.BeginTime.After(now) {
			return errors.New("未开始")
		}
		if g.Coupon.EndTime.Before(now) {
			return errors.New("已失效")
		}
	}
	// 判断封面是否为空
	var imageID uint
	if g.CoverImageID == 0 {
		err = global.XTK_DB.Table("goods_image").
			Where("goods_id = ?", req.GoodsID).
			Limit(1). // 限制只获取一条记录
			Pluck("image_id", &imageID).Error
	} else {
		imageID = g.CoverImageID
	}

	var c product.Collect
	c.ID = g.ID
	c.UserID = req.UserID
	c.GoodsID = req.GoodsID
	c.CreatedAt = time.Now()
	c.Title = g.Title
	c.CouponValue = g.Coupon.Amount
	if g.CouponID == 0 {
		c.PostCouponPrice = g.Price
	} else {
		c.PostCouponPrice = g.PostCouponPrice
	}

	c.CommissionValue = g.CommissionValue
	c.CommissionRate = g.CommissionRate
	if g.CouponID == 0 {
		c.CouponEndTime = nil
	} else {
		c.CouponEndTime = &g.Coupon.EndTime
	}
	c.ImageID = imageID

	// 判断是否已经收藏
	err = global.XTK_DB.Where("user_id = ? and goods_id = ?", c.UserID, c.GoodsID).Take(&c).Error
	if err == nil {
		return errors.New("已收藏")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// 获取该用户的收藏商品总数
	var count int64
	err = global.XTK_DB.Model(c).Where("user_id = ?", req.UserID).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 50 {
		return errors.New("最多收藏50件商品")
	}
	// 创建用户商品添加关联关系
	err = global.XTK_DB.Transaction(func(tx *gorm.DB) error {
		// 保存收藏信息
		err = global.XTK_DB.Create(&c).Error
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			return err
		}
		// 删除缓存
		key := "collect:" + strconv.Itoa(int(req.UserID))
		err = global.GVA_REDIS.Del(context.Background(), key).Err()
		return err
	})
	return err
}

// 取消收藏
func (*GoodsService) CancelCollect(req request.GoodsCollect) error {
	err := global.XTK_DB.Transaction(func(tx *gorm.DB) error {
		// 删除key
		key := "collect:" + strconv.Itoa(int(req.UserID))
		err := global.GVA_REDIS.Del(context.Background(), key).Err()
		if err != nil {
			return err
		}
		// 删除商品关联信息
		err = global.XTK_DB.
			Where("goods_id = ? and user_id = ?", req.GoodsID, req.UserID).Delete(product.Collect{}).Error
		return err
	})
	return err
}

// 商品推广
func (*GoodsService) Promotion(req request.GoodsPromotion, userID uint) error {
	// 获取商品信息 判断是否过期
	var g product.Goods
	err := global.XTK_DB.Preload("Coupon", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "begin_time", "end_time", "amount")
	}).Select("id", "cover_image_id", "commission_rate", "commission_value", "post_coupon_price", "title", "coupon_id", "price").Take(&g, "id = ?", req.GoodsID).Error
	if err != nil {
		return err
	}
	now := time.Now()
	if g.CouponID != 0 {
		if g.Coupon.BeginTime.After(now) {
			return errors.New("未开始")
		}
		if g.Coupon.EndTime.Before(now) {
			return errors.New("已失效")
		}
	}
	// 判断封面是否为空
	var imageID uint
	if g.CoverImageID == 0 {
		err = global.XTK_DB.Table("goods_image").
			Where("goods_id = ?", req.GoodsID).
			Limit(1). // 限制只获取一条记录
			Pluck("image_id", &imageID).Error
	} else {
		imageID = g.CoverImageID
	}

	var p product.Promotion
	p.ID = g.ID
	p.UserID = userID
	p.GoodsID = req.GoodsID
	p.CreatedAt = time.Now()
	p.Title = g.Title
	p.CouponValue = g.Coupon.Amount
	if g.CouponID == 0 {
		p.PostCouponPrice = g.Price
	} else {
		p.PostCouponPrice = g.PostCouponPrice
	}

	p.CommissionValue = g.CommissionValue
	p.CommissionRate = g.CommissionRate
	if g.CouponID == 0 {
		p.CouponEndTime = nil
	} else {
		p.CouponEndTime = &g.Coupon.EndTime
	}
	p.ImageID = imageID
	// 创建用户商品添加关联关系
	err = global.XTK_DB.Transaction(func(tx *gorm.DB) error {
		// 保存收藏信息
		err = global.XTK_DB.Save(&p).Error
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			return err
		}
		key := "GoodsPromotion"
		member := strconv.Itoa(int(req.GoodsID))
		_, err = global.GVA_REDIS.ZIncrBy(context.Background(), key, 1, member).Result()
		return err
	})
	return err
}
