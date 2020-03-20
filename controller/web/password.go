package web

import (
	"myzone/package/setting"
	"myzone/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ForgetPassword(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	webname := setting.ServerSetting.Sitename
	description := setting.ServerSetting.Sitebrief
	c.HTML(
		http.StatusOK,
		"forgetpass.html",
		gin.H{
			"title":       "用户重置密码",
			"islogin":     islogin,
			"sessions":    sessions,
			"webname":     webname,
			"description": description,
		},
	)
}

func ResetForgetPassword(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	webname := setting.ServerSetting.Sitename
	description := setting.ServerSetting.Sitebrief
	c.HTML(
		http.StatusOK,
		"reset_password.html",
		gin.H{
			"title":       "重设密码",
			"islogin":     islogin,
			"sessions":    sessions,
			"webname":     webname,
			"description": description,
		},
	)
}
