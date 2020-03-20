package comment

import (
	"myzone/model"
	"time"
)

func AfterAddNewComment(comment *model.Comment, sayingID int) {
	sayingInfo, _ := model.GetSayingById(sayingID)
	updateSaying := model.Saying{
		LastDate:      time.Now(),
		CommentsCnt:   sayingInfo.CommentsCnt + 1,
		LastUserID:    comment.UserID,
		LastCommentID: int(comment.ID),
	}
	model.UpdateSaying(sayingID, updateSaying)
	model.AddMyComment(comment.UserID, sayingID, int(comment.ID))
	oldUserInfo, _ := model.GetUserByID(comment.UserID)
	model.UpdateUserCommentsCnt(comment.UserID, oldUserInfo.CommentsCnt+1)
}
