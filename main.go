package main

import (
	"eicesoft/web-demo/pkg/env"
	"eicesoft/web-demo/pkg/shutdown"
	"eicesoft/web-demo/router"
	"flag"
	"math/rand"
	"time"

	"github.com/zh-five/xdaemon"
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
	if env.Get().IsDaemon() {
		xdaemon.Background("", true)
	}

	mux, err := router.InitMux()
	if err != nil {
		panic(err)
	}

	mux.StartServer()

	shutdown.NewHook().Close(
		// 关闭 http server
		func() {
			mux.Shutdown()
		},
	)
}
