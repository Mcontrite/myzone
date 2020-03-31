package web

import (
	"html"
	"myzone/model"
	"myzone/package/setting"
	"myzone/service/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Saying(c *gin.Context) {
	sayingId, _ := strconv.Atoi(c.Param("id"))
	saying, err := model.GetSayingById(sayingId)
	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", gin.H{})
		return
	}
	fcomment, _ := model.GetSayingFirstCommentByTid(sayingId)
	fcomment.MessageFmt = html.UnescapeString(fcomment.MessageFmt)
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	commentlist, _ := model.GetSayingCommentListByTid(sayingId, 500, 1)
	commentlistLen := len(commentlist)
	model.UpdateSayingViewsCnt(sayingId)
	c.HTML(
		http.StatusOK,
		"saying.html",
		gin.H{
			"saying":           saying,
			"fcomment":         fcomment,
			"islogin":          islogin,
			"sessions":         sessions,
			"commentlist":      commentlist,
			"comment_list_len": commentlistLen,
		},
	)
}

func SayingAddComment(c *gin.Context) {
	sayingId, _ := strconv.Atoi(c.Param("id"))
	sessions := user.GetSessions(c)
	islogin := user.IsLogin(c)

	c.HTML(
		http.StatusOK,
		"advance_comment.html",
		gin.H{
			"sessions":  sessions,
			"islogin":   islogin,
			"saying_id": sayingId,
		},
	)
}

func NewSaying(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)

	c.HTML(
		http.StatusOK,
		"newsaying.html",
		gin.H{
			"islogin":  islogin,
			"sessions": sessions,
		},
	)
}

func EditSaying(c *gin.Context) {
	sayingId, _ := strconv.Atoi(c.Param("id"))
	saying, _ := model.GetSayingById(sayingId)
	fcomment, _ := model.GetSayingFirstCommentByTid(sayingId)
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	webname := setting.ServerSetting.Sitename

	c.HTML(
		http.StatusOK,
		"editsaying.html",
		gin.H{
			"saying":   saying,
			"fcomment": fcomment,
			"islogin":  islogin,
			"sessions": sessions,
			"webname":  webname,
		},
	)
}
