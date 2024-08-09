package action

import (
	"admin-gin/global"
	"admin-gin/model/product/request"
	"admin-gin/service"
	"admin-gin/utils"
	"admin-gin/websocket/common"
	"encoding/json"
	"strconv"
)

// 模拟优惠券数据
func SimulateCouponAction(data interface{}, userID uint) {
	// 初始化业务参数
	dataBytes, err := json.Marshal(data)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}
	simR := request.SimulateCouponRequest{}
	err = json.Unmarshal(dataBytes, &simR)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}

	// 如果启动任务了其他进程不能启动
	_, ok := global.LockMap.Load("simulate_coupon_lock")
	if ok {
		wsMsgCode := common.MESSAGE_LEVEL_ERROR
		wsMsgMsg := "模拟优惠券任务正在执行， 请勿重复操作！"
		common.WsSendMsg(wsMsgCode, wsMsgMsg, nil, "ElMessage", []uint{userID}, "")
		global.GVA_LOG.Info("simulate_coupon_lock is exist")
		return
	}

	global.LockMap.Store("simulate_coupon_lock", 1)
	common.WsSendMsg(0, "", true, "simulate_coupon_live", []uint{}, "")

	defer func() {
		global.LockMap.Delete("simulate_coupon_lock")
		common.WsSendMsg(0, "", false, "simulate_coupon_live", []uint{}, "")
	}()

	// 调用服务
	simService := service.GroupApp.ProductServiceGroup.SimService
	// 校验参数
	err = utils.ZhValidate(&simR)
	if err != nil {
		common.WsSendMsg(
			common.MESSAGE_LEVEL_ERROR,
			err.Error(),
			nil,
			common.FORWHERE_ELMESSAGE,
			[]uint{userID},
			"",
		)
		return
	}

	// 最多一次性生成100个数据
	if simR.Count > 100 {
		simR.Count = 100
	}
	if simR.AmountFrom >= simR.AmountTo {
		common.WsSendMsg(
			common.MESSAGE_LEVEL_ERROR,
			"券面额设置错误",
			nil,
			common.FORWHERE_ELMESSAGE,
			[]uint{userID},
			"",
		)
		return
	}

	count := simService.SimulateCoupon(simR)
	common.WsSendMsg(
		common.MESSAGE_LEVEL_INFO,
		"模拟优惠券完成, 共计生成"+strconv.Itoa(count)+"个数据",
		nil,
		common.FORWHERE_ELMESSAGE,
		[]uint{userID},
		"",
	)

}
