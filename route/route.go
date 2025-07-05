package route

import (
	"GuiXinSchool/controller"

	"github.com/gin-gonic/gin"
)


type Engine struct{
	e *gin.Engine
}

func NewEngine(authCtrl *controller.AuthCtrl,userCtrl *controller.UserCtrl) *Engine {
	e := gin.Default()

	g := e.Group("/api/v1")
	{
		g.POST("/login",authCtrl.Login)
	}
    
    userG := e.Group("/api/v1/user")
    {
        userG.PUT("/update_name",authCtrl.JWTAuthMiddleware(), userCtrl.UpdateUserName)
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