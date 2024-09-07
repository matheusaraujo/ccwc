package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Options struct {
	CountBytes bool
	CountWords bool
	CountLines bool
	CountChars bool
}

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

	content, err := io.ReadAll(file)
	if err != nil {
		return 0, err
	}

	charCount := int64(utf8.RuneCount(content))

	return charCount, nil
}

func wc(options Options, filePath string) (string, error) {
	var result string

	if options.CountLines {
		lines, err := countLines(filePath)
		if err != nil {
			return "", err
		}
		result += strconv.FormatInt(lines, 10) + " "
	}

	if options.CountWords {
		words, err := countWords(filePath)
		if err != nil {
			return "", err
		}
		result += strconv.FormatInt(words, 10) + " "
	}

	if options.CountBytes {
		fileSize, err := countBytes(filePath)
		if err != nil {
			return "", err
		}
		result += strconv.FormatInt(fileSize, 10) + " "
	}

	if options.CountChars {
		characters, err := countChars(filePath)
		if err != nil {
			return "", err
		}
		result += strconv.FormatInt(characters, 10) + " "
	}

	result = "   " + result + filePath

	return result, nil
}
