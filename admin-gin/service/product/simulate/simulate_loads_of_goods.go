package simulate

import (
	"admin-gin/global"
	"admin-gin/model/product"
	"admin-gin/process"
	"admin-gin/utils"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func (*SimService) SimulateLoadOfGoods() {
	// 目前数据库中有10万个数据, 以这10万个为基础 遍历每个数据 复制100份
	var goods []product.Goods
	var pageSize = 1
	var page = 1

	for {
		offset := (page - 1) * pageSize
		err := global.XTK_DB.
			Limit(pageSize).
			Preload("Images", func(db *gorm.DB) *gorm.DB {
				return db.Select("goods_id, url")
			}).
			Offset(offset).Find(&goods).Error
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			continue
		}
		// 拿到了10条真实数据
		for _, g := range goods {
			// 放大100陪
			for i := 0; i < 100; i++ {
				g.ID = 0
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
				// 由于查询出来的关联图片 对应的是原理的 商品id 直接往里放的话 复合组件冲突,不会新增新的关联, 所以要将商品id去除, 才好创建新的关联

				// 放入队列, 刷入数据库
				process.ToMysqlJobCh <- g
			}
		}
		if len(goods) == 0 {
			break
		}
		page++
	}

	fmt.Println("模拟了100 条数据")

	// 清理模拟的基础数据
	CleanSimulateBaseData()
}
