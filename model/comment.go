package model

type Comment struct {
	Model
	UserID         int    `gorm:"default:0" json:"user_id"`          //
	SayingID       int    `gorm:"default:0" json:"saying_id"`        //主题id
	QuoteCommentId int    `gorm:"default:0" json:"quote_comment_id"` //引用pid，可能不存在
	Message        string `gorm:"default:''" json:"message"`         //内容，用户提示的原始数据
	MessageFmt     string `gorm:"default:''" json:"message_fmt"`     //过滤后的html内容
	Isfirst        int    `gorm:"default:0" json:"isfirst"`          //是否为首帖
	UserIP         string `gorm:"default:''" json:"userip"`          //发帖时用户ip
	User           User
	Saying         Saying
}

func GetCommentById(id int) (comment Comment, err error) {
	err = db.Model(&Comment{}).Where("id = ?", id).First(&comment).Error
	return
}

func GetSayingFirstCommentByTid(tid int) (comment Comment, err error) {
	err = db.Model(&Comment{}).Where("saying_id = ?", tid).Where("isfirst = ?", 1).First(&comment).Error
	return
}

func GetSayingCommentListByTid(tid int, limit int, page int) (comment []Comment, err error) {
	err = db.Preload("User").Preload("User.Group").Preload("Attach").Model(&Comment{}).Where("saying_id = ?", tid).Where("isfirst = ?", 0).Offset((page - 1) * limit).Limit(limit).Find(&comment).Error
	return
}

func AddComment(comment *Comment) (*Comment, error) {
	err := db.Create(comment).Error
	return comment, err
}

func UpdateComment(id int, comment Comment) (upComment Comment, err error) {
	err = db.Model(&Comment{}).Where("id = ?", id).Updates(comment).Error
	upComment, err = GetCommentById(id)
	return
}

func CountCommentNum() (commentNum int, err error) {
	err = db.Model(&Comment{}).Where("isfirst = ?", 0).Count(&commentNum).Error
	return
}

func DelCommentsOfSaying(tids []string) (err error) {
	err = db.Unscoped().Where("saying_id in (?)", tids).Delete(&Comment{}).Error
	return
}
