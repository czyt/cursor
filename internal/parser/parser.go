package parser

import (
	"bytes"
	"io"
	"regexp"
	"strconv"
)

var (
	codeMatchRegexp = regexp.MustCompile(`data: (".+")`)
)

func Parse(reader io.Reader) ([]byte, error) {
	data := new(bytes.Buffer)
	content, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	matchedContent := codeMatchRegexp.FindAllStringSubmatch(string(content), -1)
	for _, s := range matchedContent {

		unquote, err := strconv.Unquote(s[1])
		if err != nil {
			continue
		}
		data.WriteString(unquote)
	}
	//data = bytes.NewBufferString(unquote)
	return data.Bytes(), nil
}
