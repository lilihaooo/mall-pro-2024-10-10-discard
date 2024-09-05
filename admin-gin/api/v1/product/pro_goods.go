package product

import (
	"admin-gin/global"
	"admin-gin/model/common/response"
	"admin-gin/model/product/request"
	"admin-gin/utils"
	"github.com/gin-gonic/gin"
)

type GoodsApi struct{}

func (GoodsApi) GetGoodsSearchList(c *gin.Context) {
	var req request.ProGoodsSearch
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if req.PageSize <= 0 {
		req.PageSize = 20
	}
	if req.Page <= 1 {
		req.Page = 1
	}

	list, total, err := goodsService.GoodsSearchByMysql(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, "获取成功", c)
}

func (GoodsApi) GoodsInfo(c *gin.Context) {
	var req request.GoodsInfo
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if req.ID == 0 {
		response.FailWithMessage("id 不能为空", c)
		return
	}

	goods, err := goodsService.GoodsInfo(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(goods, c)

}

func (GoodsApi) SuggestionList(c *gin.Context) {
	var req request.Suggestion
	err := c.ShouldBindQuery(&req)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.ZhValidate(&req)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}

	list := goodsService.SearchSuggestionKeywordList(req)
	response.OkWithDetailed(list, "success", c)
}

func (GoodsApi) UploadGoodsImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var req request.GoodsImageUpload
	if err = c.ShouldBindQuery(&req); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err = utils.ZhValidate(req); err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = goodsService.GoodsUploadImage(file, req.ID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("success", c)

}

func (GoodsApi) DeleteGoodsImage(c *gin.Context) {
	var req request.GoodsImageDelete
	if err := c.ShouldBindQuery(&req); err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.ZhValidate(req); err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	err := goodsService.DeleteGoodsImage(req.ID)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("success", c)
}

func (GoodsApi) SetGoodsImageCover(c *gin.Context) {
	var req request.SetGoodsImageCover
	if err := c.ShouldBindJSON(&req); err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.ZhValidate(req); err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}

	err := goodsService.SetCoverImage(req.GoodsID, req.ImageID)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("success", c)
}

func (GoodsApi) UpdateGoodsBaseInfo(c *gin.Context) {
	var req request.UpdateGoodsBaseInfo
	if err := c.ShouldBindJSON(&req); err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.ZhValidate(req); err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}

	err := goodsService.UpdateBaseInfo(req)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("success", c)
}

func (GoodsApi) UpdateGoodsCouponInfo(c *gin.Context) {
	var req request.UpdateGoodsCouponInfo
	if err := c.ShouldBindJSON(&req); err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.ZhValidate(req); err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}

	err := goodsService.UpdateCouponInfo(req)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("success", c)
}
