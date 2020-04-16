package main

import (
	"dirInfo/process"
	"dirInfo/web"
	"flag"

	"github.com/gin-gonic/gin"
)

func init() {
	flag.StringVar(&process.DirService, "dirService", "http://192.168.154.132:9000/v1/directory", "the dirService address")
}
func main() {
	flag.Parse()
	r := gin.Default()
	r.GET("/v1/directory/infomation", web.DirInfo)
	r.Run(":9001")
}
