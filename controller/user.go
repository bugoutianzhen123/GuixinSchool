package controller

import (
	"GuiXinSchool/pkg"
	"GuiXinSchool/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserCtrl struct {
	us *service.UserSvc
}

type UpdateUserNameReq struct {
	StuID string `json:"stu_id" binding:"required"`
}


func(u *UserCtrl) UpdateUserName(c *gin.Context) {
	var req UpdateUserNameReq
	if err:= pkg.ParseURL(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, pkg.WithMsg(pkg.ParamErrResp, err.Error()))
		return
	}

	if err := u.us.UpdateName(c, req.StuID, req.StuID); err != nil {
		c.JSON(http.StatusInternalServerError, pkg.WithMsg(pkg.AuthResp, "更新用户名失败"))
		return
	}
	c.JSON(http.StatusOK, pkg.WithMsg(pkg.SuccessResp, "用户名更新成功"))
}