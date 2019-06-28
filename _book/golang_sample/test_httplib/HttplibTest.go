package main

import (
	"fmt"

	//	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
)

var log *logs.BeeLogger

//httplib Get
func init() {
	log = logs.NewLogger(1000)
	log.SetLogger("console", "")
}

func HttplibGet(str string) {
	req := httplib.Get(str)
	str, err := req.String()
	if err != nil {
		log.Error((err.Error()))
	}
	fmt.Println(str)
}
func main() {
	HttplibGet("http://api.heclouds.com/devices/530698847")
}
