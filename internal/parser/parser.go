package parser

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
)

var (
	headerData = []byte("data: ")
)

func Parse(reader io.Reader) ([]byte, error) {
	data := new(bytes.Buffer)
	isFinished := false
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() && !isFinished {
		line := scanner.Bytes()
		line = bytes.TrimSpace(line)
		if !bytes.HasPrefix(line, headerData) {
			continue
		}
		line = bytes.TrimPrefix(line, headerData)
		if string(line) == "[DONE]" {
			isFinished = true
		}
		if !isFinished {
			unquote, err := strconv.Unquote(string(line))
			if err != nil {
				continue
			}

			data.WriteString(unquote)
		}
	}

	return data.Bytes(), nil
}
