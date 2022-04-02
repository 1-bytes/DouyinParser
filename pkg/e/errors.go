package e

import "DouyinParser/pkg/logger"

var MsgFlags = map[int]string{
	SUCCESS:                  "success",
	ERROR:                    "fail",
	InvalidParams:            "请求参数错误",
	ErrorDouyinParserUrlFail: "解析抖音URL失败",
}

// GetMsg get error information based on Code.
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}

// HasError any errors will be saved to the log.
func HasError(err error) bool {
	if err != nil {
		sugar := logger.Logger.Sugar()
		sugar.Errorf("An unpredictable error was caught: %s", err)
		return true
	}
	return false
}
