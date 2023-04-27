package gotool

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"os"
)

const (
	PermFile fs.FileMode = 0666
	PermDir              = 0755
)

// file is exists
// 判断地址是否存在
func FileExists(filepath string) bool {
	info, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// read file content to map instance
func FileReadToMap(path string) (map[string]interface{}, error) {
	var fileData map[string]interface{}
	err := FileReadForInterface(path, &fileData)
	if err != nil {
		return nil, err
	}
	return fileData, nil
}

func FileReadForInterface(path string, i interface{}) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, i)
}

func FileWriteWithInterface(path string, data interface{}) error {
	var writeBytes []byte
	var err error
	switch data.(type) {
	case string:
		writeBytes = []byte(data.(string))
	case []byte:
		writeBytes = data.([]byte)
	default:
		writeBytes, err = json.Marshal(data)
		if err != nil {
			return err
		}
	}
	return ioutil.WriteFile(path, writeBytes, PermFile)
}
