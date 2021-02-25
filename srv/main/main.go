package main

import (
	"fmt"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/urfave/negroni"
	"net/http"
	"webdav/srv/dwnpic/bing"
	"webdav/srv/wdv_test"
)

func main() {
	// 抓取图片
	go bing.GetBingPicture()

	// webdav服务路由
	//wdvRouter := router.WDVRouter()
	//wdvHandler := negroni.New()
	//wdvHandler.Use(negroni.NewRecovery())
	//wdvHandler.UseHandler(wdvRouter)

	//webdavServer := &http.Server{Addr: ":80", Handler: wdvHandler}

	wdvHandler := negroni.New()
	wdvHandler.Use(negroni.NewRecovery())
	wdvHandler.UseHandlerFunc(wdv_test.Dav)

	webdavServer := &http.Server{Addr: ":80", Handler: wdvHandler}

	// 启动所有服务
	err := gracehttp.Serve(
		webdavServer,
	)

	if err != nil {
		fmt.Println(err)
		return
	}
}
