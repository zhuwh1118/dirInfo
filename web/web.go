package web

import (
	"dirInfo/process"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DirInfo(c *gin.Context) {
	path := c.Query("path")
	dirInfo := process.DirInfos{}
	var ch = make(chan int, 5)
	dirInfo, err := process.GetDirInfo(path, dirInfo, ch)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	dirInfo.Path = path
	c.JSON(http.StatusOK, dirInfo)
	return
}
