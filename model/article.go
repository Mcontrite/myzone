package model

import "time"

type Article struct {
	Model
	UserID       int       `gorm:"default:0" json:"user_id"`        //
	Title        string    `gorm:"default:''" json:"title"`         // 文章
	ViewsCnt     int       `gorm:"default:0" json:"views_cnt"`      //查看次数, 剥离出去，单独的服务，避免 cache 失效
	ReplysCnt    int       `gorm:"default:0" json:"replys_cnt"`     //回复数
	FavouriteCnt int       `gorm:"default:0" json:"favourite_cnt"`  //被收藏数
	FilesNum     int       `gorm:"default:0" json:"files_num"`      //附件中包含的文件数
	LastDate     time.Time `json:"last_date"`                       //最后回复时间
	FirstReplyID int       `gorm:"default:0" json:"first_reply_id"` //首贴 rid
	LastReplyID  int       `gorm:"default:0" json:"last_reply_id"`  //最后回复的 rid
	LastUserID   int       `gorm:"default:0" json:"last_user_id"`   //最近参与的 uid
	UserIP       string    `gorm:"default:''" json:"userip"`        //发帖时用户ip
	User         User      `json:"user"`
	LastUser     User
	Attach       []Attach
}

func GetArticles(whereMap interface{}, order string, limit int, page int) (article []Article, err error) {
	err = db.Model(&Article{}).Preload("User").Where(whereMap).Order(order).Offset((page - 1) * limit).Limit(limit).Find(&article).Error
	return
}

func GetArticleList(page int) (articles []Article, err error) {
	if page <= 1 {
		page = 1
	}
	err = db.Preload("User").Model(&Article{}).Order("updated_at desc").Offset((page - 1) * PAGE_SIZE).Limit(PAGE_SIZE).Find(&articles).Error
	return
}

func GetArticleTotleCount() (totle int, err error) {
	err = db.Model(&Article{}).Count(&totle).Error
	return
}

func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

func GetArticleById(id int) (article Article, err error) {
	err = db.Preload("User").Where("id = ?", id).Model(&Article{}).First(&article).Error
	return
}

func AddArticle(article *Article) (*Article, error) {
	err := db.Model(&Article{}).Create(article).Error
	return article, err
}

func UpdateArticle(id int, article Article) (uparticle Article, err error) {
	err = db.Model(&Article{}).Where("id = ?", id).Updates(article).Error
	uparticle, err = GetArticleById(id)
	return
}

func UpdateArticlePro(id int, items map[string]interface{}) (uparticle Article, err error) {
	err = db.Model(&Article{}).Where("id = ?", id).Updates(items).Error
	uparticle, err = GetArticleById(id)
	return
}

func DelArticle(ids []string) (err error) {
	err = db.Unscoped().Where("id in (?)", ids).Delete(&Article{}).Error
	return
}

func UpdateArticleViewsCnt(id int) error {
	article, _ := GetArticleById(id)
	return db.Model(&Article{}).Where("id = ?", id).Update("views_cnt", article.ViewsCnt+1).Error
}

func UpdateArticleFilesNum(id int, num int) error {
	return db.Model(&Article{}).Where("id = ?", id).Update("files_num", num).Error
}

func UpdateArticleFavouriteCnt(id int, num int) error {
	return db.Model(&Article{}).Where("id = ?", id).Update("favourite_cnt", num).Error
}

func CountArticlesNum() (articlesNum int, err error) {
	err = db.Model(&Article{}).Count(&articlesNum).Error
	return
}

func GetArticlesByIDs(ids []string) (articles []*Article, err error) {
	err = db.Model(&Article{}).Preload("User").Where("id in (?)", ids).Find(&articles).Error
	return
}
