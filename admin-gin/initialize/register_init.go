package initialize

import (
	_ "admin-gin/source/example"
	_ "admin-gin/source/system"
)

// 什么也不做，只导入源包以便可以注册初始化
func init() {
	// do nothing,only import source package so that inits can be registered

}
