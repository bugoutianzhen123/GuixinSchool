package aiclient

import "context"

type AIClient interface {
	Chat(ctx context.Context,req ChatRequest) (ChatResponse,error)
	// TODO
	// ChatStream(ctx context.Context,req ChatRequest,fn func(chunk string) error) error
}

type ChatRequest interface {
	IsStream() bool //是否流式
	GetContent() []byte //获取内容
}
type ChatResponse interface {
	GetContent() []byte //获取内容
}