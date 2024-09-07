package main

import (
	"bufio"
	"bytes"
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

func countBytes(r io.Reader) (int64, error) {
	var size int64
	if file, ok := r.(*os.File); ok {
		fileInfo, err := file.Stat()
		if err != nil {
			return 0, err
		}
		size = fileInfo.Size()
	} else {
		content, err := io.ReadAll(r)
		if err != nil {
			return 0, err
		}
		size = int64(len(content))
	}
	return size, nil
}

func countLines(r io.Reader) (int64, error) {
	scanner := bufio.NewScanner(r)
	var lineCount int64

	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lineCount, nil
}

func countWords(r io.Reader) (int64, error) {
	scanner := bufio.NewScanner(r)
	var wordCount int64

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

func countChars(r io.Reader) (int64, error) {
	content, err := io.ReadAll(r)
	if err != nil {
		return 0, err
	}
	return int64(utf8.RuneCount(content)), nil
}

func wc(options Options, filePath *string) (string, error) {
	var r io.Reader
	var file *os.File
	var err error
	var content []byte

	if filePath == nil {
		content, err = io.ReadAll(os.Stdin)
		if err != nil {
			return "", err
		}
		r = bytes.NewReader(content)
	} else {
		file, err = os.Open(*filePath)
		if err != nil {
			return "", err
		}
		defer file.Close()
		r = file
	}

	var result string

	if options.CountLines {
		lines, err := countLines(r)
		if err != nil {
			return "", err
		}
		result += strconv.FormatInt(lines, 10) + " "
	}

	if filePath != nil {
		file.Close()
		file, err = os.Open(*filePath)
		if err != nil {
			return "", err
		}
		defer file.Close()
		r = file
	} else {
		r = bytes.NewReader(content)
	}

	if options.CountWords {
		words, err := countWords(r)
		if err != nil {
			return "", err
		}
		result += strconv.FormatInt(words, 10) + " "
	}

	if filePath != nil {
		file.Close()
		file, err = os.Open(*filePath)
		if err != nil {
			return "", err
		}
		defer file.Close()
		r = file
	} else {
		r = bytes.NewReader(content)
	}

	if options.CountBytes {
		fileSize, err := countBytes(r)
		if err != nil {
			return "", err
		}
		result += strconv.FormatInt(fileSize, 10) + " "
	}

	if filePath != nil {
		file.Close()
		file, err = os.Open(*filePath)
		if err != nil {
			return "", err
		}
		defer file.Close()
		r = file
	} else {
		r = bytes.NewReader(content)
	}

	if options.CountChars {
		characters, err := countChars(r)
		if err != nil {
			return "", err
		}
		result += strconv.FormatInt(characters, 10) + " "
	}

	if filePath == nil {
		result = "   " + result
	} else {
		result = "   " + result + *filePath
	}

	return result, nil
}
