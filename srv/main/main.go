package main

import (
	"fmt"
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/urfave/negroni"
	"webdav/srv/router"
)

func main() {
	// 抓取图片
	//go bing.GetBingPicture()

	// webdav服务路由
	wdvRouter := router.WDVRouter()
	wdvHandler := negroni.New()
	wdvHandler.Use(negroni.NewRecovery())
	wdvHandler.UseHandler(wdvRouter)

	webdavServer := &http.Server{Addr: ":8888", Handler: wdvHandler}
	// 启动所有服务
	err := gracehttp.Serve(
		webdavServer,
	)

	if err != nil {
		fmt.Println(err)
		return
	}
}
