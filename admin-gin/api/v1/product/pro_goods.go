package product

import (
	"admin-gin/global"
	"admin-gin/model/common/response"
	"admin-gin/model/product"
	"admin-gin/model/product/request"
	"admin-gin/utils"
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
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

// 商品收藏
func (GoodsApi) GoodsCollect(c *gin.Context) {
	var req request.GoodsCollect
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
	// 判断是否越权
	userIDS := utils.GetUserID(c)
	if req.UserID != userIDS {
		global.GVA_LOG.Error("水平越权")
		response.FailWithMessage("水平越权", c)
		return
	}
	// 调用服务
	err := goodsService.Collect(req)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("success", c)
}

// 取消收藏
func (GoodsApi) GoodsCancelCollect(c *gin.Context) {
	var req request.GoodsCollect
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
	// 判断是否越权
	userIDS := utils.GetUserID(c)
	if req.UserID != userIDS {
		global.GVA_LOG.Error("水平越权")
		response.FailWithMessage("水平越权", c)
		return
	}
	// 调用服务
	err := goodsService.CancelCollect(req)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("success", c)
}

// 批量取消收藏
func (GoodsApi) BatchCancelCollect(c *gin.Context) {
	var req request.BatchCollect
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
	userIDS := utils.GetUserID(c)
	err := global.XTK_DB.Transaction(func(tx *gorm.DB) error {
		// 删除key
		key := "collect:" + strconv.Itoa(int(userIDS))
		err := global.GVA_REDIS.Del(context.Background(), key).Err()
		if err != nil {
			return err
		}
		// 删除商品关联信息
		err = global.XTK_DB.Where("id in ?", req.IDs).Delete(product.Collect{}).Error
		return err
	})
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("success", c)
}

// 我的收藏
func (GoodsApi) MyCollect(c *gin.Context) {
	var req request.MyCollect
	if err := c.ShouldBindQuery(&req); err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 设置默认值
	if req.Page == 0 {
		req.Page = 1
	}
	offset := (req.Page - 1) * 20

	userID := utils.GetUserID(c)

	var cList []product.Collect
	var total int64

	q := global.XTK_DB.Where("user_id = ?", userID)

	err := q.Model(product.Collect{}).Count(&total).Error
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = q.Preload("Image").Order("created_at Desc").Offset(offset).Limit(10).Find(&cList).Error
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 返回列表
	response.OkWithDetailed(response.PageResult{
		List:     cList,
		Total:    total,
		Page:     req.Page,
		PageSize: 20,
	}, "获取成功", c)
}

// 商品推广
func (GoodsApi) GoodsPromotion(c *gin.Context) {

}
