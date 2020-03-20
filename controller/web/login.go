package web

import (
	"myzone/package/setting"
	"myzone/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	webname := setting.ServerSetting.Sitename
	description := setting.ServerSetting.Sitebrief
	c.HTML(
		http.StatusOK,
		"login.html",
		gin.H{
			"title":       "Home Page",
			"islogin":     islogin,
			"sessions":    sessions,
			"description": description,
			"webname":     webname,
		},
	)
}
