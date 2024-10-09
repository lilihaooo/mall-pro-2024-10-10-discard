package product_ser

import (
	"fmt"
	"task/global"
	"task/models/gormmodel"
)

func (s *GoodsService) SaveGoodsImage() error {
	var gis []gormmodel.GoodsImage
	var maxID uint
	maxID = 0
	for {
		err := global.DB.
			Model(gormmodel.Goods{}).
			Where("id > ?", maxID).
			Where("data_from = 2").
			Order("id asc").
			Limit(30000).
			Select("id as goods_id", "cover_image_id as image_id").
			Scan(&gis).Error
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		if len(gis) == 0 {
			fmt.Println("完成")
			return nil
		}

		res := global.DB.Save(&gis)
		if res.Error != nil {
			fmt.Println(err.Error())
			return err
		}
		fmt.Println(res.RowsAffected)
		maxID = gis[len(gis)-1].GoodsID
	}
}
