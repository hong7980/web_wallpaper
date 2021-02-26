package webdav

import (
	"fmt"
	"golang.org/x/net/webdav"
	"net/http"

	"webdav/srv/config"
)

var (
	bingFs = &webdav.Handler{
		FileSystem: webdav.Dir(config.ServerCfg.ImgPath + "/bing"),
		LockSystem: webdav.NewMemLS(),
	}
)

func WebDav(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println(r.Method)
	fmt.Println(r.Header)
	switch r.Method {
	case "PUT", "DELETE", "PROPPATCH", "MKCOL", "COPY", "MOVE":
		http.Error(w, "WebDAV: Read Only!!!", http.StatusForbidden)
		return
	}
	bingFs.ServeHTTP(w, r)
	next(w, r)
}
