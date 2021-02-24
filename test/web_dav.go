package main

import (
	"golang.org/x/net/webdav"
	"net/http"
)

//var (
//	fs = &webdav.Handler{
//		FileSystem: webdav.Dir("../img"),
//		LockSystem: webdav.NewMemLS(),
//	}
//)
//
//func main() {
//	http.HandleFunc("/", webDav)
//	http.ListenAndServe(":8089", nil)
//	fmt.Println("start")
//}
//
//func webDav(w http.ResponseWriter, req *http.Request) {
//	switch req.Method {
//	case "PUT", "DELETE", "PROPPATCH", "MKCOL", "COPY", "MOVE":
//		http.Error(w, "WebDAV: Read Only!!!", http.StatusForbidden)
//		return
//	}
//	//fs = &webdav.Handler{
//	//	FileSystem: webdav.Dir("../img"),
//	//	LockSystem: webdav.NewMemLS(),
//	//}
//	fmt.Println("1")
//	fs.ServeHTTP(w, req)
//}

func main() {
	fs := &webdav.Handler{
		FileSystem: webdav.Dir("./img"),
		LockSystem: webdav.NewMemLS(),
	}
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "PUT", "DELETE", "PROPPATCH", "MKCOL", "COPY", "MOVE":
			http.Error(w, "WebDAV: Read Only!!!", http.StatusForbidden)
			return
		}
		fs.ServeHTTP(w, req)
	})
	http.ListenAndServe(":8089", nil)
}
