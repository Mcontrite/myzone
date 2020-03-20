package user

import (
	"regexp"

	"github.com/astaxie/beego/validation"
)

func AddUserValid(v *validation.Validation, username string, password string) {
	ValidName(v, username)
	ValidPassword(v, password)
}

func LoginValidWithName(v *validation.Validation, name string, password string) {
	ValidPassword(v, password)
	ValidNameRequired(v, name)
}

func ValidName(v *validation.Validation, username string) {
	pass, _ := regexp.MatchString("[a-zA-Z0-9]{3,16}", username)
	if !pass {
		v.SetError("username", "名称只能是3-16位字母数字组合")
	}
}

func ValidNameRequired(v *validation.Validation, username string) {
	v.Required(username, "username").Message("密码不能为空")
}

func ValidPassword(v *validation.Validation, password string) {
	v.Required(password, "password").Message("密码不能为空")
}
