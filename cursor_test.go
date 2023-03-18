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

func TestClient_Conversation(t *testing.T) {
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
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}

func TestClient_ConversationChat(t *testing.T) {
	cli := NewClient()
	response, err := cli.Conversation(ConversationRequest{
		UserRequest: UserRequest{
			Message:              "explain this regexp",
			CurrentFileName:      "gopher.go",
			CurrentSelection:     "re := regexp.MustCompile(`data: \\\" executing\\\\\\\\\\\\\\\"\\\\)\\\\\\\\\\\\\\\\n\\\\}\\\\\\\\\\\\\\\\n\\\"`)\\n",
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
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}
