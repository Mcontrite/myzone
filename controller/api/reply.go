package v1

import (
	"myzone/model"
	"myzone/package/app"
	"myzone/package/logging"
	"myzone/package/rcode"
	"myzone/package/session"
	reply_service "myzone/service/reply"
	"os"
	"strconv"
	"strings"

	file_package "myzone/package/file"

	"github.com/gin-gonic/gin"
)

func AddReply(c *gin.Context) {
	tid, _ := strconv.Atoi(c.DefaultPostForm("articleid", "1"))
	docutype, _ := strconv.Atoi(c.DefaultPostForm("doctuype", "0"))
	message := c.DefaultPostForm("message", "")
	uid, _ := strconv.Atoi(session.GetSession(c, "userid"))
	uip := c.ClientIP()
	code := rcode.SUCCESS
	attachFileString := c.PostForm("attachfiles")
	attachfiles := []string{}
	filesNum := 0
	if len(attachFileString) > 0 {
		attachfiles = strings.Split(attachFileString, ",")
		filesNum = len(attachfiles)
	}
	reply := &model.Reply{
		ArticleID:  tid,
		UserID:     uid,
		Isfirst:    0,
		UserIP:     uip,
		Doctype:    docutype,
		Message:    message,
		MessageFmt: message,
		FilesNum:   filesNum,
	}
	newReply, err := model.AddReply(reply)
	if err != nil {
		logging.Info("回复入库错误", err.Error())
		code = rcode.ERROR_SQL_INSERT_FAIL
		app.JsonErrResponse(c, code)
		return
	}
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
				ArticleID:   int(tid),
				ReplyID:     int(newReply.ID),
				UserID:      uid,
				Filename:    fname,
				Orgfilename: forginname,
				Filetype:    ftype,
				Filesize:    fsize,
			}
			model.AddAttach(attach)
		}
	}
	reply_service.AfterAddNewReply(newReply, tid)
	app.JsonOkResponse(c, code, nil)
}

func UpdateReply(c *gin.Context) {
	reply_id, _ := strconv.Atoi(c.DefaultPostForm("reply_id", "1"))
	doctype, _ := strconv.Atoi(c.DefaultPostForm("doctype", "0"))
	message := c.DefaultPostForm("message", "")
	uid, _ := strconv.Atoi(session.GetSession(c, "userid"))
	uip := c.ClientIP()
	code := rcode.SUCCESS
	oldReply, err := model.GetReplyById(reply_id)
	if err != nil {
		code = rcode.ERROR_UNFIND_DATA
		app.JsonErrResponse(c, code)
		return
	}
	if oldReply.UserID != uid {
		code = rcode.UNPASS
		app.JsonErrResponse(c, code)
		return
	}
	reply := model.Reply{
		UserIP:  uip,
		Doctype: doctype,
		Message: message,
	}
	model.UpdateReply(reply_id, reply)
	app.JsonOkResponse(c, code, nil)
}
