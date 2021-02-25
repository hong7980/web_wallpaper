package web_dav

import (
	"net/http"

	"golang.org/x/net/webdav"
)

var (
	Fs = &webdav.Handler{
		FileSystem: webdav.Dir("./img"),
		LockSystem: webdav.NewMemLS(),
	}
)

func WebDav(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	switch req.Method {
	case "PUT", "DELETE", "PROPPATCH", "MKCOL", "COPY", "MOVE":
		http.Error(w, "WebDAV: Read Only!!!", http.StatusForbidden)
		next(w, req)
		return
	}
	//fmt.Println(os.Getwd())
	Fs.ServeHTTP(w, req)
	//fmt.Println("2")
	next(w, req)
}
