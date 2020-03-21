package model

type Group struct {
	Model
	Name            string `gorm:"default:''" json:"name"`           //用户组名称
	Allowread       int    `gorm:"default:0" json:"allowread"`       //允许访问
	Allowarticle    int    `gorm:"default:0" json:"allowarticle"`    //允许发文章
	Allowsaying     int    `gorm:"default:0" json:"allowsaying"`     //允许发说说
	Allowreply      int    `gorm:"default:0" json:"allowreply"`      //允许回复
	Allowcomment    int    `gorm:"default:0" json:"allowcomment"`    //允许评论
	Allowattach     int    `gorm:"default:0" json:"allowattach"`     //允许上传文件
	Allowdown       int    `gorm:"default:0" json:"allowdown"`       //允许下载文件
	Allowupdate     int    `gorm:"default:0" json:"allowupdate"`     //允许编辑
	Allowdelete     int    `gorm:"default:0" json:"allowdelete"`     //
	Allowdeleteuser int    `gorm:"default:0" json:"allowdeleteuser"` //
	Allowviewip     int    `gorm:"default:0" json:"allowviewip"`     //允许查看用户敏感信息
}

func GetUserGroupList() (glist []Group, err error) {
	err = db.Model(&Group{}).Find(&glist).Error
	return
}
