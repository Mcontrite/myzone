package reply

import (
	"myzone/model"
	"time"
)

func AfterAddNewReply(reply *model.Reply, articleID int) {
	articleInfo, _ := model.GetArticleById(articleID)
	updateArticle := model.Article{
		LastDate:    time.Now(),
		ReplysCnt:   articleInfo.ReplysCnt + 1,
		LastUserID:  reply.UserID,
		LastReplyID: int(reply.ID),
	}
	model.UpdateArticle(articleID, updateArticle)
	model.AddMyReply(reply.UserID, articleID, int(reply.ID))
	oldUserInfo, _ := model.GetUserByID(reply.UserID)
	model.UpdateUserReplysCnt(reply.UserID, oldUserInfo.ReplysCnt+1)
}
