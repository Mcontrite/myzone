package v1

import (
	"myzone/model"
	"myzone/package/app"
	"myzone/package/rcode"
	"myzone/package/session"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Addarticlefavourite(c *gin.Context) {
	tid, _ := strconv.Atoi(c.DefaultPostForm("articleid", "1"))
	uid, _ := strconv.Atoi(session.GetSession(c, "userid"))
	code := rcode.SUCCESS
	articleInfo, err := model.GetArticleById(tid)
	if err != nil {
		code = rcode.ERROR_UNFIND_DATA
		app.JsonErrResponse(c, code)
		return
	}
	isfav, _ := model.CheckFavourite(uid, tid)
	action := 0
	favNum := articleInfo.FavouriteCnt
	if isfav == 0 {
		model.AddMyFavourite(uid, tid)
		model.UpdateArticleFavouriteCnt(tid, articleInfo.FavouriteCnt+1)
		action = 1
		favNum++
	} else {
		model.DelMyFavourite(uid, tid)
		if articleInfo.FavouriteCnt > 0 {
			model.UpdateArticleFavouriteCnt(tid, articleInfo.FavouriteCnt-1)
		}
		action = 0
		favNum--
	}
	app.JsonOkResponse(c, code, map[string]interface{}{"action": action, "fav_num": favNum})
}
