package model

type Attach struct {
	Model
	UserID       int    `gorm:"default:0" json:"user_id"`       //用户id
	ArticleID    int    `gorm:"default:0" json:"article_id"`    //文章id
	ReplyID      int    `gorm:"default:0" json:"reply_id"`      //回复id
	Comment      string `gorm:"default:''" json:"comment"`      //文件注释 方便于搜索
	Filename     string `gorm:"default:''" json:"filename"`     //文件名称，会过滤，并且截断，保存后的文件名，不包含URL前缀 upload_url
	Orgfilename  string `gorm:"default:''" json:"orgfilename"`  //上传的原文件名
	Filetype     string `gorm:"default:''" json:"filetype"`     //image/txt/zip，小图标显示
	Filesize     int    `gorm:"default:0" json:"filesize"`      //文件尺寸，单位字节
	Width        int    `gorm:"default:0" json:"width"`         //width > 0 则为图片
	Height       int    `gorm:"default:0" json:"height"`        //
	DownloadsNum int    `gorm:"default:0" json:"downloads_num"` //下载次数
	Isimage      int    `gorm:"default:0" json:"isimage"`       //是否为图片
}

func AddAttach(attach *Attach) (*Attach, error) {
	err := db.Model(&Attach{}).Create(attach).Error
	return attach, err
}

func GetAttachsByReplyId(replyId int) (attachs []Attach, err error) {
	err = db.Model(&Attach{}).Where("reply_id = ?", replyId).Find(&attachs).Error
	return
}

func DelAttach(id int) error {
	return db.Model(&Attach{}).Where("id = ?", id).Unscoped().Delete(&Attach{}).Error
}

func CountAttachsNum() (accachsNum int, err error) {
	err = db.Model(&Attach{}).Count(&accachsNum).Error
	return
}

func DelAttachsOfArticle(tids []string) (err error) {
	err = db.Unscoped().Where("article_id in (?)", tids).Delete(&Attach{}).Error
	return
}
