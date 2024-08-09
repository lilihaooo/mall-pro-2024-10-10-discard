package product

import (
	"admin-gin/model/common/response"
	"admin-gin/model/product/request"
	"github.com/gin-gonic/gin"
)

type ShopApi struct{}

func (ShopApi) ShopList(c *gin.Context) {
	var req request.ProShopList
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

	list, total, err := shopService.ShopList(req)
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
