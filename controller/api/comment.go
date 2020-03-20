package v1

import (
	"myzone/model"
	"myzone/package/app"
	"myzone/package/logging"
	"myzone/package/rcode"
	"myzone/package/session"
	comment_service "myzone/service/comment"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddComment(c *gin.Context) {
	tid, _ := strconv.Atoi(c.DefaultPostForm("sayingid", "1"))
	message := c.DefaultPostForm("message", "")
	uid, _ := strconv.Atoi(session.GetSession(c, "userid"))
	uip := c.ClientIP()
	code := rcode.SUCCESS
	comment := &model.Comment{
		SayingID:   tid,
		UserID:     uid,
		Isfirst:    0,
		UserIP:     uip,
		Message:    message,
		MessageFmt: message,
	}
	newComment, err := model.AddComment(comment)
	if err != nil {
		logging.Info("回复帖子入库错误", err.Error())
		code = rcode.ERROR_SQL_INSERT_FAIL
		app.JsonErrResponse(c, code)
		return
	}
	comment_service.AfterAddNewComment(newComment, tid)
	app.JsonOkResponse(c, code, nil)
}

func UpdateComment(c *gin.Context) {
	comment_id, _ := strconv.Atoi(c.DefaultPostForm("comment_id", "1"))
	message := c.DefaultPostForm("message", "")
	uid, _ := strconv.Atoi(session.GetSession(c, "userid"))
	uip := c.ClientIP()
	code := rcode.SUCCESS
	oldComment, err := model.GetCommentById(comment_id)
	if err != nil {
		code = rcode.ERROR_UNFIND_DATA
		app.JsonErrResponse(c, code)
		return
	}
	if oldComment.UserID != uid {
		code = rcode.UNPASS
		app.JsonErrResponse(c, code)
		return
	}
	comment := model.Comment{
		UserIP:  uip,
		Message: message,
	}
	model.UpdateComment(comment_id, comment)
	app.JsonOkResponse(c, code, nil)
}
