package admin

import (
	"myzone/model"
	"myzone/package/rcode"
	"myzone/package/session"
	"myzone/service/user"
	"myzone/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminLogin(c *gin.Context) {
	sessions := user.GetSessions(c)
	c.HTML(200, "rlogin.html", gin.H{"sessions": sessions})
}

func AdminLoginCheck(c *gin.Context) {
	username := session.GetSession(c, "username")
	password := c.DefaultPostForm("password", "")
	code := rcode.INVALID_PARAMS
	// 2，验证邮箱和密码
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	maps["username"] = username
	user, err := model.GetUser(maps)
	if err != nil {
		code = rcode.ERROR_NOT_EXIST_USER
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": rcode.GetMessage(code),
			"data":    data,
		})
		return
	}
	// 获取加密的密码
	hashPassword := user.Password
	if !utils.VerifyString(password, hashPassword) {
		code = rcode.ERROR_NOT_EXIST_USER
		c.JSON(http.StatusOK, gin.H{
			"code":    code,
			"message": rcode.GetMessage(code),
			"data":    data,
		})
		return
	}
	code = rcode.SUCCESS
	token, time, err := utils.GenerateToken(user.Username, password)
	if err != nil {
		code = rcode.ERROR_AUTH_TOKEN
	} else {
		data["token"] = token
		data["exp_time"] = time
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": rcode.GetMessage(code),
		"data":    data,
	})
}
