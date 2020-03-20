package utils

import (
	"math"
	"myzone/package/setting"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.ServerSetting.PageSize
	}
	return result
}

func Pagination_tpl(url, text, active string) string {
	g_pagination_tpl := "<li class='page-item{active}'><a href='{url}' class='page-link'>{text}</a></li>"
	g_pagination_tpl = strings.Replace(g_pagination_tpl, "{url}", url, 1)
	g_pagination_tpl = strings.Replace(g_pagination_tpl, "{text}", text, 1)
	g_pagination_tpl = strings.Replace(g_pagination_tpl, "{active}", active, 1)
	return g_pagination_tpl
}

//bootstrap 翻页，命名与 bootstrap 保持一致
func Pagination(url string, totalnum int, page int, pagesize int) (s string) {
	if pagesize == 0 {
		pagesize = 20
	}
	totalpage := math.Ceil(float64(totalnum) / float64(pagesize))
	if totalpage < 2 {
		return
	}
	page = int(math.Min(float64(totalpage), float64(page)))
	shownum := 5
	start := int(math.Max(1, float64(page-shownum)))
	end := int(math.Max(totalpage, float64(page+shownum)))
	// 不足 $shownum，补全左右两侧
	right := page + shownum - int(totalpage)
	if right > 0 {
		start -= right
		start = int(math.Max(1, float64(start)))
	}
	left := page - shownum
	if left < 0 {
		end -= left
		end = int(math.Min(totalpage, float64(end)))
	}
	if page != 1 {
		url := strings.Replace(url, "{page}", strconv.Itoa(page-1), 1)
		s += Pagination_tpl(url, "◀", "")
	}
	if start > 1 {
		text := "1 "
		if start > 2 {
			text += "..."
		}
		url := strings.Replace(url, "{page}", "1", 1)
		s += Pagination_tpl(url, text, "")
	}
	for i := start; i <= end; i++ {
		text := ""
		if i == page {
			text += " active"
		}
		url := strings.Replace(url, "{page}", strconv.Itoa(i), 1)
		s += Pagination_tpl(url, strconv.Itoa(i), text)
	}
	if end != int(totalpage) {
		text := ""
		if (int(totalpage) - end) > 1 {
			text = "..." + strconv.Itoa(int(totalpage))
		} else {
			text = strconv.Itoa(int(totalpage))
		}
		url := strings.Replace(url, "{page}", strconv.Itoa(int(totalpage)), 1)
		s += Pagination_tpl(url, text, "")
	}
	if page != int(totalpage) {
		url := strings.Replace(url, "{page}", strconv.Itoa(page+1), 1)
		s += Pagination_tpl(url, "▶", "")
	}
	return
}
