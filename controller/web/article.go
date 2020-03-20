package web

import (
	"html"
	"myzone/model"
	"myzone/package/setting"
	article_service "myzone/service/article"
	"myzone/service/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Article(c *gin.Context) {
	articleId, _ := strconv.Atoi(c.Param("id"))
	article, err := model.GetArticleById(articleId)
	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", gin.H{})
		return
	}
	freply, _ := model.GetArticleFirstReplyByTid(articleId)
	freply.MessageFmt = html.UnescapeString(freply.MessageFmt)
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	replylist, _ := model.GetArticleReplyListByTid(articleId, 500, 1)
	replylistLen := len(replylist)
	attachs, _ := model.GetAttachsByReplyId(int(freply.ID))
	isfav, _ := model.CheckFavourite(sessions.Userid, articleId)
	userNewestArticles, _ := article_service.GetUserArticles(article.UserID)
	model.UpdateArticleViewsCnt(articleId)
	description := setting.ServerSetting.Sitebrief
	c.HTML(
		http.StatusOK,
		"article.html",
		gin.H{
			"article":              article,
			"freply":               freply,
			"islogin":              islogin,
			"sessions":             sessions,
			"replylist":            replylist,
			"reply_list_len":       replylistLen,
			"user_newest_articles": userNewestArticles,
			"attachs":              attachs,
			"isfav":                isfav,
			"description":          description,
		},
	)
}

func ArticleAddReply(c *gin.Context) {
	articleId, _ := strconv.Atoi(c.Param("id"))
	sessions := user.GetSessions(c)
	islogin := user.IsLogin(c)
	description := setting.ServerSetting.Sitebrief
	c.HTML(
		http.StatusOK,
		"advance_reply.html",
		gin.H{
			"sessions":    sessions,
			"islogin":     islogin,
			"article_id":  articleId,
			"description": description,
		},
	)
}

func NewArticle(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	description := setting.ServerSetting.Sitebrief
	c.HTML(
		http.StatusOK,
		"newarticle.html",
		gin.H{
			"islogin":     islogin,
			"sessions":    sessions,
			"description": description,
		},
	)
}

func EditArticle(c *gin.Context) {
	articleId, _ := strconv.Atoi(c.Param("id"))
	article, _ := model.GetArticleById(articleId)
	freply, _ := model.GetArticleFirstReplyByTid(articleId)
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	attachs, _ := model.GetAttachsByReplyId(int(freply.ID))
	webname := setting.ServerSetting.Sitename
	description := setting.ServerSetting.Sitebrief
	c.HTML(
		http.StatusOK,
		"editarticle.html",
		gin.H{
			"article":     article,
			"freply":      freply,
			"islogin":     islogin,
			"sessions":    sessions,
			"attachs":     attachs,
			"webname":     webname,
			"description": description,
		},
	)
}
