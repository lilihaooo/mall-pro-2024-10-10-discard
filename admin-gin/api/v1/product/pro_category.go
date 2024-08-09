package product

import (
	"github.com/gin-gonic/gin"

	"admin-gin/model/common/response"
)

type CategoryApi struct{}

func (CategoryApi) GetCategoryTree(c *gin.Context) {
	tree, err := categoryService.GetCategoryTree()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(tree, c)
}
