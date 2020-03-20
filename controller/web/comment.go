package web

import (
	"myzone/model"
	"myzone/service/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func EditComment(c *gin.Context) {
	commentId, _ := strconv.Atoi(c.Param("id"))
	comment, _ := model.GetCommentById(commentId)
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	c.HTML(
		http.StatusOK,
		"editcomment.html",
		gin.H{
			"comment":  comment,
			"islogin":  islogin,
			"sessions": sessions,
		},
	)
}
