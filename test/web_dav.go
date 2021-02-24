package main

import (
	"net/http"

	"golang.org/x/net/webdav"
)

func main() {
	http.HandleFunc("/", webDav)
	http.ListenAndServe(":8080", nil)
}

func webDav(w http.ResponseWriter, req *http.Request) {
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
