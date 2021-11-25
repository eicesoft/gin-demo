package main

import (
	"context"
	"eicesoft/web-demo/config"
	"eicesoft/web-demo/pkg/db"
	"eicesoft/web-demo/pkg/env"
	"eicesoft/web-demo/pkg/logger"
	"eicesoft/web-demo/pkg/shutdown"
	"eicesoft/web-demo/router"
	"flag"
	"fmt"
	"github.com/zh-five/xdaemon"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	if !flag.Parsed() {
		flag.Parse()
	}

	rand.Seed(time.Now().Unix())
}

// @title Gin MVC demo service
// @version 0.1.1
// @description  This is a sample server for gin mvc demo
// @contact.name kelezyb
// @contact.url
// @contact.email eicesoft@qq.com
// @license.name MIT
// @BasePath
func main() {
	if env.Get().IsDaemon() == true {
		xdaemon.Background("", true)
	}

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

	mux, err := router.InitMux(loggers, dbRepo)
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Addr:           ":" + config.Get().Server.Port,
		Handler:        mux,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 28, //256M
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			loggers.Fatal("http server startup err", zap.Error(err))
		}
	}()

	shutdown.NewHook().Close(
		// 关闭 http server
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				loggers.Error("server shutdown err", zap.Error(err))
			} else {
				loggers.Info("server shutdown success")
			}
		},

		func() {
			if err := dbRepo.DbWClose(); err != nil {
				loggers.Error("dbw close err", zap.Error(err))
			} else {
				loggers.Info("dbw close success")
			}

			if err := dbRepo.DbRClose(); err != nil {
				loggers.Error("dbr close err", zap.Error(err))
			} else {
				loggers.Info("dbr close success")
			}
		},
	)
}
