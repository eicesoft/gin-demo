package router

import (
	"eicesoft/web-demo/config"
	"eicesoft/web-demo/internal/middleware"
	"eicesoft/web-demo/pkg/core"
	"eicesoft/web-demo/pkg/db"
	"eicesoft/web-demo/pkg/env"
	"eicesoft/web-demo/pkg/logger"
	"eicesoft/web-demo/pkg/mux"
	"fmt"

	"go.uber.org/zap"
)

func InitMux() (core.Mux, error) {
	loggers, err := logger.NewJSONLogger(
		logger.WithDebugLevel(),
		logger.WithField("app", fmt.Sprintf("%s[%s]", config.Get().Server.Name, env.Get().Value())),
		logger.WithTimeLayout("2006-01-02 15:04:05"),
		logger.WithFileRotationP(config.ProjectLogFile()),
	)
	if err != nil {
		panic(err)
	}
	defer loggers.Sync()

	// 初始化 DB
	dbRepo, err := db.New()
	if err != nil {
		loggers.Fatal("new db err", zap.Error(err))
	}

	m, err := core.New(loggers, dbRepo)
	if err != nil {
		panic(err)
	}

	r := new(mux.Resource)
	r.Mux = m
	r.Middles = middleware.New(loggers)
	//r.RegistryRouters()
	setApiRouter(r)

	return m, nil
}
