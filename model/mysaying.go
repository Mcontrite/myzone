package model

type MySaying struct {
	Model
	UserID   int `gorm:"default:0" json:"user_id"`   //
	SayingID int `gorm:"default:0" json:"saying_id"` //
	User     User
	Saying   Saying
}

func GetMySayingList(uid int, page int, limit int, Orderby string) (sayings []MySaying, err error) {
	if page <= 1 {
		page = 1
	}
	if limit == 0 {
		limit = 20
	}
	err = db.Preload("User").Preload("Saying").Model(&MySaying{}).Where("user_id = ?", uid).Offset((page - 1) * 20).Limit(200).Order(Orderby).Find(&sayings).Error
	return
}

func AddMySaying(userID int, sayingID int) (mySaying *MySaying, err error) {
	mySaying = &MySaying{
		UserID:   userID,
		SayingID: sayingID,
	}
	err = db.Create(mySaying).Error
	return
}

func DelMySayingsOfSaying(tids []string) (err error) {
	err = db.Unscoped().Where("saying_id in (?)", tids).Delete(&MySaying{}).Error
	return
}
