package model

import "time"

type Cache struct {
	Model
	K      string    `json:"k"`      //
	V      string    `json:"v"`      //
	Expiry time.Time `json:"expiry"` //
}
