package deepseek

import (
	"GuiXinSchool/pkg/aiclient"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ChatRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}


// GetContent returns the content of the messages in the request
func (c ChatRequest) GetContent() []byte {
	data, err := json.Marshal(c)
	if err != nil {
		return nil // or handle error as needed
	}
	return data
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}


type ChatResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// GetContent returns the content of the first choice in the response
func (c ChatResponse) GetContent() []byte {
	data, err := json.Marshal(c.Choices[0].Message.Content)
	if err != nil {
		return nil // or handle error as needed
	}
	return data
}

type ChatSvc struct {
	cli     *http.Client
	baseURL string
	apiKey  string
}

func NewChatSvc(apiKey string) *ChatSvc {
	return &ChatSvc{
		cli:     &http.Client{},
		baseURL: "https://api.deepseek.com/chat/completions",
		apiKey:  apiKey,
	}
}

func (c *ChatSvc) Chat(ctx context.Context, req aiclient.ChatRequest) (aiclient.ChatResponse, error) {
	
	//先进行类型断言
	_, ok := req.(ChatRequest)
	if !ok {
		return nil, fmt.Errorf("invalid request type, expected deepseek.ChatRequest")
	}
	
	var response ChatResponse

	// 序列化请求体
	reqBody, err := json.Marshal(req)
	if err != nil {
		return response, fmt.Errorf("failed to marshal request: %w", err)
	}

	// 创建 HTTP 请求
	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.baseURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return response, fmt.Errorf("failed to create request: %w", err)
	}

	// 设置请求头
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+c.apiKey)

	// 发送请求
	resp, err := c.cli.Do(httpReq)
	if err != nil {
		return response, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return response, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	// 解析响应体
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return response, fmt.Errorf("failed to decode response: %w", err)
	}

	return response, nil
}
