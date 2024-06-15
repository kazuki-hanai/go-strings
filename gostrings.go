package gostrings

import (
	"bufio"
	"io"
)

func isPrintable(b byte) bool {
	return b >= 32 && b <= 126
}

// Read strings from reader and return them as a slice of strings
func GetStrings(f io.Reader, minWordCount, bufsize int) ([]string, error) {
	buf := bufio.NewReader(f)

	result := make([]string, 0)
	str := make([]byte, 0)

	for b, err := buf.ReadByte(); ; b, err = buf.ReadByte() {
		if err == io.EOF {
			break
		}

		if isPrintable(b) {
			str = append(str, b)
		} else {
			if minWordCount <= len(str) {
				result = append(result, string(str))
			}
			str = make([]byte, 0)
		}
	}

	return result, nil
}
