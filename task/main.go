package main

import (
	"task/core"
	"task/global"
	"task/models/esmodel"
	"task/service/product_ser"
)

func main() {
	core.InitConf()
	core.InitEs()
	core.InitGorm()
	defer core.CloseGormLogFile()
	core.InitLogrus()
	defer core.CloseLogFile()

	// 任务一: 自动补全
	//saveSuggestionIndex() 将清空全部自动补全索引!!! 谨慎
	////将补全词语加入关键词库
	//sugService := new(product_ser.SuggestionService)
	//sugService.EsAddSuggestionKeyword()
	//

	//任务二: Mysql to Es
	//saveGoodsIndex() //将清空全部goods索引!!! 谨慎
	proService := new(product_ser.GoodsService)
	proService.GoodsMysql2EsTask()

	//任务三 品牌
	//获取店铺名 判断是否为官方旗舰店 然后为商品添加 品牌信息
	//proService := new(product_ser.GoodsService)
	//err := proService.SetGoodsBrands()
	//if err != nil {
	//	fmt.Println("错误")
	//}

	// 任务四 封面图刷为 不是null
	//proService := new(product_ser.GoodsService)
	//err := proService.SetGoodsCoverImageID()
	//if err != nil {
	//	fmt.Println("错误")
	//}

	// 任务五 数据库备份
	//proService := new(product_ser.GoodsService)
	//err := proService.MysqlBackup()
	//if err != nil {
	//	fmt.Println("错误")
	//}

}

func saveSuggestionIndex() {
	//创建索引库
	eg := esmodel.Suggestion{}
	err := eg.CreateIndex()
	if err != nil {
		global.Logrus.Error(err.Error())
		return
	}
}

func saveGoodsIndex() {
	//创建索引库
	eg := esmodel.Goods{}
	err := eg.CreateIndex()
	if err != nil {
		global.Logrus.Error(err.Error())
		return
	}
}
