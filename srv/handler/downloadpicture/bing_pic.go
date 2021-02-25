package downloadpicture

import (
	"fmt"
	"net/http"
	"webdav/srv/dwnpic/bing"
)

func GetBingPic(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	bing.GetBingPicture()
	fmt.Fprintln(w, "update bing picture successful")
	next(w, r)
}
