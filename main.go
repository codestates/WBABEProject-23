package main

import (
	cf "lecture/WBABEProject-23/config"
	"lecture/WBABEProject-23/controller"
	"lecture/WBABEProject-23/model"
	"lecture/WBABEProject-23/route"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

func main() {
	config := cf.GetConfig("./config/config.toml")

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
	}
}
