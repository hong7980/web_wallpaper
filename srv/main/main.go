package main

import (
	"fmt"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/jasonlvhit/gocron"
	"github.com/urfave/negroni"
	"net/http"
	"webdav/srv/dwnpic/bing"
	"webdav/srv/router"
)

func main() {
	// 抓取图片
	s := gocron.NewScheduler()
	s.Every(10).Hours().Do(bing.GetBingPicture)
	s.Start()

	//webdav服务路由
	wdvRouter := router.WDVRouter()
	wdvHandler := negroni.New()
	wdvHandler.Use(negroni.NewRecovery())
	wdvHandler.UseHandler(wdvRouter)

	webdavServer := &http.Server{Addr: ":80", Handler: wdvHandler}

	//zhuaquluyou
	dwnpicRouter := router.DWNPICRouter()
	dwnpicHandler := negroni.New()
	dwnpicHandler.Use(negroni.NewRecovery())
	dwnpicHandler.UseHandler(dwnpicRouter)

	dwnpicServer := &http.Server{Addr: ":50001", Handler: dwnpicHandler}

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
