package process

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync/atomic"
)

var client *http.Client

type DirInfos struct {
	Path      string
	DirCount  int64
	FileCount int64
	TotalSize int64
}

type (
	result struct {
		Path string
		Dirs []dirs
	}

	dirs struct {
		Name  string
		IsDir bool
		Size  int64
	}
)

var DirService string

func init() {
	client = new(http.Client)
}

func GetDirInfo(path string, dirInfo DirInfos, ch chan int) (DirInfos, error) {
	ch <- 1
	var dir result
	blob, err := getDir(path)
	if err != nil {
		return dirInfo, err
	}
	err = json.Unmarshal(blob, &dir)
	if err != nil {
		return dirInfo, err
	}
	for _, d := range dir.Dirs {
		if d.IsDir {
			<-ch
			go GetDirInfo(d.Name, dirInfo, ch)
			newDirCount := atomic.AddInt64(&dirInfo.DirCount, 1)
			dirInfo.DirCount = newDirCount
		} else {
			newFileCount := atomic.AddInt64(&dirInfo.FileCount, 1)
			dirInfo.FileCount = newFileCount
			newTotalCount := atomic.AddInt64(&dirInfo.TotalSize, d.Size)
			dirInfo.TotalSize = newTotalCount
		}
	}
	return dirInfo, nil
}

func getDir(path string) ([]byte, error) {
	address := fmt.Sprintf("%s?path=%s", DirService, path)
	resp, err := client.Get(address)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	blob, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return blob, nil
}
