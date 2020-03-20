package v1

import (
	"myzone/model"
	"myzone/package/app"
	file_package "myzone/package/file"
	"myzone/package/rcode"
	"myzone/package/session"
	"myzone/package/upload"
	"strconv"

	"github.com/gin-gonic/gin"
)

var pixgif = "data:image/gif;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVQImWNgYGBgAAAABQABh6FO1AAAAABJRU5ErkJggg=="

func CkeditorUpload(c *gin.Context) {
	file, _ := c.FormFile("upload")
	userid := session.GetSession(c, "userid")
	fileName := file.Filename
	if !upload.CheckImageSize2(file) {
		c.JSON(200, gin.H{
			"fileName": fileName,
			"uploaded": 1,
			"url":      pixgif,
		})
		return
	}
	newFilename := file_package.MakeFileName(userid, fileName)
	filepath := "upload/article/" + userid
	filepath, err := file_package.CreatePathInToday(filepath)
	if err != nil {
		c.JSON(200, gin.H{
			"fileName": fileName,
			"uploaded": 1,
			"url":      pixgif,
		})
		return
	}
	fullName := filepath + "/" + newFilename
	c.SaveUploadedFile(file, fullName)
	c.JSON(200, gin.H{
		"fileName": fileName,
		"uploaded": 1,
		"url":      "/" + fullName,
	})
}

func UploadFile(c *gin.Context) {
	action := c.Query("action")
	uid := c.Query("uid")
	code := rcode.SUCCESS
	file, _ := c.FormFile("upload")
	fileName := file.Filename
	newFilename := file_package.MakeFileName(uid, fileName)
	if !upload.CheckImageSize2(file) {
		code = rcode.ERROR_IMAGE_TOO_LARGE
		app.JsonErrResponse(c, code)
		return
	}
	filepath := "upload/" + action + "/" + uid + "/"
	err := file_package.CreatePath(filepath)
	if err != nil {
		code = rcode.ERROR_FILE_CREATE_FAIL
		app.JsonErrResponse(c, code)
		return
	}
	fullName := filepath + newFilename
	err = c.SaveUploadedFile(file, fullName)
	if err != nil {
		code = rcode.ERROR_FILE_SAVE_FAIL
		app.JsonErrResponse(c, code)
		return
	}
	c.JSON(200, gin.H{"filename": "测试图", "filetype": 1, "url": "/" + fullName, "attatchid": 99})
}

func UploadAttach(c *gin.Context) {
	userid := session.GetSession(c, "userid")
	code := rcode.SUCCESS
	file, _ := c.FormFile("upload")
	fileName := file.Filename
	fileType := file_package.GetType(fileName)
	newFilename := file_package.MakeFileName(userid, fileName)
	if !upload.CheckImageSize2(file) {
		code = rcode.ERROR_IMAGE_TOO_LARGE
		app.JsonErrResponse(c, code)
		return
	}
	filepath := "upload/attach/" + userid
	filepath, err := file_package.CreatePathInToday(filepath)
	if err != nil {
		code = rcode.ERROR_FILE_CREATE_FAIL
		app.JsonErrResponse(c, code)
		return
	}
	fullName := filepath + "/" + newFilename
	err = c.SaveUploadedFile(file, fullName)
	if err != nil {
		code = rcode.ERROR_FILE_SAVE_FAIL
		app.JsonErrResponse(c, code)
		return
	}
	data := map[string]interface{}{"orgfilename": fileName, "filetype": fileType, "url": fullName}
	app.JsonOkResponse(c, code, data)
}

func UploadAddAttach(c *gin.Context) {
	userid := session.GetSession(c, "userid")
	articleId, _ := strconv.Atoi(c.DefaultPostForm("article_id", "0"))
	posarticleId := articleId
	replyId, _ := strconv.Atoi(c.PostForm("reply_id"))
	code := rcode.SUCCESS
	file, _ := c.FormFile("upload")
	fileName := file.Filename
	fileType := file_package.GetType(fileName)
	fileSize := file.Size
	newFilename := file_package.MakeFileName(userid, fileName)
	if !upload.CheckImageSize2(file) {
		code = rcode.ERROR_IMAGE_TOO_LARGE
		app.JsonErrResponse(c, code)
		return
	}
	filepath := "upload/attach/" + userid
	filepath, err := file_package.CreatePathInToday(filepath)
	if err != nil {
		code = rcode.ERROR_FILE_CREATE_FAIL
		app.JsonErrResponse(c, code)
		return
	}
	fullName := filepath + "/" + newFilename
	err = c.SaveUploadedFile(file, fullName)
	if err != nil {
		code = rcode.ERROR_FILE_SAVE_FAIL
		app.JsonErrResponse(c, code)
		return
	}
	replyInfo, _ := model.GetReplyById(replyId)
	if articleId == 0 {
		articleId = replyInfo.ArticleID
	}
	useridInt, _ := strconv.Atoi(userid)
	model.AddAttach(&model.Attach{
		ArticleID:   articleId,
		ReplyID:     replyId,
		UserID:      useridInt,
		Filesize:    int(fileSize),
		Filename:    fullName,
		Orgfilename: fileName,
		Filetype:    fileType,
	})
	if posarticleId != 0 {
		articleInfo, _ := model.GetArticleById(articleId)
		model.UpdateArticleFilesNum(articleId, articleInfo.FilesNum+1)
	}
	model.UpdateReplyFilesNum(replyId, replyInfo.FilesNum+1)
	data := map[string]interface{}{"orgfilename": fileName, "filetype": fileType, "url": fullName}
	app.JsonOkResponse(c, code, data)
}

func DeleteAttach(c *gin.Context) {
	userid := session.GetSession(c, "userid")
	_ = userid
	attachId, _ := strconv.Atoi(c.PostForm("attach_id"))
	articleId, _ := strconv.Atoi(c.DefaultPostForm("article_id", "0"))
	replyId, _ := strconv.Atoi(c.DefaultPostForm("reply_id", "0"))
	code := rcode.SUCCESS
	if articleId != 0 {
		articleInfo, _ := model.GetArticleById(articleId)
		if articleInfo.FilesNum != 0 {
			model.UpdateArticleFilesNum(articleId, articleInfo.FilesNum-1)
		}
		replyId = articleInfo.FirstReplyID
	}
	if replyId != 0 {
		replyInfo, _ := model.GetReplyById(replyId)
		if replyInfo.FilesNum != 0 {
			model.UpdateReplyFilesNum(replyId, replyInfo.FilesNum-1)
		}
	}
	model.DelAttach(attachId)
	app.JsonOkResponse(c, code, nil)
}
