package article

import (
	"myzone/model"
)

func GetUserArticles(uid int) (articles []model.Article, err error) {
	whereMap := &model.Article{UserID: uid}
	order := "created_at desc"
	limit := 10
	articles, err = model.GetArticles(whereMap, order, limit, 1)
	return
}

func AfterAddNewArticle(article *model.Article) {
	articleID := article.ID
	userID := article.UserID
	model.AddMyArticle(userID, int(articleID))
	oldUserInfo, _ := model.GetUserByID(article.UserID)
	model.UpdateUserArticlesCnt(article.UserID, oldUserInfo.ArticlesCnt+1)
}

func DelArticles(tids []string) (err error) {
	err = model.DelReplysOfArticle(tids)
	if err != nil {
		return
	}
	err = model.DelMyArticlesOfArticle(tids)
	if err != nil {
		return
	}
	err = model.DelMyReplysOfArticle(tids)
	if err != nil {
		return
	}
	err = model.DelMyFavouritesOfArticle(tids)
	if err != nil {
		return
	}
	err = model.DelAttachsOfArticle(tids)
	if err != nil {
		return
	}
	err = model.DelArticle(tids)
	if err != nil {
		return
	}
	return
}

func GetArticlesByIDs(tidArr []string) (articles []*model.Article, err error) {
	articles, err = model.GetArticlesByIDs(tidArr)
	return
}
