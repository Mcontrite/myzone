package web

import (
	"myzone/model"
	"myzone/service/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func EditReply(c *gin.Context) {
	replyId, _ := strconv.Atoi(c.Param("id"))
	reply, _ := model.GetReplyById(replyId)
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	attachs, _ := model.GetAttachsByReplyId(replyId)
	c.HTML(
		http.StatusOK,
		"editreply.html",
		gin.H{
			"reply":    reply,
			"islogin":  islogin,
			"sessions": sessions,
			"attachs":  attachs,
		},
	)
}
