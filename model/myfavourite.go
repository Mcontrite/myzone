package model

type MyFavourite struct {
	Model
	UserID    int `gorm:"default:0" json:"user_id"`    //
	ArticleID int `gorm:"default:0" json:"article_id"` //
	User      User
	Article   Article
}

func GetMyFavouriteList(uid int, page int, limit int, Orderby string) (articles []MyFavourite, err error) {
	if page <= 1 {
		page = 1
	}
	if limit == 0 {
		limit = PAGE_SIZE
	}
	err = db.Preload("User").Preload("Article").Model(&MyArticle{}).Where("user_id = ?", uid).Offset((page - 1) * limit).Limit(limit).Order(Orderby).Find(&articles).Error
	return
}

func AddMyFavourite(userID int, articleID int) (myFavourite *MyFavourite, err error) {
	myFavourite = &MyFavourite{
		UserID:    userID,
		ArticleID: articleID,
	}
	err = db.Create(myFavourite).Error
	return
}

func DelMyFavourite(uid int, articleId int) error {
	return db.Unscoped().Where("user_id = ?", uid).Where("article_id = ?", articleId).Delete(&MyFavourite{}).Error
}

func CheckFavourite(uid int, tid int) (fav int, err error) {
	err = db.Model(&MyFavourite{}).Where("user_id = ? and article_id = ?", uid, tid).Count(&fav).Error
	return
}

func DelMyFavouritesOfArticle(tids []string) (err error) {
	err = db.Unscoped().Where("article_id in (?)", tids).Delete(&MyFavourite{}).Error
	return
}
