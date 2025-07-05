package aiclient

type TongyiClient struct {
	apiKey string
	model  string
}

func NewTongyiClient(apiKey string) *TongyiClient {
	return &TongyiClient{
		apiKey: apiKey,
		model:  "qwen-turbo", // 默认模型
	}
}
