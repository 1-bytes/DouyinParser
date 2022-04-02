package bootstrap

import (
	"DouyinParser/pkg/captcha"
	"DouyinParser/pkg/config"
)

// SetupCaptcha 初始化验证码.
func SetupCaptcha() {
	cfg := config.GetString("captcha.style_json")
	captcha.GetInstance().SetConfig(cfg)
}
