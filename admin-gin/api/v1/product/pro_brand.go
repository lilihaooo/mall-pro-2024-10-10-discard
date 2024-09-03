package product

import (
	"admin-gin/global"
	"admin-gin/model/common/response"
	"admin-gin/model/product"
	"github.com/gin-gonic/gin"
)

type TagApi struct{}

func (TagApi) TagList(c *gin.Context) {
	// 获取全部标签
	var tags []product.Tag
	if err := global.XTK_DB.Find(&tags).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(tags, c)

}
