package router

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"webdav/srv/handler/downloadpicture"
)

const (
	downloadBingPicRoute = "/bing"
)

// DWNPICRouter builds a download picture router
func DWNPICRouter() *mux.Router {
	router := mux.NewRouter()

	// 路由
	router.PathPrefix(downloadBingPicRoute).
		Handler(negroni.New(
			negroni.HandlerFunc(downloadpicture.GetBingPic)))

	return router
}
