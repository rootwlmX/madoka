// Package util 工具
package util

import (
	"bufio"
	"encoding/json"
	"madoka/models"
	"os"
)

var _cfg *models.AppInfo = nil

// ParseConfig parse config
func ParseConfig(path string) (*models.AppInfo, error) {

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)

	if err = decoder.Decode(&_cfg); err != nil {
		return nil, err
	}
	return _cfg, nil
}
