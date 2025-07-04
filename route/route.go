package route

import "github.com/gin-gonic/gin"


type Engine struct{
	e *gin.Engine
}

func NewEngine() *Engine {
	e := gin.Default()

	g := e.Group("/api/v1")
	{
		g.POST("/login",)
	}


	return &Engine{
		e: e,
	}
}


func (e *Engine) Run() {
	if err := e.e.Run(); err != nil {
		panic(err)
	}
}