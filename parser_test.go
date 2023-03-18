package cursor

import (
	"bytes"
	"fmt"
	"strconv"
	"testing"
)

func TestParser(t *testing.T) {
	code := []byte(

		`data: "<|BEGIN_type|>"
data: "\n"
data: "continue"
data: "\n"
data: "<|END_type|>"
data: "\n"
data: "<|BEGIN_message|>"
data: "package"
data: " main"
data: "\n\n"
data: "import"
data: " (\n"
data: "\t"
data: "\"fmt"
data: "\"\n"
data: "\t"
data: "\"sync"
data: "\"\n"
data: ")\n\n"
data: "func"
data: " main"
data: "()"
data: " {\n"
data: "\tvar"
data: " wg"
data: " sync"
data: ".Wait"
data: "Group"
data: "\n"
data: "\tfor"
data: " i"
data: " :="
data: " "
data: "0"
data: ";"
data: " i"
data: " < "
data: "5"
data: ";"
data: " i"
data: "++"
data: " {\n"
data: "\t"
data: "\twg"
data: ".Add"
data: "("
data: "1"
data: ")\n"
data: "\t"
data: "\tgo"
data: " func"
data: "(num"
data: " int"
data: ")"
data: " {\n"
data: "\t\t"
data: "\tdefer"
data: " wg"
data: ".Done"
data: "()\n"
data: "\t\t"
data: "\tfmt"
data: ".Printf"
data: "(\""
data: "G"
data: "or"
data: "outine"
data: " %"
data: "d"
data: "\\n"
data: "\","
data: " num"
data: ")\n"
data: "\t"
data: "\t"
data: "}("
data: "i"
data: ")\n"
data: "\t"
data: "}\n"
data: "\twg"
data: ".Wait"
data: "()\n"
data: "\tfmt"
data: ".Println"
data: "(\""
data: "All"
data: " gor"
data: "outines"
data: " finished"
data: " executing\")\n}\n"
data: "<|END_message|>"
data: [DONE]`)
	reader := bytes.NewBuffer(code)
	parse := Parse(reader)
	fmt.Println(string(parse))
}

func TestTakePart(t *testing.T) {
	part := takePart(`data: "\tdefer"`)
	unquote, _ := strconv.Unquote(part[0])
	fmt.Println(unquote)
}
