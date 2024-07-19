package initialize

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"time"
)

// createLogDirectory 确保日志目录存在
func createLogDirectory() {
	logDir := "./logs"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.MkdirAll(logDir, os.ModePerm)
		if err != nil {
			panic("Failed to create log directory")
		}
	}
}

func Logger() {
	//logger, _ := zap.NewDevelopment()
	//zap.ReplaceGlobals(logger)

	createLogDirectory()

	logFilePath := filepath.Join("./logs", time.Now().Format(time.DateOnly)+".log")

	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    20,   // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 7,    // 日志文件最多保存多少个备份
		MaxAge:     30,   // 文件最多保存多少天
		Compress:   true, // 是否压缩
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.MessageKey = "message"

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),        // 控制台输出
			zapcore.AddSync(lumberJackLogger), // 文件输出
		),
		zap.NewAtomicLevelAt(zap.DebugLevel),
	)

	logger := zap.New(core)
	zap.ReplaceGlobals(logger)
}
