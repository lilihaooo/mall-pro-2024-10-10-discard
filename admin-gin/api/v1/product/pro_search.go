package product

import (
	"admin-gin/global"
	"admin-gin/model/common/response"
	"admin-gin/model/product/request"
	"admin-gin/utils"
	"admin-gin/utils/redis"
	"github.com/gin-gonic/gin"
)

type SearchApi struct{}

func (SearchApi) SuggestionList(c *gin.Context) {
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
	list := searchService.SearchSuggestionKeywordList(req)
	response.OkWithDetailed(list, "success", c)
}

func (SearchApi) HotList(c *gin.Context) {
	zset := redis.NewResZSet()
	z, err := zset.GetMembers(c, "HotSearch", 0, -1)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	var list []interface{}
	for _, v := range z {
		list = append(list, v.Member)
	}
	response.OkWithDetailed(list, "success", c)
}
