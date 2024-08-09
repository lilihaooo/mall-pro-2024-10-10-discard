package product

import (
	"admin-gin/model/common/response"
	"admin-gin/model/product/request"
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

	// 查看参数
	//v := reflect.ValueOf(req)
	//typeOfReq := v.Type()
	//for i := 0; i < v.NumField(); i++ {
	//	fmt.Printf("%s: %v\n", typeOfReq.Field(i).Name, v.Field(i).Interface())
	//}

	list, total, err := goodsService.GoodsSearchByEs(req)
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
