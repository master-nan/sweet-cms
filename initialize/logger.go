package initialize

import "go.uber.org/zap"

func Logger() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}
