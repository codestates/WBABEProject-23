package main

import (
	"context"
	"fmt"
	cf "lecture/WBABEProject-23/config"
	"lecture/WBABEProject-23/controller"
	"lecture/WBABEProject-23/logger"
	"lecture/WBABEProject-23/model"
	"lecture/WBABEProject-23/route"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

func main() {
	config := cf.GetConfig("./config/config.toml")

	if err := logger.InitLogger(config); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	logger.Debug("ready server....")
	//model 모듈 선언
	if mod, err := model.NewModel(config); err != nil {
		panic(err)
	} else if ctl, err := controller.NewController(mod); err != nil {
		panic(err)
	} else if router, err := route.NewRouter(ctl); err != nil {
		panic(err)
	} else {
		mapi := &http.Server{
			Addr:           config.Server.Host,
			Handler:        router.Index(),
			ReadTimeout:    5 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}

		g.Go(func() error {
			return mapi.ListenAndServe()
		})
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		logger.Warn("Shutdown Server ...")

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		if err := mapi.Shutdown(ctx); err != nil {
			logger.Error("Server Shutdown:", err)
		}

		select {
		case <-ctx.Done():
			logger.Info("timeout of 1 seconds.")
		}

		logger.Info("Server exiting")
	}
	if err := g.Wait(); err != nil {
		logger.Error(err)
	}

}
