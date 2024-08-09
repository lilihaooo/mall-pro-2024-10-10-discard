package product

import (
	"admin-gin/global"
	"admin-gin/model/product"
	"admin-gin/model/product/request"
	"admin-gin/model/product/response"
	"admin-gin/model/system"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
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
		Preload("Shop", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name, logo, logistics_score, product_score, service_score") // 这里选择 Shop 表中的 id 和 name 字段
		}).
		Preload("Media").
		Preload("Images").
		Preload("Tags").
		Preload("Brand").
		Preload("Coupon").
		Find(&goods, req.ID).
		Error
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
func (*GoodsService) GoodsSearchByEs(req request.ProGoodsSearchV2) (respGoods []response.GoodsSearchV2, total int64, err error) {
	once.Do(func() {
		InitCategory()
	})
	// 构建 复合 查询条件
	boolQuery := elastic.NewBoolQuery()

	// 关键词搜索 - 模糊匹配
	if req.Keyword != "" {
		boolQuery.Must(elastic.NewMatchPhraseQuery("all", req.Keyword).Slop(0))
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
		fmt.Println(p)
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
	// 品牌库   查看品牌名不为空的商品
	if req.IsBrand == 1 {
		boolQuery.MustNot(elastic.NewTermQuery("brand_name.keyword", ""))
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
		return nil, 0, err
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
		return nil, 0, err
	}

	// 解析结果
	for _, hit := range esRes.Hits.Hits {
		var g product.EsGoods
		err = json.Unmarshal(hit.Source, &g)
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			return nil, 0, err
		}

		var gResp response.GoodsSearchV2
		err = copier.Copy(&gResp, g)
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			return nil, 0, err
		}
		respGoods = append(respGoods, gResp)
	}
	return
}

// 获取搜索词条
func (*GoodsService) SearchSuggestionKeywordList(req request.Suggestion) (collect []string) {
	keyword := req.Keyword
	// 构建 Suggest 查询
	suggestName := "keyword_suggest"
	suggestQuery := elastic.NewCompletionSuggester(suggestName).
		Prefix(keyword).
		Field("keyword").
		Size(10)

	// 执行查询
	searchResult, err := global.ESClient.Search().
		Index("suggestion_index"). // 设置索引名称
		Suggester(suggestQuery).   // 添加 suggest 查询
		Do(context.Background())   // 执行查询
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil
	}

	// 处理查询结果
	if suggestResult, found := searchResult.Suggest[suggestName]; found {
		for _, entry := range suggestResult {
			for _, option := range entry.Options {
				collect = append(collect, option.Text)
			}
		}
	} else {
		global.GVA_LOG.Info("无推荐")
		return
	}
	return
}
