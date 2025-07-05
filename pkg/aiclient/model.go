package aiclient

type Message struct {
	Role    string `json:"role"` // user/assistant/system
	Content string `json:"content"`
}

type ChatRequest struct {
	Messages  []Message `json:"messages"`
	Model     string    `json:"model,omitempty"`
	MaxTokens int       `json:"max_tokens,omitempty"`
}

type ChatResponse struct {
	Message  Message `json:"message"`
	Usage    Usage   `json:"usage"`
	Finished bool    `json:"finished"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
