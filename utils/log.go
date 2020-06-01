package utils

import (
	"encoding/json"
	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

func InitZap() *zap.SugaredLogger {
	configJson := []byte(`
{
	"level": "debug",
	"encoding": "json",
	"outputPaths": ["./pgman.log"],
	"errorOutputPaths": ["stderr", "./pgman.error.log"],
	"encoderConfig": {
		"messageKey": "message",
		"levelKey": "level",
	    "levelEncoder": "lowercase",
		"timeKey": "time",
		"timeEncoder": "epoch"
	}
}
	`)

	var cfg zap.Config
	if err := json.Unmarshal(configJson, &cfg); err != nil {
		panic(err)
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	Log = logger.Sugar()
	return Log
}
