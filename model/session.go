package model

type Session struct {
	Model
	UserId    int    `gorm:"default:0" json:"user_id"`    //用户id 未登录为 0，可以重复
	URL       string `gorm:"default:''" json:"url"`       //当前访问 url
	IP        int    `gorm:"default:0" json:"ip"`         //
	Useragent string `gorm:"default:''" json:"useragent"` //
	Data      string `gorm:"default:''" json:"data"`      //
	Bigdata   int    `gorm:"default:0" json:"bigdata"`    //是否有大数据
	LastDate  int    `gorm:"default:0" json:"last_date"`  //上次活动时间
}
