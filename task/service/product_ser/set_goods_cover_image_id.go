package product_ser

import (
	"fmt"
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
			Offset(offset).
			Limit(pageSize).
			Find(&ims).Error
		if err != nil {
			global.Logrus.Error(err)
			return err
		}
		if len(ims) == 0 {
			fmt.Println("完成")
			break
		}

		// 遍历每个店铺名
		for _, im := range ims {
			err = global.DB.Model(gormmodel.Goods{}).Where("id = ?", im.GoodsID).Update("cover_image_id", im.ImageID).Error
			if err != nil {
				global.Logrus.Error(err)
				return err
			}

		}
		fmt.Println(pageNum)
		pageNum++
	}

	return nil
}
