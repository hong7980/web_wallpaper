package web_dav

import (
	"fmt"
	"net/http"

	"golang.org/x/net/webdav"
)

var (
	Fs = &webdav.Handler{
		FileSystem: webdav.Dir("./img"),
		LockSystem: webdav.NewMemLS(),
	}
)

func WebDav(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	switch r.Method {
	case "PUT", "DELETE", "PROPPATCH", "MKCOL", "COPY", "MOVE":
		http.Error(w, "WebDAV: Read Only!!!", http.StatusForbidden)
		return
	}
	//fmt.Println(os.Getwd())
	Fs.ServeHTTP(w, r)
	fmt.Println("2")
	next(w, r)
}
