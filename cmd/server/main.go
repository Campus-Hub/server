package main

import (
	"errors"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Campus-Hub/server/internal/model"

	"github.com/Campus-Hub/server/config"
	_ "github.com/Campus-Hub/server/internal/model"
	"github.com/Campus-Hub/server/pkg/logger"
	"github.com/Campus-Hub/server/router"
	"github.com/golang/glog"
)

// Setup each System before server.
func init() {
	logger.Setup()
	config.Setup()
	model.Setup()
	// cache.Setup()
}

func main() {
	flag.Parse()
	defer glog.Flush()
	go startListen() //转载路由
	{
		osSignals := make(chan os.Signal, 1)
		// 开启信号监听
		signal.Notify(osSignals, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
		s := <-osSignals
		glog.Info("exit! ", s)
	}
	glog.Flush()
}

func startListen() {
	//gin.SetMode(gin.ReleaseMode)
	logger.Logger.Error("UploadResourceService", errors.New("test"))
	server := http.Server{
		Addr:    ":4000",
		Handler: router.Setup(),
	}
	err := server.ListenAndServe()
	if err != nil {
		println("")
		//logger.Logger.Infoln("绑定HTTP到 %s 失败！可能是端口已经被占用，或用户权限不足", config.Config.AppName)
		logger.Logger.Infoln(err)
	}
}
