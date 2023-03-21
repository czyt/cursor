package parser

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
)

var (
	headerData   = []byte("data: ")
	beginType    = []byte("<|BEGIN_type|>")
	endType      = []byte("<|END_type|>")
	beginMessage = []byte("<|BEGIN_message|>")
	endMessage   = []byte("<|END_message|>")
	end          = []byte("[DONE]")
)

func Parse(reader io.Reader) ([]byte, error) {

	var (
		isMessageBegin = false
		isFinished     = false
	)
	data := new(bytes.Buffer)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() && !isFinished {
		line := scanner.Bytes()
		line = bytes.TrimSpace(line)
		if !bytes.HasPrefix(line, headerData) {
			continue
		}

		line = bytes.TrimPrefix(line, headerData)

		// Message begin
		if bytes.Contains(line, beginMessage) {
			isMessageBegin = true
			continue
		}
		// message end
		if bytes.EqualFold(line, end) {
			isFinished = true
			continue
		}
		// special content skip
		if bytes.Contains(line, beginType) ||
			bytes.Contains(line, endType) ||
			bytes.Contains(line, endMessage) {
			continue
		}

		if !isFinished && isMessageBegin {
			unquote, err := strconv.Unquote(string(line))
			if err != nil {
				continue
			}

			data.WriteString(unquote)
		}
	}

	return data.Bytes(), nil
}
