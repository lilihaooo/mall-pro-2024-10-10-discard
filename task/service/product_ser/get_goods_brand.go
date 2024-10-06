package product_ser

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
	"task/global"
	"task/models/gormmodel"
)

func (s *GoodsService) SetGoodsBrands() error {

	fmt.Println("更新品牌表")
	// 遍历店铺表
	var shops []gormmodel.Shop
	err := global.DB.
		Where("name LIKE ?", "%官方旗舰店").
		Select("id", "name").
		Find(&shops).Error

	if err != nil {
		global.Logrus.Error(err)
		return err
	}

	// 遍历每个店铺名
	for _, shop := range shops {
		if strings.Contains(shop.Name, "官方旗舰店") {
			// 将该名字保存到品牌表
			// 修改本店铺中全部商品的品牌id
			err = global.DB.Transaction(func(tx *gorm.DB) error {
				// 添加品牌信息
				var b gormmodel.Brand
				b.Name = strings.TrimSuffix(shop.Name, "官方旗舰店")
				err = tx.Create(&b).Error
				if err != nil {
					global.Logrus.Error(err)
					return err
				}
				// 修改商品品牌
				err = tx.Table("goods").Where("shop_id=?", shop.ID).Update("brand_id", b.ID).Error
				if err != nil {
					return err
				}
				return nil
			})

		}
	}
	return nil
}
