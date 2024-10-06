package product_ser

import (
	"context"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
	"strconv"
	"task/global"
	"task/models/esmodel"
	"task/models/gormmodel"
	"time"
)

type GoodsService struct{}

type categoryAPart struct {
	ID       uint
	ParentID uint
	Name     string
}

var categoryMap = make(map[uint]categoryAPart)

func (s *GoodsService) GoodsMysql2EsTask() {
	// 查询出全部分类信息 用于搜索以及响应
	initCategoryMap()
	num := 0
	a := time.Now()
	var maxGoodsID uint = 0
	var batchCount = 1000
	for {
		// 从数据库中获取100条记录
		var goods []gormmodel.Goods
		err := global.DB.
			Preload("Shop", func(db *gorm.DB) *gorm.DB {
				return db.Select("id, name")
			}).
			Preload("Coupon", func(db *gorm.DB) *gorm.DB {
				return db.Select("id, title, amount, begin_time, end_time, coupon_total, coupon_cover")
			}).
			Preload("Image").
			Preload("Images").
			Preload("Brand", func(db *gorm.DB) *gorm.DB {
				return db.Select("id, name")
			}).
			Preload("Media", func(db *gorm.DB) *gorm.DB {
				return db.Select("uid, media_name")
			}).
			Preload("Tags", func(db *gorm.DB) *gorm.DB {
				return db.Select("id")
			}).
			Preload("Platform", func(db *gorm.DB) *gorm.DB {
				return db.Select("id, name")
			}).
			Order("id asc").
			Where("id > ?", maxGoodsID).
			Where("is_to_es = 0").
			Limit(batchCount).
			Find(&goods).Error
		if err != nil {
			global.Logrus.Error(err.Error())
			panic(err)
		}
		var ids = make([]uint, len(goods))

		var esGoods []esmodel.Goods
		for i, g := range goods {
			ids[i] = g.ID
			esg := esmodel.Goods{}
			esg.ID = g.ID
			esg.Title = g.Title
			esg.Description = g.Description
			esg.CategoryID = g.CategoryID
			cName := findLongCategoryName(g.CategoryID, categoryMap)
			esg.CategoryName = cName
			esg.ShopName = g.Shop.Name
			esg.CouponID = g.CouponID
			esg.CouponBeginTime = g.Coupon.BeginTime.Format("2006-01-02 15:04:05")
			esg.CouponEndTime = g.Coupon.EndTime.Format("2006-01-02 15:04:05")
			esg.CouponAmount = g.Coupon.Amount
			esg.CouponTotal = g.Coupon.CouponTotal
			esg.CouponCover = g.Coupon.CouponCover
			var img = ""
			if len(g.Images) > 0 {
				img = g.Images[0].Url
			}
			if g.Image.Url != "" {
				img = g.Image.Url
			}

			esg.CoverImage = img
			esg.Status = g.Status
			esg.SalesAll = g.SalesAll
			esg.SalesMonth = g.SalesMonth
			esg.SalesDay = g.SalesDay
			esg.Sales2Hour = g.Sales2Hour
			esg.PlatformName = g.Platform.Name
			esg.PlatformID = g.PlatformID
			esg.MediaUID = *g.MediaUID
			esg.MediaName = g.Media.MediaName
			esg.CommissionRate = g.CommissionRate
			esg.CommissionValue = g.CommissionValue
			esg.PostCouponPrice = g.PostCouponPrice
			esg.Price = g.Price
			esg.Inventory = g.Inventory
			esg.ExperienceScore = g.ExperienceScore
			esg.CreatedAt = g.CreatedAt.Format("2006-01-02 15:04:05")
			esg.PushHot = g.PushHot
			esg.DataFrom = g.DataFrom
			esg.BrandName = g.Brand.Name
			esg.BrandID = g.Brand.ID

			var etag []uint
			for _, one := range g.Tags {
				etag = append(etag, one.ID)
			}
			esg.Tags = etag

			esGoods = append(esGoods, esg)
		}

		//将商品表的is_to_es字段改为1
		err = global.DB.Model(&gormmodel.Goods{}).Where("id in ?", ids).Update("is_to_es", 1).Error
		if err != nil {
			global.Logrus.Error(err.Error())
		} else {
			s.toEs(esGoods)
			l := len(goods)
			num += l
			global.Logrus.Info("已同步: " + strconv.Itoa(num) + " 条记录")
		}

		count := len(goods)
		if count < batchCount {
			b := time.Since(a)
			global.Logrus.Info("完成, 用时: " + b.String())
			break
		}
		maxGoodsID = goods[batchCount-1].ID

		//
		//global.ESClient.Flush()
		//// 使用Count API来获取索引中的总记录数量
		//total, _ := global.ESClient.Count("goods_index").Do(context.Background())
		//// 打印总记录数量
		//fmt.Printf("总记录数量为：%d\n", total)
		//}

	}
}

func initCategoryMap() {
	var categoryAParts []categoryAPart
	err := global.DB.Model(gormmodel.Category{}).Select("id, parent_id, name").Scan(&categoryAParts).Error
	if err != nil {
		global.Logrus.Error(err.Error())
	}
	// 将全部分类信息组装为一个map 用于查找long(path)分类名
	for _, one := range categoryAParts {
		categoryMap[one.ID] = one
	}
}

func (s *GoodsService) toEs(esGoods []esmodel.Goods) {
	// note 可以将商品id指定为文档id
	// 构建批量请求
	bulkRequest := global.ESClient.Bulk()
	for _, item := range esGoods {
		indexReq := elastic.NewBulkIndexRequest().Index(item.Index()).Doc(item)
		bulkRequest = bulkRequest.Add(indexReq)
	}

	// 执行批量请求
	bulkResponse, err := bulkRequest.Do(context.Background()) // Refresh("true")立刻刷新
	//bulkResponse, err := bulkRequest.Do(context.Background())
	if err != nil {
		global.Logrus.Error(err)
		return
	}

	// 处理批量响应
	if bulkResponse.Errors {
		// 处理错误
		for _, failedItem := range bulkResponse.Failed() {
			global.Logrus.Error("失败加入的文档 " + failedItem.Id)
		}
	}
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
