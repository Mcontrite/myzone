package web

import (
	"myzone/model"
	"myzone/package/setting"
	"myzone/service/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	MY_INFO = iota
	MY_PASSWORD
	MY_AVATAR
	MY_NAME
	MY_ARTICLES
	MY_SAYINGS
	MY_FAVS
	MY_REPLYS
	MY_COMMENTS
	MY_NOTICE
)

const (
	U_INFO = iota
	U_ARTICLES
	U_SAYINGS
	U_REPLYS
	U_COMMENTS
)

var (
	webname     string
	description string
)

func init() {
	webname = setting.ServerSetting.Sitename
	description = setting.ServerSetting.Sitebrief
}

func MyInfo(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	uid := sessions.Userid
	userinfo, _ := user.GetUserByID(uid)
	tpl := "my_info.html"
	c.HTML(200, tpl, gin.H{
		"islogin":     islogin,
		"sessions":    sessions,
		"userinfo":    userinfo,
		"webname":     webname,
		"description": description,
	})
}

func MyEdit(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	tpl := "my_edit.html"
	c.HTML(200, tpl, gin.H{
		"islogin":     islogin,
		"sessions":    sessions,
		"webname":     webname,
		"description": description,
	})
}

func MyArticle(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	uid := sessions.Userid
	myArticles, _ := model.GetMyArticleList(uid, page, 20, "created_at desc")
	tpl := "my_article.html"
	c.HTML(200, tpl, gin.H{
		"islogin":     islogin,
		"sessions":    sessions,
		"myarticles":  myArticles,
		"webname":     webname,
		"description": description,
	})
}

func MySaying(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	uid := sessions.Userid
	mySayings, _ := model.GetMySayingList(uid, page, 20, "created_at desc")
	tpl := "my_saying.html"
	c.HTML(200, tpl, gin.H{
		"islogin":     islogin,
		"sessions":    sessions,
		"mysayings":   mySayings,
		"webname":     webname,
		"description": description,
	})
}

func MyFavorite(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	uid := sessions.Userid
	favArticles, _ := model.GetMyFavouriteList(uid, 1, 200, "created_at desc")
	tpl := "my_favorite.html"
	c.HTML(200, tpl, gin.H{
		"islogin":     islogin,
		"sessions":    sessions,
		"articles":    favArticles,
		"webname":     webname,
		"description": description,
	})
}

func MyReply(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	uid := sessions.Userid
	replys, _ := model.GetMyReplyList(uid, page, 20, "created_at desc")
	tpl := "my_reply.html"
	c.HTML(200, tpl, gin.H{
		"islogin":     islogin,
		"sessions":    sessions,
		"replys":      replys,
		"webname":     webname,
		"description": description,
	})
}

func MyComment(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	uid := sessions.Userid
	comments, _ := model.GetMyCommentList(uid, page, 20, "created_at desc")
	tpl := "my_comment.html"
	c.HTML(200, tpl, gin.H{
		"islogin":     islogin,
		"sessions":    sessions,
		"comments":    comments,
		"webname":     webname,
		"description": description,
	})
}

func UserInfo2(c *gin.Context) {
	action, _ := strconv.Atoi(c.DefaultQuery("action", "0"))
	var tpl string
	switch action {
	case U_ARTICLES:
		tpl = "u_article.html"
	case U_SAYINGS:
		tpl = "u_saying.html"
	case U_REPLYS:
		tpl = "u_replys.html"
	case U_COMMENTS:
		tpl = "u_comments.html"
	default:
		tpl = "u_info.html"
	}
	c.HTML(200, tpl, gin.H{})
}

func UserInfo(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	uid, _ := strconv.Atoi(c.Param("id"))
	userinfo, _ := user.GetUserByID(uid)
	tpl := "u_info.html"
	c.HTML(200, tpl, gin.H{
		"islogin":     islogin,
		"sessions":    sessions,
		"userinfo":    userinfo,
		"webname":     webname,
		"description": description,
	})
}

func UserArticle(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	uid, _ := strconv.Atoi(c.Param("id"))
	userinfo, _ := user.GetUserByID(uid)
	myArticles, _ := model.GetMyArticleList(uid, page, 20, "created_at desc")
	tpl := "u_articles.html"
	c.HTML(200, tpl, gin.H{
		"islogin":     islogin,
		"sessions":    sessions,
		"myarticles":  myArticles,
		"userinfo":    userinfo,
		"webname":     webname,
		"description": description,
	})
}

func UserSaying(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	uid, _ := strconv.Atoi(c.Param("id"))
	userinfo, _ := user.GetUserByID(uid)
	mySayings, _ := model.GetMySayingList(uid, page, 20, "created_at desc")
	tpl := "u_sayings.html"
	c.HTML(200, tpl, gin.H{
		"islogin":     islogin,
		"sessions":    sessions,
		"mysayings":   mySayings,
		"userinfo":    userinfo,
		"webname":     webname,
		"description": description,
	})
}

func UserReply(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	uid, _ := strconv.Atoi(c.Param("id"))
	userinfo, _ := user.GetUserByID(uid)
	replys, _ := model.GetMyReplyList(uid, page, 20, "created_at desc")
	tpl := "u_replys.html"
	c.HTML(200, tpl, gin.H{
		"islogin":     islogin,
		"sessions":    sessions,
		"replys":      replys,
		"userinfo":    userinfo,
		"webname":     webname,
		"description": description,
	})
}

func UserComment(c *gin.Context) {
	islogin := user.IsLogin(c)
	sessions := user.GetSessions(c)
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	uid, _ := strconv.Atoi(c.Param("id"))
	userinfo, _ := user.GetUserByID(uid)
	comments, _ := model.GetMyCommentList(uid, page, 20, "created_at desc")
	tpl := "u_comments.html"
	c.HTML(200, tpl, gin.H{
		"islogin":     islogin,
		"sessions":    sessions,
		"comments":    comments,
		"userinfo":    userinfo,
		"webname":     webname,
		"description": description,
	})
}
