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

const PAGE_SIZE int = 10

func Index(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	articleList, _ := model.GetArticleList(page)
	sayingList, _ := model.GetSayingList(page)
	articleTotle, _ := model.GetArticleTotleCount()
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	pages := utils.Pagination("?page={page}", articleTotle, page, PAGE_SIZE)
	newestUser, _ := user.GetNewestTop12Users()
	articlesNum, _ := model.CountArticlesNum()
	sayingsNum, _ := model.CountSayingsNum()
	replysNum, _ := model.CountReplyNum()
	usersNum, _ := model.CountUserNum()
	webname := setting.ServerSetting.Sitename
	description := setting.ServerSetting.Sitebrief
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"articleList":  articleList,
			"sayingList":   sayingList,
			"islogin":      islogin,
			"sessions":     sessions,
			"newestuser":   newestUser,
			"pages":        pages,
			"articles_num": articlesNum,
			"sayings_num":  sayingsNum,
			"replys_num":   replysNum,
			"users_num":    usersNum,
			"webname":      webname,
			"description":  description,
		},
	)
}
