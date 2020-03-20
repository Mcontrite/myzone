package user

import (
	"myzone/model"
)

func GetUserByID(uid int) (user model.User, err error) {
	wmap := map[string]interface{}{"id": uid}
	user, err = model.GetUser(wmap)
	return
}

func ResetPassword(newPassword string, uid int) (err error) {
	var wmap = make(map[string]interface{})
	wmap["id"] = uid
	err = model.UpdateUser(wmap, map[string]interface{}{"password": newPassword})
	return
}

func ResetAvatar(newAvatar string, uid int) (err error) {
	var wmap = make(map[string]interface{})
	wmap["id"] = uid
	err = model.UpdateUser(wmap, map[string]interface{}{"avatar": "/" + newAvatar})
	return
}

func ResetName(newName string, uid int) (err error) {
	var wmap = make(map[string]interface{})
	wmap["id"] = uid
	err = model.UpdateUser(wmap, map[string]interface{}{"username": newName})
	return
}

func DelUserByID(uid int) (err error) {
	wmap := map[string]interface{}{"id": uid}
	err = model.DelUser(wmap)
	return
}

func IsAdmin(ugid int) string {
	if ugid > 0 && ugid < 6 {
		return "1"
	}
	return "0"
}

func GetNewestTop12Users() (userList []model.User, err error) {
	return model.GetUsers(20, "id desc", 1)
}
