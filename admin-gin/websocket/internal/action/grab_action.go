package action

import (
	"admin-gin/global"
	"admin-gin/model/product/request"
	"admin-gin/service"
	"admin-gin/service/product/goods_grab"
	"admin-gin/websocket/common"
	"encoding/json"
	"fmt"
)

const (
	GRAB_BEGIN = 1
	GRAB_END   = 0
	GRAB_LIVE  = 2
)

func GrabAction(data interface{}, userID uint) {
	// 初始化业务参数
	dataBytes, err := json.Marshal(data)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}
	grabR := request.GrabRequest{}
	err = json.Unmarshal(dataBytes, &grabR)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}
	// 执行业务
	if grabR.Option == GRAB_BEGIN {
		// 如果启动任务了其他进程不能启动
		_, ok := global.LockMap.Load("grab_lock")
		if ok {
			wsMsgCode := common.MESSAGE_LEVEL_ERROR
			wsMsgMsg := "任务正在执行， 请勿重复操作！"
			common.WsSendMsg(
				wsMsgCode,
				wsMsgMsg,
				nil,
				"ElMessage",
				[]uint{userID},
				"",
			)
			global.GVA_LOG.Info("grab_lock is exist")
			return
		}

		global.LockMap.Store("grab_lock", 1)
		goods_grab.IsCancel = false
		// 开始抓取
		common.WsSendMsg(0, "", true, "grab_live", []uint{}, "")
		grabService := service.GroupApp.ProductServiceGroup.GoodsGrabService
		grabService.GoodsGrab(userID, grabR)
		global.LockMap.Delete("grab_lock")
		common.WsSendMsg(0, "", false, "grab_live", []uint{}, "")
		fmt.Println("结束采集释放资源")
	}

	if grabR.Option == GRAB_END {
		// 如果启动任务了其他进程不能启动
		_, ok := global.LockMap.Load("grab_lock")
		if !ok {
			wsMsgCode := common.MESSAGE_LEVEL_ERROR
			wsMsgMsg := "没有正在执行的任务！"
			common.WsSendMsg(wsMsgCode, wsMsgMsg, nil, common.FORWHERE_ELMESSAGE, []uint{userID}, "")
			return
		}
		// 判断该任务是否是当前用户开启的任务
		if userID != goods_grab.UserID {
			wsMsgCode := common.MESSAGE_LEVEL_ERROR
			wsMsgMsg := "无权终止"
			common.WsSendMsg(wsMsgCode, wsMsgMsg, nil, common.FORWHERE_ELMESSAGE, []uint{userID}, "")
			return
		}
		if goods_grab.IsCancel {
			wsMsgCode := common.MESSAGE_LEVEL_ERROR
			wsMsgMsg := "已关闭"
			common.WsSendMsg(wsMsgCode, wsMsgMsg, nil, common.FORWHERE_ELMESSAGE, []uint{userID}, "")
			return
		}
		wsMsgCode := common.MESSAGE_LEVEL_INFO
		wsMsgMsg := "正在终止！"
		common.WsSendMsg(wsMsgCode, wsMsgMsg, nil, common.FORWHERE_GRABACTIONLOG, []uint{userID}, "")
		goods_grab.IsCancel = true
	}
	if grabR.Option == GRAB_LIVE {
		_, ok := global.LockMap.Load("grab_lock")
		common.WsSendMsg(0, "", ok, common.FORWHERE_GRABLIVE, []uint{}, "")
		return
	}
}
