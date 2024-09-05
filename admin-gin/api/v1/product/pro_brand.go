package product

import (
	"admin-gin/global"
	"admin-gin/model/common/response"
	"admin-gin/model/product"
	"github.com/gin-gonic/gin"
)

type BrandApi struct{}

func (BrandApi) BrandList(c *gin.Context) {
	// 获取全部标签
	var brands []product.Brand
	if err := global.XTK_DB.Find(&brands).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(brands, c)
}
