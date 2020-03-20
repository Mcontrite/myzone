package admin

import (
	"myzone/model"

	"github.com/gin-gonic/gin"
)

func GetSayingList(c *gin.Context) {
	var sayings []model.Saying
	sayings, _ = model.GetSayings(map[string]interface{}{"deleted_at": nil}, "created_at desc", 100, 1)
	c.HTML(200, "asayinglist.html", gin.H{"sayings": sayings})
}
