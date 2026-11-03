package util

import (
	"bufio"
	"bytes"
	"fmt"
	"iter"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strconv"
	"strings"
)

func filename(dayPart string, isSample bool) string {
	yearDir := detectYearDir()
	folder := filepath.Join(yearDir, "inputs")
	if isSample {
		folder = filepath.Join(yearDir, "samples")
	}

	_, fileName, _, ok := runtime.Caller(0)
	if !ok {
		return filepath.Join(folder, fmt.Sprintf("%s.txt", dayPart))
	}
	prefixPath := filepath.Dir(fileName)
	basePath := filepath.Join(prefixPath, "..", folder)

	return filepath.Join(basePath, fmt.Sprintf("%s.txt", dayPart))
}

func detectYearDir() string {
	const defaultYear = "2024"
	for skip := 2; skip < 10; skip++ {
		if _, fileName, _, ok := runtime.Caller(skip); ok {
			if year := extractYear(fileName); year != "" {
				return year
			}
		} else {
			break
		}
	}
	return defaultYear
}

func extractYear(path string) string {
	parts := strings.Split(filepath.ToSlash(path), "/")
	for i, part := range slices.Backward(parts) {
		if part == "days" && i > 0 {
			candidate := parts[i-1]
			if len(candidate) == 4 {
				if _, err := strconv.Atoi(candidate); err == nil {
					return candidate
				}
			}
		}
	}
	return ""
}

func ReadBytes(dayPart string, isSample bool) (iter.Seq[byte], error) {
	return readSeq(filename(dayPart, isSample), byteSplit, asByte)
}

func ReadStrings(dayPart string, isSample bool, delim byte) (iter.Seq[string], error) {
	return readSeq(filename(dayPart, isSample), delimSplit(delim), asString)
}

func ReadIntLists(dayPart string, isSample bool, delim byte) (iter.Seq[[]int], error) {
	return readSeq(filename(dayPart, isSample), delimSplit(delim), asIntList)
}

func ReadInts(dayPart string, isSample bool, delim byte) (iter.Seq[int], error) {
	return readSeq(filename(dayPart, isSample), delimSplit(delim), asInt)
}

func asByte(data []byte) (byte, error) {
	return data[0], nil
}

func asString(data []byte) (string, error) {
	return string(data), nil
}

func asInt(data []byte) (int, error) {
	return strconv.Atoi(string(data))
}

func asIntList(data []byte) ([]int, error) {
	var list []int
	for number := range strings.SplitSeq(string(data), ",") {
		value, err := strconv.Atoi(number)
		if err != nil {
			return nil, err
		}
		list = append(list, value)
	}

	return list, nil
}

type convertFunc[T any] func(data []byte) (T, error)

// delimSplit splits on each occurrence of delim, dropping the delimiters themselves.
func delimSplit(delim byte) bufio.SplitFunc {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		if i := bytes.IndexByte(data, delim); i >= 0 {
			// Advance past the token and the delimiter character (+1)
			return i + 1, data[0:i], nil
		}
		if atEOF {
			return len(data), data, nil
		}
		return 0, nil, nil
	}
}

// byteSplit yields one single-byte token at a time, streaming the whole file.
func byteSplit(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if len(data) == 0 {
		return 0, nil, nil
	}

	return 1, data[0:1], nil
}

func readSeq[T any](file string, split bufio.SplitFunc, conv convertFunc[T]) (iter.Seq[T], error) {
	fp, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(fp)
	scanner.Split(split)

	return func(yield func(T) bool) {
		defer fp.Close()
		for scanner.Scan() {
			value, err := conv(scanner.Bytes())
			if err != nil {
				panic(err)
			}
			if !yield(value) {
				return
			}
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}, nil
}
