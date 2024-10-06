package product_ser

import (
	"fmt"
	"gorm.io/gorm"
	"task/global"
	"task/models/gormmodel"
)

func (s *GoodsService) SetGoodsCoverImageID() error {
	// 商品图片表
	pageNum := 1
	pageSize := 100 // 每批次处理 500 条数据
	for {
		offset := (pageNum - 1) * pageSize
		var ims []gormmodel.GoodsImage
		err := global.DB.
			Preload("Goods", func(db *gorm.DB) *gorm.DB {
				return db.Select("id, title")
			}).
			Offset(offset).
			Limit(pageSize).
			Find(&ims).Error
		if err != nil {
			global.Logrus.Error(err)
			return err
		}
		if len(ims) == 0 {
			break
		}

		// 遍历每个店铺名
		for _, im := range ims {
			res := global.DB.Model(gormmodel.Goods{}).
				Where("title = ?", im.Goods.Title).
				Update("cover_image_id", im.ImageID)
			err = res.Error
			if err != nil {
				global.Logrus.Error(err)
				return err
			}
			if res.RowsAffected == 0 {
				fmt.Println("失败")
			}

		}
		fmt.Println(pageNum)
		pageNum++
	}

	return nil
}
