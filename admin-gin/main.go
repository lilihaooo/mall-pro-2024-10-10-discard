package main

import (
	_ "net/http/pprof"

	"admin-gin/core"
	"admin-gin/global"
	"admin-gin/initialize"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
	"log"
	"net/http"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title                       Gin-Vue-Admin Swagger API接口文档
// @version                     v2.6.5
// @description                 使用gin+vue进行极速开发的全栈开发基础平台
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	go func() {
		log.Println(http.ListenAndServe("127.0.0.1:6060", nil))
	}()

	global.GVA_VP = core.Viper() // 初始化Viper， 将配置文件解析到global.GVA_LOG中。并动态监听
	initialize.OtherInit()       // 初始化其他  将jwt的过期时间改为 日期格式； 检查JWT缓冲时间; 初始化全局缓存

	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)

	global.GVA_DB = initialize.Gorm() // gorm连接数据库   初始化默认的mysql 通过配置文件mysql

	initialize.Timer() // 定时任务 （定时清理数据库）

	initialize.DBList() // 初始化多数据库 通过配置文件db-list

	initialize.InitEs()
	initialize.InitKafkaProducer()

	global.XTK_DB = global.MustGetGlobalDBByDBName("xtaoke_db") // 将redu_db放到全局变量

	if global.GVA_DB != nil {
		//initialize.RegisterTables()
		// 程序结束前关闭数据库链接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}

	if global.XTK_DB != nil {
		// 程序结束前关闭数据库链接
		db, _ := global.XTK_DB.DB()
		//err := global.XTK_DB.AutoMigrate(
		//
		//	product.Image{},
		//	product.Goods{},
		//)
		//if err != nil {
		//	global.GVA_LOG.Error(err.Error())
		//	return
		//}
		defer db.Close()
	}

	if global.KafkaProducer != nil {
		defer global.KafkaProducer.Close()
	}

	core.RunWindowsServer()

}
