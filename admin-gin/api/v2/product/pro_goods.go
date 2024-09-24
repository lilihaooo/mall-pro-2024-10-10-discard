package product

import (
	"admin-gin/model/common/response"
	"admin-gin/model/product/request"
	"admin-gin/utils"
	"github.com/gin-gonic/gin"
)

type GoodsApi struct{}

func (GoodsApi) GetGoodsSearchList(c *gin.Context) {
	var req request.ProGoodsSearchV2
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if req.OrderType == "" {
		req.OrderType = "DESC"
	}
	if req.PageSize <= 0 {
		req.PageSize = 20
	}
	if req.Page <= 1 {
		req.Page = 1
	}

	// 获取用户id
	userID := utils.GetUserID(c)

	list, total, err := goodsService.GoodsSearchByEs(req, userID)
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
