package model

type MyArticle struct {
	Model
	UserID    int `gorm:"default:0" json:"user_id"`
	ArticleID int `gorm:"default:0" json:"article_id"`
	User      User
	Article   Article
}

func GetMyArticleList(uid int, page int, limit int, Orderby string) (articles []MyArticle, err error) {
	if page <= 1 {
		page = 1
	}
	if limit == 0 {
		limit = 20
	}
	err = db.Preload("User").Preload("Article").Model(&MyArticle{}).Where("user_id = ?", uid).Offset((page - 1) * 20).Limit(200).Order(Orderby).Find(&articles).Error
	return
}

func AddMyArticle(userID int, articleID int) (myArticle *MyArticle, err error) {
	myArticle = &MyArticle{
		UserID:    userID,
		ArticleID: articleID,
	}
	err = db.Create(myArticle).Error
	return
}

func DelMyArticlesOfArticle(tids []string) (err error) {
	err = db.Unscoped().Where("article_id in (?)", tids).Delete(&MyArticle{}).Error
	return
}
