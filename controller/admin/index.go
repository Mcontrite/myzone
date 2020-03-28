package admin

import (
	"myzone/model"
	"myzone/service/user"

	"github.com/gin-gonic/gin"
)

func AdminIndex(c *gin.Context) {
	sessions := user.GetSessions(c)
	articlesNum, _ := model.CountArticlesNum()
	replysNum, _ := model.CountReplyNum()
	sayingsNum, _ := model.CountSayingsNum()
	commentsNum, _ := model.CountCommentNum()
	usersNum, _ := model.CountUserNum()
	attachsNum, _ := model.CountAttachsNum()
	c.HTML(
		200,
		"aindex.html",
		gin.H{
			"sessions": sessions,
			"counts": map[string]interface{}{
				"articles": articlesNum,
				"replys":   replysNum,
				"sayings":  sayingsNum,
				"comments": commentsNum,
				"users":    usersNum,
				"attachs":  attachsNum,
			},
		})
}
