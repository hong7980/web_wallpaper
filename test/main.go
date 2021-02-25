package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"golang.org/x/net/webdav"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var i int

func main() {
	//i = 0
	//c := cron.New()
	//spec := "*/* * 1 * * ?"
	//c.AddFunc(spec, getPicture)
	//c.Start()

	var addr *string
	var path *string
	addr = flag.String("addr", ":8001", "") // listen端口，默认8080
	path = flag.String("path", "./img", "") // 文件路径，默认当前目录
	flag.Parse()
	//fmt.Println("addr=", *addr, ", path=", *path) // 在控制台输出配置
	http.ListenAndServe(*addr, &webdav.Handler{
		FileSystem: webdav.Dir(*path),
		LockSystem: webdav.NewMemLS(),
	})
}

func getPicture() {
	// 获取当前时间戳
	timeStr := strconv.FormatInt(time.Now().UnixNano(), 10)
	//fmt.Println(timeStr)

	// 请求获取图片链接
	r, err := http.Get("http://cn.bing.com/HPImageArchive.aspx?format=js&idx=" + strconv.Itoa(i) + "&n=1&nc=" + timeStr + "&pid=hp&FORM=BEHPTB")
	if err != nil {
		panic(err)
	}
	defer func() { _ = r.Body.Close() }()
	body, _ := ioutil.ReadAll(r.Body)
	//fmt.Printf("%s", body)

	// 从响应体中解析图片url
	var obj map[string]interface{}
	jsoniter.Unmarshal(body, &obj)
	imagesUrl := (obj["images"].([]interface{}))[0].(map[string]interface{})
	fmt.Println(imagesUrl["url"])

	// 发起Get请求下载图片
	r2, err2 := http.Get("http://cn.bing.com" + imagesUrl["url"].(string))
	if err2 != nil {
		panic(err2)
	}
	defer func() { _ = r2.Body.Close() }()
	body2, _ := ioutil.ReadAll(r2.Body)
	//fmt.Printf("%s", body2)

	// md5加密图片用作文件唯一标识
	h := md5.New()
	h.Write(body2)
	cipherStr := h.Sum(nil)
	fmt.Printf("%s\n", hex.EncodeToString(cipherStr)) // 输出加密结果
	_ = ioutil.WriteFile("./img/"+hex.EncodeToString(cipherStr)+".jpg", body2, 0755)
}
