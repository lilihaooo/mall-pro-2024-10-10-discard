package product

import (
	"admin-gin/api/v1"
	"admin-gin/api/v2"
	"github.com/gin-gonic/gin"
)

type GoodsRouter struct{}

func (p *GoodsRouter) InitGoodsRouter(Router *gin.RouterGroup) {
	goodsInfoRouter := Router.Group("goods")
	goodsApiV1 := v1.ApiGroupApp.ProductApiGroup.GoodsApi
	{
		goodsInfoRouter.GET("/search/list", goodsApiV1.GetGoodsSearchList) // 全部商品
		goodsInfoRouter.GET("/info", goodsApiV1.GoodsInfo)                 // 根据ID获得产品信息

		goodsInfoRouter.POST("/image/upload", goodsApiV1.UploadGoodsImage)          // 商品图片上传
		goodsInfoRouter.GET("/image/delete", goodsApiV1.DeleteGoodsImage)           // 商品图片删除
		goodsInfoRouter.POST("/image/setCover", goodsApiV1.SetGoodsImageCover)      // 商品图片删除
		goodsInfoRouter.POST("/updateBaseInfo", goodsApiV1.UpdateGoodsBaseInfo)     // 商品基本信息修改
		goodsInfoRouter.POST("/updateCouponInfo", goodsApiV1.UpdateGoodsCouponInfo) // 商品优惠券信息修改
		goodsInfoRouter.POST("/collect", goodsApiV1.GoodsCollect)                   // 收藏商品
		goodsInfoRouter.POST("/cancelCollect", goodsApiV1.GoodsCancelCollect)       // 取消收藏商品
		goodsInfoRouter.POST("/batchCancelCollect", goodsApiV1.BatchCancelCollect)  // 批量取消收藏商品
		goodsInfoRouter.GET("/myCollect", goodsApiV1.MyCollect)                     // 我的收藏
		goodsInfoRouter.POST("/promotion", goodsApiV1.GoodsPromotion)               // 推广商品
	}

	goodsApiV2 := v2.ApiGroupApp.ProductApiGroup.GoodsApi
	{
		goodsInfoRouter.GET("/v2/search/list", goodsApiV2.GetGoodsSearchList) // 根据ID获得产品信息
	}
}
