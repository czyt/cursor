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
		UserMessages: make([]UserMessage, 0),
		BotMessages:  make([]BotMessage, 0),
		ContextType:  "copilot",
		RootPath:     "C:\\Users\\czyt\\cursor-tutor",
	}, "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}

// a demo json request
//
//	{
//		"userRequest": {
//			"message": "这个有什么优点么",
//			"currentRootPath": "C:\\Users\\czyt\\cursor-tutor",
//			"currentFileName": "C:\\Users\\czyt\\cursor-tutor\\gopher.go",
//			"currentFileContents": "package main\n\nimport \"fmt\"\n\ntype C chan C\n\nfunc main() {\n  var c = make(C, 1)\n  c <- c\n  for i := 0; ; i++ {\n    select {\n    case <-c:\n    case <-c:\n      c <- c\n    default:\n      fmt.Println(i)\n      return\n    }\n  }\n}",
//			"precedingCode": [],
//			"currentSelection": "package main\n\nimport \"fmt\"\n\ntype C chan C\n\nfunc main() {\n  var c = make(C, 1)\n  c <- c\n  for i := 0; ; i++ {\n    select {\n    case <-c:\n    case <-c:\n      c <- c\n    default:\n      fmt.Println(i)\n      return\n    }\n  }\n}",
//			"suffixCode": [],
//			"copilotCodeBlocks": [],
//			"customCodeBlocks": [],
//			"codeBlockIdentifiers": [],
//			"msgType": "freeform",
//			"maxOrigLine": null
//		},
//		"userMessages": [{
//			"sender": "user",
//			"sentAt": 1679398840144,
//			"message": "使用go实现双缓存队列",
//			"conversationId": "832b0c33-b12a-4ae4-aa1c-15f69409cc81",
//			"otherCodeBlocks": [],
//			"codeSymbols": [],
//			"currentFile": "C:\\Users\\czyt\\cursor-tutor\\gopher.go",
//			"precedingCode": "",
//			"procedingCode": "",
//			"currentSelection": "package main\n\nimport \"fmt\"\n\ntype C chan C\n\nfunc main() {\n  var c = make(C, 1)\n  c <- c\n  for i := 0; ; i++ {\n    select {\n    case <-c:\n    case <-c:\n      c <- c\n    default:\n      fmt.Println(i)\n      return\n    }\n  }\n}",
//			"selection": {
//				"from": 0,
//				"to": 221
//			},
//			"msgType": "freeform"
//		}],
//		"botMessages": [{
//			"sender": "bot",
//			"sentAt": 1679398840405,
//			"conversationId": "832b0c33-b12a-4ae4-aa1c-15f69409cc81",
//			"type": "markdown",
//			"message": "\nThe code you provided is an implementation of a program that prints the number of iterations it takes to exit an infinite loop. The program uses a channel of channels to create a cyclic dependency, which is then used to create a deadlock. \n\nAs for your question, to implement a double-buffered queue in Go, you can use the built-in package \"container/ring\". Here's an example implementation:\n\n```go\npackage main\n\nimport (\n    \"container/ring\"\n    \"fmt\"\n)\n\nfunc main() {\n    size := 5\n    r := ring.New(size)\n\n    for i := 0; i < size; i++ {\n        r.Value = i\n        r = r.Next()\n    }\n\n    for i := 0; i < size; i++ {\n        fmt.Println(r.Value)\n        r = r.Next()\n    }\n}\n```\n\nIn this example, we create a new ring of size 5 and iterate through it to set the values. We then iterate through the ring again to print the values. You can modify this code to fit your specific use case.",
//			"lastToken": "<|END_message|>",
//			"finished": true,
//			"currentFile": "C:\\Users\\czyt\\cursor-tutor\\gopher.go",
//			"interrupted": false
//		}],
//		"contextType": "copilot",
//		"rootPath": "C:\\Users\\czyt\\cursor-tutor"
//	}
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
		UserMessages: make([]UserMessage, 0),
		BotMessages:  make([]BotMessage, 0),
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
