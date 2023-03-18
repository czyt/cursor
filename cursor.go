package cursor

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
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

type Decoder struct {
	r *bufio.Reader
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		r: bufio.NewReader(r),
	}
}

func (d *Decoder) Decode() (Event, error) {
	event := Event{}

	line, err := d.r.ReadString('\n')
	if err != nil {
		return event, err
	}

	if line == "\n" {
		return event, nil
	}

	if line[0] == ':' {
		return event, nil
	}

	if colon := strings.IndexByte(line, ':'); colon != -1 {
		field := line[:colon]
		value := strings.TrimLeft(line[colon+1:], " ")

		switch field {
		case "event":
			event.Name = value
		case "id":
			event.ID = value
		case "retry":
			if n, err := strconv.Atoi(value); err == nil {
				event.Retry = time.Duration(n) * time.Millisecond
			}
		default:
			if strings.HasPrefix(field, "data") {
				event.Data += value + "\n"
			}
		}
	}

	return event, nil
}

type Event struct {
	Name  string
	ID    string
	Data  string
	Retry time.Duration
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
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return
	}
	fmt.Println(string(body))
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
	all, _ := io.ReadAll(resp.Body)
	fmt.Println(string(all))
	return
}

func (c *Client) Conversation(conversation ConversationRequest) (response any, err error) {
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

	res, err := client.Do(req)
	defer res.Body.Close()
	parse := Parse(res.Body)
	fmt.Println(string(parse))
	return
}
