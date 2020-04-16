package process

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var client *http.Client

type dirInfo struct {
	Path      string
	DirCount  int
	FileCount int
	TotalSize int
}

type (
	result struct {
		Path string
		Dirs []dirs
	}

	dirs struct {
		Name  string
		IsDir bool
		Size  int
	}
)

var DirService string

func init() {
	client = new(http.Client)
}

func GetDir(path string) ([]byte, error) {
	address := fmt.Sprintf("%s?path=%s", DirService, path)
	resp, err := client.Get(address)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}
	blob, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return blob, nil
}

func Calculate(blob []byte) (dirInfo, error) {
	var dir result
	var dirInfo dirInfo
	err := json.Unmarshal(blob, &dir)
	if err != nil {
		return dirInfo, err
	}
	for _, d := range dir.Dirs {
		if d.IsDir {
			dirInfo.DirCount++
		} else {
			dirInfo.FileCount++
			dirInfo.TotalSize += d.Size
		}
	}
	dirInfo.Path = dir.Path
	return dirInfo, nil
}
