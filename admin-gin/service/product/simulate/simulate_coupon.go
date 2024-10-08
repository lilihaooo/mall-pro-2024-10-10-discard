package simulate

import (
	"admin-gin/global"
	"admin-gin/model/product"
	"admin-gin/model/product/request"
	"admin-gin/utils"
	"strconv"
	"time"
)

func (*SimService) SimulateCoupon(simReq request.SimulateCouponRequest) (count int) {
	// 中国时区
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}

	// 转换为中国时间
	chinaTime := simReq.BeginDate.In(location)
	date := chinaTime.Format("2006-01-02")
	beginDTimeStr := date + " 00:00:00"
	endDTimeStr := date + " 23:59:59"
	// 将字符串解析为时间格式一定要指定时区
	dayE, err := time.ParseInLocation("2006-01-02 15:04:05", endDTimeStr, location) // 中国时间
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}
	dayB, err := time.ParseInLocation("2006-01-02 15:04:05", beginDTimeStr, location) // 中国时间
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return
	}
	for i := 0; i < simReq.Count; i++ {
		var coModel product.Coupon
		amount := utils.RandomFloat64With1Decimal(simReq.AmountFrom, simReq.AmountTo)
		// 开始时间
		cover := utils.RandomInt(100, 1000)
		total := cover * utils.RandomInt(1, 10)
		coModel.Title = "通用券"
		coModel.Amount = amount
		coModel.MinAmount = amount * 10
		coModel.BeginTime = dayB
		coModel.EndTime = dayE.AddDate(0, 0, simReq.EffectiveDay)
		coModel.CouponTotal = total
		coModel.CouponCover = cover
		coModel.Status = 1

		err = global.XTK_DB.Save(&coModel).Error
		if err != nil {
			global.GVA_LOG.Error(err.Error())
		}
		count++
	}
	global.GVA_LOG.Info("模拟优惠券完成, 共计生成" + strconv.Itoa(count) + "个数据")
	return
}
