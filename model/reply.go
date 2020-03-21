package model

type Reply struct {
	Model
	UserID       int    `gorm:"default:0" json:"user_id"`        //用户ID
	ArticleID    int    `gorm:"default:0" json:"article_id"`     //文章id
	QuoteReplyId int    `gorm:"default:0" json:"quote_reply_id"` //引用哪个Aid，可能不存在
	Message      string `gorm:"default:''" json:"message"`       //内容原始数据
	MessageFmt   string `gorm:"default:''" json:"message_fmt"`   //过滤后的html内容
	Isfirst      int    `gorm:"default:0" json:"isfirst"`        //是否为首帖
	ImagesNum    int    `gorm:"default:0" json:"images"`         //附件中包含的图片数
	FilesNum     int    `gorm:"default:0" json:"files"`          //附件中包含的文件数
	Doctype      int    `gorm:"default:0" json:"doctype"`        //类型，0: html, 1: txt; 2: markdown; 3: ubb
	UserIP       string `gorm:"default:''" json:"userip"`        //发帖时用户ip
	User         User
	Article      Article
	Attach       []Attach
}

func GetReplyById(id int) (reply Reply, err error) {
	err = db.Model(&Reply{}).Where("id = ?", id).First(&reply).Error
	return
}

func GetArticleFirstReplyByTid(tid int) (reply Reply, err error) {
	err = db.Model(&Reply{}).Where("article_id = ?", tid).Where("isfirst = ?", 1).First(&reply).Error
	return
}

func GetArticleReplyListByTid(tid int, limit int, page int) (reply []Reply, err error) {
	err = db.Preload("User").Preload("User.Group").Preload("Attach").Model(&Reply{}).Where("article_id = ?", tid).Where("isfirst = ?", 0).Offset((page - 1) * limit).Limit(limit).Find(&reply).Error
	return
}

func AddReply(reply *Reply) (*Reply, error) {
	err := db.Create(reply).Error
	return reply, err
}

func UpdateReply(id int, reply Reply) (upReply Reply, err error) {
	err = db.Model(&Reply{}).Where("id = ?", id).Updates(reply).Error
	upReply, err = GetReplyById(id)
	return
}

func UpdateReplyFilesNum(id int, num int) error {
	return db.Model(&Reply{}).Where("id = ?", id).Update("files_num", num).Error
}

func CountReplyNum() (replyNum int, err error) {
	err = db.Model(&Reply{}).Where("isfirst = ?", 0).Count(&replyNum).Error
	return
}

func DelReplysOfArticle(tids []string) (err error) {
	err = db.Unscoped().Where("article_id in (?)", tids).Delete(&Reply{}).Error
	return
}
