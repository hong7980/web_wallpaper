package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/jasonlvhit/gocron"
	"github.com/urfave/negroni"

	"webdav/srv/config"
	"webdav/srv/dwnpic/bing"
	"webdav/srv/router"
)

func init() {
	//fmt.Println(os.Getwd())
	err := config.LoadCfg()
	if err != nil {
		fmt.Println("load cfg failed: " + err.Error())
		os.Exit(-1)
	}
}

func main() {
	// 抓取图片
	s := gocron.NewScheduler()
	s.Every(10).Hours().Do(bing.GetBingPicture)
	s.Start()

	// webdav服务路由
	wdvRouter := router.WDVRouter()
	wdvHandler := negroni.New()
	wdvHandler.Use(negroni.NewRecovery())
	wdvHandler.UseHandler(wdvRouter)
	addr := ":" + config.ServerCfg.WebdavPort

	webdavServer := &http.Server{Addr: addr, Handler: wdvHandler}

	// zhuaquluyou
	dwnpicRouter := router.DWNPICRouter()
	dwnpicHandler := negroni.New()
	dwnpicHandler.Use(negroni.NewRecovery())
	dwnpicHandler.UseHandler(dwnpicRouter)
	dwnpicAddr := ":" + config.ServerCfg.DwnpicsrvPort

	dwnpicServer := &http.Server{Addr: dwnpicAddr, Handler: dwnpicHandler}

	// 启动所有服务
	err := gracehttp.Serve(
		webdavServer,
		dwnpicServer,
	)

	if err != nil {
		fmt.Println(err)
		return
	}
}
