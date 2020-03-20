package web

import (
	"github.com/gin-gonic/gin"
)

func DeleteMod(c *gin.Context) {
	c.HTML(200, "del_mod.html", gin.H{})
}
