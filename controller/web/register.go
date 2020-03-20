package web

import (
	"myzone/package/setting"
	"myzone/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	webname := setting.ServerSetting.Sitename
	description := setting.ServerSetting.Sitebrief
	c.HTML(
		http.StatusOK,
		"register.html",
		gin.H{
			"title":       "用户注册",
			"islogin":     islogin,
			"sessions":    sessions,
			"webname":     webname,
			"description": description,
		},
	)
}
