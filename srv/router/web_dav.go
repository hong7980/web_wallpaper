package router

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"webdav/srv/handler/webdav"
)

const (
	WDVRoute = "/"
)

// WDVRouter builds a webdav router
func WDVRouter() *mux.Router {
	router := mux.NewRouter()

	// 路由
	router.PathPrefix(WDVRoute).
		Handler(negroni.New(
			negroni.HandlerFunc(webdav.WebDav)))

	return router
}
