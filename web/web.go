package web

import (
	"dirInfo/process"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DirInfo(c *gin.Context) {
	path := c.Query("path")
	blob, err := process.GetDir(path)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	dirInfo, err := process.Calculate(blob)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, dirInfo)
}
