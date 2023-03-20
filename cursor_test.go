package cursor

import "testing"

func TestClient_Tests(t *testing.T) {
	cli := NewClient()
	response, err := cli.Tests(TestsRequest{
		CachedTests:      []any{},
		FileName:         "gopher.go",
		FileContents:     "",
		TestFileName:     "",
		TestFileContents: "",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}

func TestClient_Comment(t *testing.T) {
	cli := NewClient()
	response, err := cli.Comment(CommentRequest{
		Filename: "gopher.go",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}

func TestClient_ConversationGenerate(t *testing.T) {
	cli := NewClient()
	response, err := cli.Conversation(ConversationRequest{
		UserRequest: UserRequest{
			Message:              "golang waitgroup Example",
			CurrentFileName:      "gopher.go",
			CurrentRootPath:      "C:\\Users\\czyt\\cursor-tutor",
			PrecedingCode:        []any{},
			SuffixCode:           []any{},
			CopilotCodeBlocks:    []any{},
			CustomCodeBlocks:     []any{},
			CodeBlockIdentifiers: []any{},
			MsgType:              "generate",
		},
		UserMessages: []any{},
		BotMessages:  []any{},
		ContextType:  "copilot",
		RootPath:     "C:\\Users\\czyt\\cursor-tutor",
	}, "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}

func TestClient_ConversationChat(t *testing.T) {
	cli := NewClient()
	response, err := cli.Conversation(ConversationRequest{
		UserRequest: UserRequest{
			Message:              "go语言发布时间",
			CurrentFileName:      "gopher.go",
			CurrentFileContents:  "^1[3-9]\\\\d{9}$\\n",
			CurrentSelection:     "^1[3-9]\\\\d{9}$\\n",
			CurrentRootPath:      "C:\\Users\\czyt\\cursor-tutor",
			PrecedingCode:        []any{},
			SuffixCode:           []any{},
			CopilotCodeBlocks:    []any{},
			CustomCodeBlocks:     []any{},
			CodeBlockIdentifiers: []any{},
			MsgType:              "edit",
		},
		UserMessages: []any{},
		BotMessages:  []any{},
		ContextType:  "copilot",
		RootPath:     "C:\\Users\\czyt\\cursor-tutor",
	}, "zh-CN")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}

// Todo: add message type :freeform
//{
//	"userRequest": {
//		"message": "如何优化这个程序",
//		"currentRootPath": "C:\\Users\\czyt\\cursor-tutor",
//		"currentFileName": "C:\\Users\\czyt\\cursor-tutor\\gopher.go",
//		"currentFileContents": "package main\n\nimport \"fmt\"\n\ntype C chan C\n\nfunc main() {\n  var c = make(C, 1)\n  c <- c\n  for i := 0; ; i++ {\n    select {\n    case <-c:\n    case <-c:\n      c <- c\n    default:\n      fmt.Println(i)\n      return\n    }\n  }\n}",
//		"precedingCode": [],
//		"currentSelection": "package main\n\nimport \"fmt\"\n\ntype C chan C\n\nfunc main() {\n  var c = make(C, 1)\n  c <- c\n  for i := 0; ; i++ {\n    select {\n    case <-c:\n    case <-c:\n      c <- c\n    default:\n      fmt.Println(i)\n      return\n    }\n  }\n}",
//		"suffixCode": [],
//		"copilotCodeBlocks": [],
//		"customCodeBlocks": [],
//		"codeBlockIdentifiers": [],
//		"msgType": "freeform",
//		"maxOrigLine": null
//	},
//	"userMessages": [],
//	"botMessages": [],
//	"contextType": "copilot",
//	"rootPath": "C:\\Users\\czyt\\cursor-tutor"
//}