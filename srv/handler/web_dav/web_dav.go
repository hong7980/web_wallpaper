package web_dav

import (
	"net/http"

	"golang.org/x/net/webdav"
)

func WebDav(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	switch req.Method {
	case "PUT", "DELETE", "PROPPATCH", "MKCOL", "COPY", "MOVE":
		http.Error(w, "WebDAV: Read Only!!!", http.StatusForbidden)
		return
	}
	fs := &webdav.Handler{
		FileSystem: webdav.Dir("."),
		LockSystem: webdav.NewMemLS(),
	}
	fs.ServeHTTP(w, req)
}
