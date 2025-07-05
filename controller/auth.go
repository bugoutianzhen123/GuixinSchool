package controller

import (
	"GuiXinSchool/pkg"
	"GuiXinSchool/service"
	"net/http"

	"github.com/gin-gonic/gin"
)


type AuthCtrl struct{
	as *service.AuthSvc
}

func NewAuthCtrl(as *service.AuthSvc) *AuthCtrl {
	return &AuthCtrl{
		as: as,
	}
}

type LoginReq struct {
	StuID   string `json:"stu_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login handles user login requests.
func(ac *AuthCtrl) Login(c *gin.Context) {
	var req LoginReq
	if err:=pkg.ParseBody(c,&req); err!=nil{
		c.JSON(http.StatusBadRequest, pkg.WithMsg(pkg.ParamErrResp, err.Error()))
		return
	}


	if err := ac.as.Login(c.Request.Context(), req.StuID, req.Password);err!=nil {
		c.JSON(http.StatusUnauthorized, pkg.WithMsg(pkg.AuthResp, err.Error()))
		return
	}

	type TokenData struct{
		Token string `json:"token"`
	}

	token,err := ac.as.GetToken(c.Request.Context(), req.StuID)
	if err!=nil {
		c.JSON(http.StatusInternalServerError, pkg.WithMsg(pkg.AuthResp, err.Error()))
		return
	}
	c.JSON(http.StatusOK, pkg.WithData(pkg.SuccessResp, TokenData{Token: token}))
}


func (ac *AuthCtrl) JWTAuthMiddleware() gin.HandlerFunc {
	return ac.as.JWTAuthMiddleware()
}