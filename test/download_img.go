package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type xBody struct {
	images string
}

func main() {
	getPicture2()
}

func getPicture2() {
	// huoqushijianz
	timeStr := strconv.FormatInt(time.Now().UnixNano(), 10)
	//fmt.Println(timeStr)

	r, err := http.Get("http://cn.bing.com/HPImageArchive.aspx?format=js&idx=1&n=1&nc=" + timeStr + "&pid=hp&FORM=BEHPTB")
	if err != nil {
		panic(err)
	}
	defer func() { _ = r.Body.Close() }()
	body, _ := ioutil.ReadAll(r.Body)
	//fmt.Printf("%s", body)

	// Parse the url in json
	var obj map[string]interface{}
	jsoniter.Unmarshal(body, &obj)
	imagesUrl := (obj["images"].([]interface{}))[0].(map[string]interface{})
	fmt.Println(imagesUrl["url"])

	// get images
	r2, err2 := http.Get("http://cn.bing.com" + imagesUrl["url"].(string))
	if err2 != nil {
		panic(err2)
	}
	defer func() { _ = r2.Body.Close() }()
	body2, _ := ioutil.ReadAll(r2.Body)
	//fmt.Printf("%s", body2)

	// md5
	h := md5.New()
	h.Write(body2) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	fmt.Printf("%s\n", hex.EncodeToString(cipherStr)) // 输出加密结果

	_ = ioutil.WriteFile("./img/"+hex.EncodeToString(cipherStr)+".jpg", body2, 0755)
}
