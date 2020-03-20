package admin

import (
	"myzone/model"

	"github.com/gin-gonic/gin"
)

func GetArticleList(c *gin.Context) {
	var articles []model.Article
	articles, _ = model.GetArticles(map[string]interface{}{"deleted_at": nil}, "created_at desc", 100, 1)
	c.HTML(200, "aarticlelist.html", gin.H{"articles": articles})
}
