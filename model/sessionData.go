package model

import "time"

type SessionData struct {
	SessionId string    `gorm:"primary_key" json:" - "` //
	LastDate  time.Time `json:"last_date"`              //
	Data      string    `json:"data"`                   //
}
