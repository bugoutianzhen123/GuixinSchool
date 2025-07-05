package aiclient

type OpenAIClient struct {
	apiKey string
	model  string
}

func NewOpenAIClient(apiKey string) *OpenAIClient {
	return &OpenAIClient{
		apiKey: apiKey,
		model:  "gpt-3.5-turbo", // 默认模型
	}
}
