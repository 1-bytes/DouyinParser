package config

import "DouyinParser/pkg/config"

func init() {
	config.Add("app", config.StrMap{
		// 应用名称
		"name": config.Env("APP_NAME", "DouyinParser"),
		// 当前环境，用以区分多环境
		"env": config.Env("APP_ENV", "develop"),
		// APP 安全密钥，务必去创建一个自己的 GUID 作为密钥：https://www.guidgen.com
		"key": config.Env("APP_KEY", "5f90f2f4-7545-4b3e-9b96-4bcbc0343f56"),
		// 是否开启调试模式
		"debug": config.Env("APP_DEBUG", false),
		// JWT 密钥，至少需要 512 位的长度
		"jwt_secret": config.Env("APP_JWT_SECRET", "x01jIE9ymYpbVnR2KTzvaMwdPAQ6rg754LtNeGZlHkuOJ8UqSDWihcF3oCfBsXgAS6aN8aZHsxiIgJXlJUxBPy4P6mkB59CSUPJw34pZ3ONHyK7UZfkQJZ6w20vP585pjpf9xVH9cgsSM9ziTE4QYPaHzDmYoOu1v1FGrAU13DcqFsjDiqx8O7Ty6KJSFeA5kqKE7aEKcl3OfkH3d0y9Ger2Z3Y1Bn0nKyGQvF5rm8jXFcLI6c71EpDbqtLb7o7cfks8xPc24osvvQhNRSJSWAL56wSB20uVixWeOXC5nCCQ5MTZQwA41ZIA5B2U2RwPz17GDSZOLLoYvGugJ3SA0FX7AYxsugRDN3LiWdF6yeKFxy96swyxQEXnVranoSE0n1edswVCrm6drzvEkqQMQs1Lc6gX7VakPNopaSQINsVK5Uz0UthTcu8zXLciEMx7Sl0BwX5Wv7G2SG6u7XQG94hmKtnzBUq5kaFt3flMr9paeMM9YivrCD1pSLhqlOXG"),
	})
}
