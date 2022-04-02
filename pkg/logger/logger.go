package logger

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.Logger

// InitZapLogger 初始化日志.
func InitZapLogger(logFilePath string) {
	if gin.IsDebugging() {
		Logger, _ = zap.NewProduction()
		return
	}
	writeSyncer := GetLogWriter(logFilePath)
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	Logger = zap.New(core, zap.AddCaller())
}

// GetLogWriter .
func GetLogWriter(filepath string) zapcore.WriteSyncer {
	file, _ := os.Create(filepath)
	return zapcore.AddSync(file)
}
