package middleware

import (
	"strconv"
	"strings"

	"admin-gin/global"
	"admin-gin/model/common/response"
	"admin-gin/service"
	"admin-gin/utils"
	"github.com/gin-gonic/gin"
)

var casbinService = service.GroupApp.SystemServiceGroup.CasbinService

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		waitUse, _ := utils.GetClaims(c)
		//获取请求的PATH
		path := c.Request.URL.Path
		obj := strings.TrimPrefix(path, global.GVA_CONFIG.System.RouterPrefix)
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := strconv.Itoa(int(waitUse.AuthorityId))

		// 判断策略中是否存在
		e := casbinService.Casbin()

		success, _ := e.Enforce(sub, obj, act) // 执行
		if !success {
			response.FailWithDetailed(gin.H{}, "权限不足", c)
			c.Abort()
			return
		}
		c.Next()
	}
}
