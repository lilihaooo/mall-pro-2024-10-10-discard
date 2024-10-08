package simulate

import (
	"admin-gin/global"
	"admin-gin/model/product"
)

func (*SimService) GoodsAddTags() {
	var goodsIDChan = make(chan uint, 10)
	doneChan := make(chan bool) // 用于通知主 goroutine 所有操作完成
	go func() {

		for goodsID := range goodsIDChan {
			ta := MakeRandomTags()

			var g product.Goods
			g.ID = goodsID
			err := global.XTK_DB.Model(&g).Association("Tags").Clear()
			if err != nil {
				global.GVA_LOG.Error(err.Error())
				return
			}
			err = global.XTK_DB.Model(&g).Association("Tags").Append(&ta)
			if err != nil {
				global.GVA_LOG.Error(err.Error())
				return
			}
		}
		doneChan <- true // 通知主 goroutine 所有操作完成
	}()

	pageSize := 500
	page := 1
	for {
		var goods []product.Goods
		offset := (page - 1) * pageSize
		err := global.XTK_DB.Select("id").Limit(pageSize).Offset(offset).Find(&goods).Error
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			break
		}
		if len(goods) == 0 {
			break
		}
		for _, g := range goods {
			goodsIDChan <- g.ID
		}
		page++
	}
	close(goodsIDChan)
	<-doneChan // 等待所有操作完成
	// 清理模拟的基础数据
	CleanSimulateBaseData()
}
