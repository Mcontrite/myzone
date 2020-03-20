package v1

import (
	"myzone/model"
	"myzone/package/app"
	"myzone/package/logging"
	"myzone/package/rcode"
	"myzone/package/session"
	"strconv"
	"strings"
	"time"

	saying_service "myzone/service/saying"
	user_service "myzone/service/user"

	"github.com/gin-gonic/gin"
)

func AddSaying(c *gin.Context) {
	message := c.DefaultPostForm("message", "")
	code := rcode.SUCCESS
	uid, _ := strconv.Atoi(session.GetSession(c, "userid"))
	uip := c.ClientIP()
	saying := &model.Saying{
		UserID:   uid,
		UserIP:   uip,
		LastDate: time.Now(),
	}
	newSaying, err := model.AddSaying(saying)
	if err != nil {
		logging.Info("saying入库错误", err.Error())
		code = rcode.ERROR_SQL_INSERT_FAIL
		app.JsonErrResponse(c, code)
		return
	}
	comment := &model.Comment{
		SayingID:   int(newSaying.ID),
		UserID:     uid,
		Isfirst:    1,
		UserIP:     uip,
		Message:    message,
		MessageFmt: message,
	}
	newComment, err := model.AddComment(comment)
	if err != nil {
		logging.Info("comment入库错误", err.Error())
		code = rcode.ERROR
		code = rcode.ERROR_SQL_INSERT_FAIL
		app.JsonErrResponse(c, code)
		return
	}
	model.UpdateSaying(int(newSaying.ID), model.Saying{FirstCommentID: int(newComment.ID), LastDate: time.Now()})
	saying_service.AfterAddNewSaying(newSaying)
	app.JsonOkResponse(c, code, nil)
}

func DeleteSayings(c *gin.Context) {
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
	err := saying_service.DelSayings(idsSlice)
	if err != nil {
		code = rcode.ERROR
		app.JsonErrResponse(c, code)
		return
	}
	app.JsonOkResponse(c, code, ids)
}

func UpdateSaying(c *gin.Context) {
	saying_id, _ := strconv.Atoi(c.Param("id"))
	comment_id, _ := strconv.Atoi(c.DefaultPostForm("comment_id", "1"))
	message := c.DefaultPostForm("message", "")
	uid, _ := strconv.Atoi(session.GetSession(c, "userid"))
	uip := c.ClientIP()
	code := rcode.SUCCESS
	oldSaying, err := model.GetSayingById(saying_id)
	if err != nil {
		code = rcode.ERROR_UNFIND_DATA
		app.JsonErrResponse(c, code)
		return
	}
	if oldSaying.UserID != uid {
		code = rcode.UNPASS
		app.JsonErrResponse(c, code)
		return
	}
	oldComment, err := model.GetSayingFirstCommentByTid(saying_id)
	if err != nil {
		code = rcode.ERROR_UNFIND_DATA
		app.JsonErrResponse(c, code)
		return
	}
	if int(oldComment.ID) != comment_id {
		code = rcode.UNPASS
		app.JsonErrResponse(c, code)
		return
	}
	saying := model.Saying{
		UserIP: uip,
	}
	model.UpdateSaying(saying_id, saying)
	comment := model.Comment{
		UserIP:  uip,
		Message: message,
	}
	model.UpdateComment(comment_id, comment)
	app.JsonOkResponse(c, code, nil)
}
