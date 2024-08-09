package core

import (
	"admin-gin/core/internal"
	"admin-gin/global"
	"admin-gin/utils"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// Zap 获取 zap.Logger
// Author [SliverHorn](https://github.com/SliverHorn)
func Zap() (logger *zap.Logger) {
	// 获取日志目录， 没有就创建
	if ok, _ := utils.PathExists(global.GVA_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.GVA_CONFIG.Zap.Director)
		_ = os.Mkdir(global.GVA_CONFIG.Zap.Director, os.ModePerm)
	}

	// 获得日志级别
	// 例如配置文件的级别小于FatalLevel级别， 此方法会将FatalLevel级别前包括FatalLevel 加入到levels。
	levels := global.GVA_CONFIG.Zap.Levels()

	length := len(levels)
	cores := make([]zapcore.Core, 0, length)
	for i := 0; i < length; i++ {
		core := internal.NewZapCore(levels[i])
		cores = append(cores, core)
	}

	logger = zap.New(zapcore.NewTee(cores...))
	if global.GVA_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
