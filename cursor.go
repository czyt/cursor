package cursor

import (
	"bytes"
	"encoding/json"
	"github.com/czyt/cursor/internal/parser"
	"net/http"
)

const (
	testsUrl        = "https://aicursor.com/tests"
	commentUrl      = "https://aicursor.com/comment"
	conversationUrl = "https://aicursor.com/conversation"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Tests(data TestsRequest) (response any, err error) {

	client := &http.Client{}
	payload, err := json.Marshal(data)
	req, err := http.NewRequest(http.MethodPost, testsUrl, bytes.NewBuffer(payload))

	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Cursor/0.1.1 Chrome/108.0.5359.62 Electron/22.0.0 Safari/537.36")

	res, err := client.Do(req)
	defer res.Body.Close()
	if err != nil {
		return
	}
	return response, nil
}

func (c *Client) Comment(comment CommentRequest) (response any, err error) {
	payload, err := json.Marshal(comment)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, commentUrl, bytes.NewBuffer(payload))
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Cursor/0.1.1 Chrome/108.0.5359.62 Electron/22.0.0 Safari/537.36")
	req.Header.Add("Sec-Fetch-Site", "cross-site")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	return
}

func (c *Client) Conversation(conversation ConversationRequest, language string) (response string, err error) {
	payload, err := json.Marshal(conversation)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, conversationUrl, bytes.NewBuffer(payload))
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Cursor/0.1.1 Chrome/108.0.5359.62 Electron/22.0.0 Safari/537.36")
	req.Header.Add("Sec-Fetch-Site", "cross-site")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	if language != "" {
		req.Header.Add("Accept-Language", language)
	}

	res, err := client.Do(req)
	defer res.Body.Close()
	parse := parser.Parse(res.Body)
	return string(parse), nil
}
