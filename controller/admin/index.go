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
				"users":    usersNum,
				"attachs":  attachsNum,
			},
		})
}
