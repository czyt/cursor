package cursor

import (
	"encoding/json"
	"go.czyt.tech/cursor/internal/builder"
	"go.czyt.tech/cursor/internal/parser"
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
	if err != nil {
		return nil, err
	}
	req, err := builder.NewRequest(http.MethodPost, testsUrl, "", payload)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	return response, nil
}

func (c *Client) Comment(comment CommentRequest) (response any, err error) {
	payload, err := json.Marshal(comment)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	req, err := builder.NewRequest(http.MethodPost, commentUrl, "", payload)
	if err != nil {
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	return
}

func (c *Client) Conversation(conversation ConversationRequest, language string) (response string, err error) {
	payload, err := json.Marshal(conversation)
	if err != nil {
		return "", err
	}
	client := &http.Client{}
	req, err := builder.NewRequest(http.MethodPost, conversationUrl, "", payload)
	if err != nil {
		return "", err
	}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	parse := parser.Parse(res.Body)
	return string(parse), nil
}
