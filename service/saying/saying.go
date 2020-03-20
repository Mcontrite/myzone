package saying

import (
	"myzone/model"
)

func GetUserSayings(uid int) (sayings []model.Saying, err error) {
	whereMap := &model.Saying{UserID: uid}
	order := "created_at desc"
	limit := 10
	sayings, err = model.GetSayings(whereMap, order, limit, 1)
	return
}

func AfterAddNewSaying(saying *model.Saying) {
	sayingID := saying.ID
	userID := saying.UserID
	model.AddMySaying(userID, int(sayingID))
}

func DelSayings(tids []string) (err error) {
	err = model.DelCommentsOfSaying(tids)
	if err != nil {
		return
	}
	err = model.DelMySayingsOfSaying(tids)
	if err != nil {
		return
	}
	err = model.DelMyCommentsOfSaying(tids)
	if err != nil {
		return
	}
	err = model.DelSaying(tids)
	if err != nil {
		return
	}
	return
}

func GetSayingsByIDs(tidArr []string) (sayings []*model.Saying, err error) {
	sayings, err = model.GetSayingsByIDs(tidArr)
	return
}
