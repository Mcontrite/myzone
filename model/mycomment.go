package model

type MyComment struct {
	Model
	UserID    int `gorm:"primary_key;default:0" json:"user_id"`    //
	SayingID  int `gorm:"default:0" json:"saying_id"`              //
	CommentID int `gorm:"primary_key;default:0" json:"comment_id"` //
	User      User
	Saying    Saying
	Comment   Comment
}

func GetMyCommentList(uid int, page int, limit int, Orderby string) (comments []MyComment, err error) {
	if page <= 1 {
		page = 1
	}
	if limit == 0 {
		limit = 20
	}
	err = db.Preload("Saying").Preload("Saying.User").Preload("Saying.User.Group").Preload("Comment").Model(&MyComment{}).Where("user_id = ?", uid).Offset((page - 1) * 20).Limit(200).Order(Orderby).Find(&comments).Error
	return
}

func AddMyComment(userID int, sayingID int, commentID int) (myComment *MyComment, err error) {
	myComment = &MyComment{
		UserID:    userID,
		SayingID:  sayingID,
		CommentID: commentID,
	}
	err = db.Create(myComment).Error
	return
}

func DelMyCommentsOfSaying(tids []string) (err error) {
	err = db.Unscoped().Where("saying_id in (?)", tids).Delete(&MyComment{}).Error
	return
}
