package webdav

import (
	"fmt"
	"golang.org/x/net/webdav"
	"net/http"
)

var (
	Fs = &webdav.Handler{
		FileSystem: webdav.Dir("./img"),
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
	Fs.ServeHTTP(w, r)
	next(w, r)
}
