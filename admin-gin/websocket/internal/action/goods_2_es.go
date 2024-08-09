package action

import (
	"admin-gin/global"
	"admin-gin/websocket/common"
	"fmt"
)

func Goods2ESAction(data interface{}, userID uint) {
	// 初始化业务参数
	//dataBytes, err := json.Marshal(data)
	//if err != nil {
	//	global.GVA_LOG.Error(err.Error())
	//	return
	//}
	//simR := request.SimulateRequest{}
	//err = json.Unmarshal(dataBytes, &simR)
	//if err != nil {
	//	global.GVA_LOG.Error(err.Error())
	//	return
	//}

	// 执行业务
	//if simR.Option == "coupon" {
	// 如果启动任务了其他进程不能启动
	_, ok := global.LockMap.Load("goods_2_es_lock")
	if ok {
		common.WsSendMsg(
			common.MESSAGE_LEVEL_ERROR,
			"耗时任务正在执行， 请勿重复操作！",
			nil,
			common.FORWHERE_ELMESSAGE,
			[]uint{userID},
			"",
		)
		global.GVA_LOG.Info("task_lock is exist")
		return
	}

	global.LockMap.Store("goods_2_es_lock", 1)
	common.WsSendMsg(0, "", true, "goods_2_es_live", []uint{}, "")

	// 调用服务
	//goods2ESService := service.GroupApp.ProductServiceGroup.ToESService
	//goods2ESService.Goods2Es(userID)

	global.LockMap.Delete("goods_2_es_lock")
	common.WsSendMsg(0, "", false, "goods_2_es_live", []uint{}, "")

	fmt.Println("结束迁移")
}

//}
