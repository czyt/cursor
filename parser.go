package cursor

import (
	"bufio"
	"bytes"
	"io"
	"regexp"
	"strconv"
	"strings"
)

func takePart(raw string) []string {
	re := regexp.MustCompile(`"(.+)"`)
	return re.FindAllString(raw, -1)
}
func Parse(reader io.Reader) []byte {
	buffer := bytes.Buffer{}
	replacer := strings.NewReplacer("data: ", "")
	messageHasBegin := false
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		data := replacer.Replace(scanner.Text())
		if strings.Contains(data, "<|BEGIN_message|>") {
			messageHasBegin = true
			continue
		}
		if strings.Contains(data, "<|END_message|>") {
			break
		}
		if messageHasBegin {
			part := takePart(data)
			unquote, _ := strconv.Unquote(part[0])
			buffer.WriteString(unquote)
		}
	}
	return buffer.Bytes()
}
