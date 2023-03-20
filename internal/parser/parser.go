package parser

import (
	"bufio"
	"bytes"
	"io"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

var (
	codeMatchRegexp = regexp.MustCompile(`"(.+)"`)
)

func takeCodePart(raw string) []string {
	return codeMatchRegexp.FindAllString(raw, -1)
}
func Parse(reader io.Reader) []byte {
	data := new(bytes.Buffer)
	replacer := strings.NewReplacer("data: ", "")
	messageHasBegin := false
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		content := replacer.Replace(scanner.Text())
		if strings.Contains(content, "<|BEGIN_message|>") {
			messageHasBegin = true
			continue
		}
		if strings.Contains(content, "<|END_message|>") || strings.Contains(content, "[DONE]") {
			break
		}
		if !messageHasBegin {
			continue
		}
		parsedData := takeCodePart(content)
		if len(parsedData) == 0 {
			return nil
		}
		code, _ := strconv.Unquote(parsedData[0])
		r, size := utf8.DecodeRuneInString(code)
		if size == 1 {
			data.WriteString(code)
			continue
		}
		if r < utf8.RuneSelf {
			data.WriteByte(byte(r))
			continue
		}
		data.WriteRune(r)

	}
	return data.Bytes()
}
