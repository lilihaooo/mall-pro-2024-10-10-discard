package initialize

import (
	"admin-gin/global"
	"admin-gin/model/system"
	systemReq "admin-gin/model/system/request"
	"admin-gin/service"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

// 管理员授权全部路由信息
func AdminApiAuthority() {
	// 为超级管理员添加全部角色
	global.GVA_LOG.Info("请保证超级管理员的AuthorityID(角色id) 为1")

	rules := [][]string{}
	var cmr systemReq.CasbinInReceive
	// 将全部路由信息赋值给cmr
	cmr.AuthorityId = 1
	for _, api := range global.Apis {
		rules = append(rules, []string{"1", api.Path, api.Method})

		cmr.CasbinInfos = append(cmr.CasbinInfos, systemReq.CasbinInfo{
			Path:   api.Path,
			Method: api.Method,
		})
	}

	// 清空全部授权信息
	if err := global.GVA_DB.Unscoped().Where("1 = 1").Delete(&gormadapter.CasbinRule{}).Error; err != nil {
		panic(err)
	}

	// 清空api表
	var api system.SysApi
	if err := global.GVA_DB.Unscoped().Where("1 = 1").Delete(api).Error; err != nil {
		panic(err)
	}

	err := global.GVA_DB.Transaction(func(tx *gorm.DB) error {

		// 批量创建 api
		if err := tx.Create(&global.Apis).Error; err != nil {
			return err
		}
		// 为超级管理员重新生成全部api授权
		casbinService := service.GroupApp.SystemServiceGroup.CasbinService
		e := casbinService.Casbin()
		_, err := e.AddPolicies(rules)
		if err != nil {
			return err
		}
		// 是否内存
		global.Apis = nil
		return nil
	})
	if err != nil {
		panic(err)
	}
}
