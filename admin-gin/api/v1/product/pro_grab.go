package product

import (
	"admin-gin/global"
	"admin-gin/model/common/response"
	"admin-gin/model/product"
	"admin-gin/model/product/request"
	"admin-gin/utils"
	"github.com/gin-gonic/gin"
)

type GrobApi struct{}

func (GrobApi) GetGrabRecord(c *gin.Context) {
	var req request.GrabRecord
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

	list, total, err := goodsService.GrabRecord(req)

	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, "获取成功", c)

}

func (GrobApi) GetGrabLog(c *gin.Context) {
	var req request.ProGrabLog
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.ZhValidate(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var logs []product.GrabLog
	err = global.XTK_DB.Where("grab_record_uuid =?", req.UUID).Order("created_at desc").Find(&logs).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(logs, c)
}
