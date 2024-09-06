package main

import (
	"os"
	"strconv"
)

func countBytes(filePath string) (int64, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}

	return fileInfo.Size(), nil
}

func wc(flagCountBytes bool, filePath string) (string, error) {

	if flagCountBytes {
		fileSize, err := countBytes(filePath)
		if err != nil {
			return "", err
		}
		return strconv.FormatInt(fileSize, 10) + " " + filePath, nil
	}

	return "", nil

}
