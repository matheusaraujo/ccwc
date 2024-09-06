package main

import (
	"os"
	"strconv"
)

func wc(countBytes bool, filePath string) (string, error) {

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return "", err
	}

	var fileSize int64

	if countBytes {
		fileSize = fileInfo.Size()
	} else {
		fileSize = 0
	}

	return strconv.FormatInt(fileSize, 10) + " " + filePath, nil

}
