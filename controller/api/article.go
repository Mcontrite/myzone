package v1

import (
	"myzone/model"
	"myzone/package/app"
	"myzone/package/logging"
	"myzone/package/rcode"
	"myzone/package/session"
	"os"
	"strconv"
	"strings"
	"time"

	file_package "myzone/package/file"
	article_service "myzone/service/article"
	user_service "myzone/service/user"

	"github.com/gin-gonic/gin"
)

func AddArticle(c *gin.Context) {
	doctype, _ := strconv.Atoi(c.DefaultPostForm("doctype", "0"))
	title := c.DefaultPostForm("title", "")
	message := c.DefaultPostForm("message", "")
	attachFileString := c.PostForm("attachfiles")
	attachfiles := []string{}
	filesNum := 0
	code := rcode.SUCCESS
	if len(attachFileString) > 0 {
		attachfiles = strings.Split(attachFileString, ",")
		filesNum = len(attachfiles)
	}
	uid, _ := strconv.Atoi(session.GetSession(c, "userid"))
	uip := c.ClientIP()
	article := &model.Article{
		UserID:   uid,
		UserIP:   uip,
		Title:    title,
		FilesNum: filesNum,
		LastDate: time.Now(),
	}
	newArticle, err := model.AddArticle(article)
	if err != nil {
		logging.Info("article入库错误", err.Error())
		code = rcode.ERROR_SQL_INSERT_FAIL
		app.JsonErrResponse(c, code)
		return
	}
	reply := &model.Reply{
		ArticleID:  int(newArticle.ID),
		UserID:     uid,
		Isfirst:    1,
		UserIP:     uip,
		Doctype:    doctype,
		Message:    message,
		MessageFmt: message,
	}
	newReply, err := model.AddReply(reply)
	if err != nil {
		logging.Info("reply入库错误", err.Error())
		code = rcode.ERROR
		code = rcode.ERROR_SQL_INSERT_FAIL
		app.JsonErrResponse(c, code)
		return
	}
	model.UpdateArticle(int(newArticle.ID), model.Article{FirstReplyID: int(newReply.ID), LastDate: time.Now()})
	article_service.AfterAddNewArticle(newArticle)
	if len(attachFileString) > 0 {
		for _, attachfile := range attachfiles {
			file := strings.Split(attachfile, "|")
			fname := file[0]
			forginname := file[1]
			ftype := file_package.GetType(fname)
			ofile, err := os.Open(fname)
			defer ofile.Close()
			if err != nil {
				continue
			}
			fsize, _ := file_package.GetSize(ofile)
			attach := &model.Attach{
				ArticleID:   int(newArticle.ID),
				ReplyID:     int(newReply.ID),
				UserID:      uid,
				Filename:    fname,
				Orgfilename: forginname,
				Filetype:    ftype,
				Filesize:    fsize,
			}
			_, err = model.AddAttach(attach)
			if err != nil {
				logging.Info("attach入库错误", err.Error())
				code = rcode.ERROR_SQL_INSERT_FAIL
				app.JsonErrResponse(c, code)
				return
			}
		}
	}
	app.JsonOkResponse(c, code, nil)
}

type Tids struct {
	Tidarr []string `json:"tidarr"`
}

func DeleteArticles(c *gin.Context) {
	ids := c.PostForm("tidarr")
	code := rcode.SUCCESS
	idsSlice := strings.Split(ids, ",")
	uid, _ := strconv.Atoi(session.GetSession(c, "userid"))
	isadmin := user_service.IsAdmin(uid)
	if isadmin == "0" {
		code = rcode.UNPASS
		app.JsonErrResponse(c, code)
		return
	}
	err := article_service.DelArticles(idsSlice)
	if err != nil {
		code = rcode.ERROR
		app.JsonErrResponse(c, code)
		return
	}
	app.JsonOkResponse(c, code, ids)
}

func UpdateArticle(c *gin.Context) {
	article_id, _ := strconv.Atoi(c.Param("id"))
	reply_id, _ := strconv.Atoi(c.DefaultPostForm("reply_id", "1"))
	doctype, _ := strconv.Atoi(c.DefaultPostForm("doctype", "0"))
	title := c.DefaultPostForm("title", "")
	message := c.DefaultPostForm("message", "")
	uid, _ := strconv.Atoi(session.GetSession(c, "userid"))
	uip := c.ClientIP()
	code := rcode.SUCCESS
	oldArticle, err := model.GetArticleById(article_id)
	if err != nil {
		code = rcode.ERROR_UNFIND_DATA
		app.JsonErrResponse(c, code)
		return
	}
	if oldArticle.UserID != uid {
		code = rcode.UNPASS
		app.JsonErrResponse(c, code)
		return
	}
	oldReply, err := model.GetArticleFirstReplyByTid(article_id)
	if err != nil {
		code = rcode.ERROR_UNFIND_DATA
		app.JsonErrResponse(c, code)
		return
	}
	if int(oldReply.ID) != reply_id {
		code = rcode.UNPASS
		app.JsonErrResponse(c, code)
		return
	}
	article := model.Article{
		UserIP: uip,
		Title:  title,
	}
	model.UpdateArticle(article_id, article)
	reply := model.Reply{
		UserIP:  uip,
		Doctype: doctype,
		Message: message,
	}
	model.UpdateReply(reply_id, reply)
	app.JsonOkResponse(c, code, nil)
}

// 添加附件
// 直接添加到表中，因为以及各有了文章  所以可以直接添加
func AddarticleAttach(c *gin.Context) {
	// 获取文件内容
	// 获取articleid replyid uid
	// 修改article表的files字段 + 1
	// 在attach表中添加一天新的记录
}

// 删除的附件  知己额删除  提供好attach的id  就能删除
func DelarticleAttach(c *gin.Context) {
	// 删除数据内容  删除文件内容
	// 获取articleid
	// 修改article表的files字段 - 1
	// 在attach表中直接删除记录
}
