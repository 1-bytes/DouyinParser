package bootstrap

import configs "DouyinParser/config"

// Setup 初始化指定的服务.
func Setup() {
	autoLoader(
		configs.Initialize, // 配置文件
		//SetupDB,            // 数据库
		SetupLogger, // 日志
		//SetupValidation,    // 表单验证
		SetupCaptcha, // 验证码
	)
}

// autoLoader 自动加载初始化.
func autoLoader(funcName ...func()) {
	// 只是单纯的初始化服务模块，没有参数，没有返回值！！
	for _, v := range funcName {
		v()
	}
}
