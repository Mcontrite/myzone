package model

type MyReply struct {
	Model
	UserID    int `gorm:"primary_key;default:0" json:"user_id"`  //
	ArticleID int `gorm:"default:0" json:"article_id"`           //
	ReplyID   int `gorm:"primary_key;default:0" json:"reply_id"` //
	User      User
	Article   Article
	Reply     Reply
}

func GetMyReplyList(uid int, page int, limit int, Orderby string) (replys []MyReply, err error) {
	if page <= 1 {
		page = 1
	}
	if limit == 0 {
		limit = 20
	}
	err = db.Preload("Article").Preload("Article.User").Preload("Article.User.Group").Preload("Reply").Model(&MyReply{}).Where("user_id = ?", uid).Offset((page - 1) * 20).Limit(200).Order(Orderby).Find(&replys).Error
	return
}

func AddMyReply(userID int, articleID int, replyID int) (myReply *MyReply, err error) {
	myReply = &MyReply{
		UserID:    userID,
		ArticleID: articleID,
		ReplyID:   replyID,
	}
	err = db.Create(myReply).Error
	return
}

func DelMyReplysOfArticle(tids []string) (err error) {
	err = db.Unscoped().Where("article_id in (?)", tids).Delete(&MyReply{}).Error
	return
}
