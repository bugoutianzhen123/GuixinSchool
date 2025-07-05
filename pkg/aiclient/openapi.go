package aiclient

import "context"

type AIClient interface {
	Chat(ctx context.Context,req ChatRequest) (ChatResponse,error)
}

type ChatRequest interface {
	GetContent() []byte //获取内容
}
type ChatResponse interface {
	GetContent() []byte //获取内容
}