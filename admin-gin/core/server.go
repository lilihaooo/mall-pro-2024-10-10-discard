package core

import (
	"admin-gin/global"
	"admin-gin/initialize"
	"admin-gin/service/system"
	"fmt"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	// 启动芒果db
	if global.GVA_CONFIG.System.UseMongo {
		err := initialize.Mongo.Initialization()
		if err != nil {
			zap.L().Error(fmt.Sprintf("%+v", err))
		}
	}
	// 从db加载jwt数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}

	router := initialize.Routers()
	// init admin ApiAuthority
	initialize.AdminApiAuthority()

	router.Static("/form-generator", "./resource/page")
	address := fmt.Sprintf("0.0.0.0:%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, router)
	fmt.Printf(`
	当前版本:v2.6.5
	GVA讨论社区:https://support.qq.com/products/371961
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
`, address)

	// 断开监听打印错误
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
