package builder

import (
	"bytes"
	"net/http"
)

func NewRequest(method string, host string, acceptLang string, payload []byte) (request *http.Request, err error) {
	req, err := http.NewRequest(method, host, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Cursor/0.1.1 Chrome/108.0.5359.62 Electron/22.0.0 Safari/537.36")
	req.Header.Add("Sec-Fetch-Site", "cross-site")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	if acceptLang != "" {
		req.Header.Add("Accept-Language", acceptLang)
	}
	return req, nil
}
