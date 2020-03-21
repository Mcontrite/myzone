package model

import (
	"myzone/utils"
	"time"
)

type User struct {
	Model
	GroupID      int       `gorm:"default:0" json:"group_id"`      //用户组编号
	Username     string    `gorm:"default:''" json:"username"`     //用户名
	Password     string    `gorm:"default:''" json:"password"`     //密码
	Avatar       string    `gorm:"default:''" json:"avatar"`       //用户头像
	ArticlesCnt  int       `gorm:"default:0" json:"articles_cnt"`  //发帖数
	ReplysCnt    int       `gorm:"default:0" json:"replys_cnt"`    //回复数
	SayingsCnt   int       `gorm:"default:0" json:"sayings_cnt"`   //发帖数
	CommentsCnt  int       `gorm:"default:0" json:"comments_cnt"`  //回复数
	FavouriteCnt int       `gorm:"default:0" json:"favourite_cnt"` //收藏数
	CreateIp     string    `gorm:"default:''" json:"create_ip"`    //创建时IP
	LoginIp      string    `gorm:"default:''" json:"login_ip"`     //登录时IP
	LoginDate    time.Time `json:"login_date"`                     //登录时间
	LoginsCnt    int       `gorm:"default:0" json:"logins_cnt"`    //登录次数
	Group        Group
}

func GetUser(maps interface{}) (user User, err error) {
	err = db.Preload("Group").Model(&User{}).Where(maps).First(&user).Error
	return
}

func GetUserByID(id int) (user User, err error) {
	err = db.Preload("Group").Model(&User{}).Where("id = ?", id).First(&user).Error
	return
}

func GetUsers(num int, order string, maps interface{}) (user []User, err error) {
	err = db.Preload("Group").Model(&User{}).Order(order).Limit(num).Find(&user).Error
	return
}

func ExistUserByName(username string) bool {
	var user User
	db.Model(&User{}).Select("id").Where("username = ?", username).First(&user)
	if user.ID > 0 {
		return true
	}
	return false
}

func AddUser(username, password, ip string) (user *User, err error) {
	password, err = utils.BcryptString(password)
	if err != nil {
		return
	}
	user = &User{
		Username:  username,
		Password:  password,
		CreateIp:  ip,
		LoginDate: time.Now(),
	}
	err = db.Create(user).Error
	return
}

func AddUserPro(userinfo *User) (user *User, err error) {
	userinfo.Password, err = utils.BcryptString(userinfo.Password)
	if err != nil {
		return
	}
	user = userinfo
	err = db.Create(userinfo).Error
	return
}

func UpdateUser(whereMaps interface{}, updateItems map[string]interface{}) (err error) {
	err = db.Model(&User{}).Where(whereMaps).Updates(updateItems).Error
	return
}

func DelUser(whereMaps interface{}) (err error) {
	err = db.Unscoped().Where(whereMaps).Delete(&User{}).Error
	return
}

func UpdateUserArticlesCnt(id int, newArticlesCnt int) (err error) {
	err = db.Model(&User{}).Where("id = ?", id).Update("articles_cnt", newArticlesCnt).Error
	return
}

func UpdateUserSayingsCnt(id int, newSayingsCnt int) (err error) {
	err = db.Model(&User{}).Where("id = ?", id).Update("sayings_cnt", newSayingsCnt).Error
	return
}

func UpdateUserReplysCnt(id int, newReplysCnt int) (err error) {
	err = db.Model(&User{}).Where("id = ?", id).Update("replys_cnt", newReplysCnt).Error
	return
}

func UpdateUserCommentsCnt(id int, newCommentsCnt int) (err error) {
	err = db.Model(&User{}).Where("id = ?", id).Update("comments_cnt", newCommentsCnt).Error
	return
}

func CountUserNum() (usersNum int, err error) {
	err = db.Model(&User{}).Count(&usersNum).Error
	return
}
