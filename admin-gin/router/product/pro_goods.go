package product

import (
	"admin-gin/api/v1"
	"admin-gin/api/v2"
	"admin-gin/router/common"
	"github.com/gin-gonic/gin"
)

type GoodsRouter struct{}

func (p *GoodsRouter) InitGoodsRouter(Router *gin.RouterGroup) {
	var apiGroup = "goods"
	goodsInfoRouter := Router.Group(apiGroup)
	goodsApiV1 := v1.ApiGroupApp.ProductApiGroup.GoodsApi
	{
		common.RegisterRouteWithComment(goodsInfoRouter, "GET", "search/list", "全部商品列表", apiGroup, goodsApiV1.GetGoodsSearchList)
		common.RegisterRouteWithComment(goodsInfoRouter, "GET", "info", "获得产品信息", apiGroup, goodsApiV1.GoodsInfo)

		common.RegisterRouteWithComment(goodsInfoRouter, "POST", "image/upload", "商品图片上传", apiGroup, goodsApiV1.UploadGoodsImage)
		common.RegisterRouteWithComment(goodsInfoRouter, "GET", "image/delete", "商品图片删除", apiGroup, goodsApiV1.DeleteGoodsImage)
		common.RegisterRouteWithComment(goodsInfoRouter, "POST", "image/setCover", "设置商品图片", apiGroup, goodsApiV1.SetGoodsImageCover)
		common.RegisterRouteWithComment(goodsInfoRouter, "POST", "updateBaseInfo", "商品基本信息修改", apiGroup, goodsApiV1.UpdateGoodsBaseInfo)
		common.RegisterRouteWithComment(goodsInfoRouter, "POST", "updateCouponInfo", "商品优惠券信息修改", apiGroup, goodsApiV1.UpdateGoodsCouponInfo)
		common.RegisterRouteWithComment(goodsInfoRouter, "POST", "collect", "收藏商品", apiGroup, goodsApiV1.GoodsCollect)
		common.RegisterRouteWithComment(goodsInfoRouter, "POST", "cancelCollect", "取消收藏商品", apiGroup, goodsApiV1.GoodsCancelCollect)
		common.RegisterRouteWithComment(goodsInfoRouter, "POST", "batchCancelCollect", "批量取消收藏商品", apiGroup, goodsApiV1.BatchCancelCollect)
		common.RegisterRouteWithComment(goodsInfoRouter, "GET", "myCollect", "我的收藏", apiGroup, goodsApiV1.MyCollect)
		common.RegisterRouteWithComment(goodsInfoRouter, "GET", "myPromotion", "我的推广", apiGroup, goodsApiV1.MyPromotion)
		common.RegisterRouteWithComment(goodsInfoRouter, "POST", "promotion", "推广商品", apiGroup, goodsApiV1.GoodsPromotion)
	}

	goodsApiV2 := v2.ApiGroupApp.ProductApiGroup.GoodsApi
	{
		common.RegisterRouteWithComment(goodsInfoRouter, "GET", "v2/search/list", "获得商品搜索列表信息", apiGroup, goodsApiV2.GetGoodsSearchList)
	}
}
