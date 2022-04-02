package bootstrap

import "DouyinParser/pkg/validation"

// SetupValidation 初始化表单验证器.
func SetupValidation() {
	validation.Init()
}
