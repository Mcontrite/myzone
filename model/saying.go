package model

import "time"

type Saying struct {
	Model
	UserID         int       `gorm:"default:0" json:"user_id"`          //
	ViewsCnt       int       `gorm:"default:0" json:"views_cnt"`        //查看次数, 剥离出去，单独的服务，避免 cache 失效
	CommentsCnt    int       `gorm:"default:0" json:"comments_cnt"`     //回复数
	LastDate       time.Time `json:"last_date"`                         //最后回复时间
	FirstCommentID int       `gorm:"default:0" json:"first_comment_id"` //首贴 pid
	LastCommentID  int       `gorm:"default:0" json:"last_comment_id"`  //最后回复的 pid
	LastUserID     int       `gorm:"default:0" json:"last_user_id"`     //最近参与的 uid
	UserIP         string    `gorm:"default:''" json:"userip"`          //发帖时用户ip
	User           User      `json:"user"`
	LastUser       User
}

func GetSayings(whereMap interface{}, order string, limit int, page int) (saying []Saying, err error) {
	err = db.Model(&Saying{}).Preload("User").Where(whereMap).Order(order).Offset((page - 1) * limit).Limit(limit).Find(&saying).Error
	return
}

func GetSayingById(id int) (saying Saying, err error) {
	err = db.Preload("User").Where("id = ?", id).Model(&Saying{}).First(&saying).Error
	return
}

func GetSayingList(page int) (sayings []Saying, err error) {
	if page <= 1 {
		page = 1
	}
	err = db.Preload("User").Model(&Saying{}).Order("created_at desc").Offset((page - 1) * PAGE_SIZE).Limit(PAGE_SIZE).Find(&sayings).Error
	return
}

func GetSayingTotleCount() (totle int, err error) {
	err = db.Model(&Saying{}).Count(&totle).Error
	return
}

func GetSayingTotal(maps interface{}) (count int) {
	db.Model(&Saying{}).Where(maps).Count(&count)
	return
}

func AddSaying(saying *Saying) (*Saying, error) {
	err := db.Model(&Saying{}).Create(saying).Error
	return saying, err
}

func UpdateSaying(id int, saying Saying) (upsaying Saying, err error) {
	err = db.Model(&Saying{}).Where("id = ?", id).Updates(saying).Error
	upsaying, err = GetSayingById(id)
	return
}

func UpdateSayingPro(id int, items map[string]interface{}) (upsaying Saying, err error) {
	err = db.Model(&Saying{}).Where("id = ?", id).Updates(items).Error
	upsaying, err = GetSayingById(id)
	return
}

func DelSaying(ids []string) (err error) {
	err = db.Unscoped().Where("id in (?)", ids).Delete(&Saying{}).Error
	return
}

func UpdateSayingViewsCnt(id int) error {
	saying, _ := GetSayingById(id)
	return db.Model(&Saying{}).Where("id = ?", id).Update("views_cnt", saying.ViewsCnt+1).Error
}

func CountSayingsNum() (sayingsNum int, err error) {
	err = db.Model(&Saying{}).Count(&sayingsNum).Error
	return
}

func GetSayingsByIDs(ids []string) (sayings []*Saying, err error) {
	err = db.Model(&Saying{}).Preload("User").Where("id in (?)", ids).Find(&sayings).Error
	return
}
