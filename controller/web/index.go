package web

import (
	"myzone/model"
	"myzone/package/setting"
	"myzone/service/user"
	"myzone/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const PAGE_SIZE int = 6

func Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	articleList, _ := model.GetArticleList(page)
	articleTotle, _ := model.GetArticleTotleCount()
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	pages := utils.Pagination("?page={page}", articleTotle, page, PAGE_SIZE)
	newestUser, _ := user.GetNewestTop12Users()
	articlesNum, _ := model.CountArticlesNum()
	replysNum, _ := model.CountReplyNum()
	usersNum, _ := model.CountUserNum()
	webname := setting.ServerSetting.Sitename
	description := setting.ServerSetting.Sitebrief
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"articleList":  articleList,
			"islogin":      islogin,
			"sessions":     sessions,
			"newestuser":   newestUser,
			"pages":        pages,
			"articles_num": articlesNum,
			"replys_num":   replysNum,
			"users_num":    usersNum,
			"webname":      webname,
			"description":  description,
		},
	)
}

func SayingIndex(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	commentList, _ := model.GetCommentList(page)
	sayingList, _ := model.GetSayingList(page)
	commentTotle, _ := model.GetCommentTotleCount()
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	pages := utils.Pagination("?page={page}", commentTotle, page, PAGE_SIZE)
	newestUser, _ := user.GetNewestTop12Users()
	sayingsNum, _ := model.CountSayingsNum()
	commentsNum, _ := model.CountCommentNum()
	usersNum, _ := model.CountUserNum()
	webname := setting.ServerSetting.Sitename
	description := setting.ServerSetting.Sitebrief
	c.HTML(
		http.StatusOK,
		"sayingindex.html",
		gin.H{
			"commentList":  commentList,
			"sayingList":   sayingList,
			"islogin":      islogin,
			"sessions":     sessions,
			"newestuser":   newestUser,
			"pages":        pages,
			"sayings_num":  sayingsNum,
			"comments_num": commentsNum,
			"users_num":    usersNum,
			"webname":      webname,
			"description":  description,
		},
	)
}
