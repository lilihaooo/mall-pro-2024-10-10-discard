package initialize

import (
	"github.com/songzhibin97/gkit/cache/local_cache"

	"admin-gin/global"
	"admin-gin/utils"
)

// 初始化全局本地缓存，默认过期时间为配置文件中的值
// 检查缓冲时间的格式

func OtherInit() {
	// 解析过期时间str为时间格式: 配置文件的格式1d33m3s的格式， 由于不能解析d所以要将d(天）转为24h, 才能解析为时间格式
	dr, err := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}

	// 解析缓冲时间： 实际没有用，起检查作用
	_, err = utils.ParseDuration(global.GVA_CONFIG.JWT.BufferTime)
	if err != nil {
		panic(err)
	}

	// 初始化全局黑色缓存
	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)
}
