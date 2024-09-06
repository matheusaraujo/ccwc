package main

import (
	"os"
	"strconv"
)

func wc(filePath string) (string, error) {

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return "", err
	}

	fileSize := fileInfo.Size()

	return strconv.FormatInt(fileSize, 10) + " " + filePath, nil

}
