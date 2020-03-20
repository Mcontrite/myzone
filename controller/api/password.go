package v1

import (
	"myzone/model"
	"myzone/package/app"
	"myzone/package/rcode"
	"myzone/utils"

	"github.com/gin-gonic/gin"
)

// 用户重设密码
func UserResetPassword(c *gin.Context) {
	password := c.PostForm("password")
	code := rcode.SUCCESS
	password, _ = utils.BcryptString(password)
	var wmap = make(map[string]interface{})
	err := model.UpdateUser(wmap, map[string]interface{}{"password": password})
	if err != nil {
		code = rcode.ERROR_SQL_UPDATE_FAIL
		app.JsonErrResponse(c, code)
		return
	}
	app.JsonOkResponse(c, code, nil)
}
