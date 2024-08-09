package action

import (
	"admin-gin/global"
	"admin-gin/service"
	"admin-gin/websocket/common"
)

func GoodsAddTagsAction(userID uint) {
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

	_, ok := global.LockMap.Load("task_lock")
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

	global.LockMap.Store("task_lock", 1)
	common.WsSendMsg(0, "", true, "task_live", []uint{}, "")

	defer func() {
		global.LockMap.Delete("task_lock")
		common.WsSendMsg(0, "", false, "task_live", []uint{}, "")
	}()

	//调用服务
	simService := service.GroupApp.ProductServiceGroup.SimService
	simService.GoodsAddTags()

}
