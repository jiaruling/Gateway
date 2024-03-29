package controller

import (
	"encoding/json"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jiaruling/Gateway/dao"
	"github.com/jiaruling/Gateway/dto"
	"github.com/jiaruling/Gateway/global"
	"github.com/jiaruling/Gateway/middleware"
	"github.com/jiaruling/golang_utils/lib"
)

type AdminLoginController struct{}

func AdminLoginRegister(group *gin.RouterGroup) {
	adminLogin := &AdminLoginController{}
	group.POST("/login", adminLogin.AdminLogin)
	group.DELETE("/logout", adminLogin.AdminLoginOut)
}

// done: 管理员登陆
// AdminLogin godoc
// @Summary 管理员登陆
// @Description 管理员登陆
// @Tags 管理员接口
// @ID /admin_login/login
// @Accept  json
// @Produce  json
// @Param body body dto.AdminLoginInput true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOutput} "success"
// @Router /admin_login/login [post]
func (adminlogin *AdminLoginController) AdminLogin(c *gin.Context) {
	params := &dto.AdminLoginInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	//1. params.UserName 取得管理员信息 admininfo
	//2. admininfo.salt + params.Password sha256 => saltPassword
	//3. saltPassword==admininfo.password
	tx := lib.GetMysqlGorm()
	admin := &dao.Admin{}
	admin, err := admin.LoginCheck(tx, params)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}

	//设置session
	sessInfo := &dto.AdminSessionInfo{
		ID:        admin.Id,
		UserName:  admin.UserName,
		LoginTime: time.Now(),
	}
	sessBts, err := json.Marshal(sessInfo)
	if err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	sess := sessions.Default(c)
	sess.Set(global.AdminSessionInfoKey, string(sessBts))
	sess.Save()

	out := &dto.AdminLoginOutput{Token: admin.UserName}
	middleware.ResponseSuccess(c, out)
}

// done: 管理员退出
// AdminLogin godoc
// @Summary 管理员退出
// @Description 管理员退出
// @Tags 管理员接口
// @ID /admin_login/logout
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /admin_login/logout [delete]
func (adminlogin *AdminLoginController) AdminLoginOut(c *gin.Context) {
	sess := sessions.Default(c)
	sess.Delete(global.AdminSessionInfoKey)
	sess.Save()
	middleware.ResponseSuccess(c, "")
}
