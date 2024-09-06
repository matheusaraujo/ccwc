package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
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

func countWords(filePath string) (int64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var wordCount int64 = 0

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		wordCount += int64(len(words))
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return wordCount, nil
}

func countChars(filePath string) (int64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var charCount int64 = 0

	for scanner.Scan() {
		line := scanner.Text()
		charCount += int64(utf8.RuneCountInString(line)) + 2
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return charCount, nil
}

func wc(shouldCountBytes bool, shouldCountWords bool, shouldCountLines bool, shouldCountChars bool, filePath string) (string, error) {
	var result string

	if shouldCountLines {
		lines, err := countLines(filePath)
		if err != nil {
			return "", err
		}
		result += strconv.FormatInt(lines, 10) + " "
	}

	if shouldCountWords {
		words, err := countWords(filePath)
		if err != nil {
			return "", err
		}
		result += strconv.FormatInt(words, 10) + " "
	}

	if shouldCountBytes {
		fileSize, err := countBytes(filePath)
		if err != nil {
			return "", err
		}
		result += strconv.FormatInt(fileSize, 10) + " "
	}

	if shouldCountChars {
		characters, err := countChars(filePath)
		if err != nil {
			return "", err
		}
		result += strconv.FormatInt(characters, 10) + " "
	}

	result = "   " + result + filePath

	return result, nil
}
