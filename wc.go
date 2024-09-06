package main

import (
	"bufio"
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

func countLines(filePath string) (int64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lineCount int64 = 0

	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lineCount, nil
}

func wc(shouldCountBytes bool, shouldCountLines bool, filePath string) (string, error) {
	var result string

	if shouldCountLines {
		lines, err := countLines(filePath)
		if err != nil {
			return "", err
		}
		result += strconv.FormatInt(lines, 10) + " "
	}

	if shouldCountBytes {
		fileSize, err := countBytes(filePath)
		if err != nil {
			return "", err
		}
		result += strconv.FormatInt(fileSize, 10) + " "
	}

	result += filePath

	return result, nil
}
