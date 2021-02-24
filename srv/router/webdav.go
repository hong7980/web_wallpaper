package router

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"webdav/srv/handler/web_dav"
)

const (
	WDVRoute = "/"
)

// WDVRouter builds a web_dav router
func WDVRouter() *mux.Router {
	router := mux.NewRouter()

	// 路由
	router.Path(WDVRoute).
		Handler(negroni.New(
			negroni.HandlerFunc(web_dav.WebDav))).
		Name("webdav").Methods("GET")

	return router
}
