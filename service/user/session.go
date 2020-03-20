package user

import (
	"myzone/model"
	"myzone/package/session"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserSession struct {
	Username       string `json:"username"`
	Userid         int    `json:"userid"`
	Useravatar     string `json:"useravatar"`
	Userarticlecnt int    `json:"userarticlecnt"`
	Userreplycnt   int    `json:"userreplycnt"`
	Usersayingcnt  int    `json:"usersayingcnt"`
	Usercommentcnt int    `json:"usercommentcnt"`
	Isadmin        string `json:"isadmin"`
}

func LoginSession(c *gin.Context, user model.User, sok chan int) {
	session.SetSession(c, "username", user.Username)
	session.SetSession(c, "userid", strconv.Itoa(int(user.ID)))
	session.SetSession(c, "useravatar", user.Avatar)
	session.SetSession(c, "userarticlecnt", strconv.Itoa(user.ArticlesCnt))
	session.SetSession(c, "userreplycnt", strconv.Itoa(user.ReplysCnt))
	session.SetSession(c, "usersayingcnt", strconv.Itoa(user.SayingsCnt))
	session.SetSession(c, "usercommentcnt", strconv.Itoa(user.CommentsCnt))
	session.SetSession(c, "isadmin", IsAdmin(user.GroupID))
	sok <- 1
}

func GetSessions(c *gin.Context) (sessions *UserSession) {
	username := session.GetSession(c, "username")
	userid, _ := strconv.Atoi(session.GetSession(c, "userid"))
	useravatar := session.GetSession(c, "useravatar")
	userarticlecnt, _ := strconv.Atoi(session.GetSession(c, "userarticlecnt"))
	userreplycnt, _ := strconv.Atoi(session.GetSession(c, "userreplycnt"))
	usersayingcnt, _ := strconv.Atoi(session.GetSession(c, "usersayingcnt"))
	usercommentcnt, _ := strconv.Atoi(session.GetSession(c, "usercommentcnt"))
	isadmin := session.GetSession(c, "isadmin")
	sessions = &UserSession{
		Username:       username,
		Userid:         userid,
		Useravatar:     useravatar,
		Userarticlecnt: userarticlecnt,
		Userreplycnt:   userreplycnt,
		Usersayingcnt:  usersayingcnt,
		Usercommentcnt: usercommentcnt,
		Isadmin:        isadmin,
	}
	return
}

func LogoutSession(c *gin.Context) {
	session.DeleteSession(c, "username")
	session.DeleteSession(c, "userid")
	session.DeleteSession(c, "useravatar")
	session.DeleteSession(c, "userarticlecnt")
	session.DeleteSession(c, "userreplycnt")
	session.DeleteSession(c, "usersayingcnt")
	session.DeleteSession(c, "usercommentcnt")
	session.DeleteSession(c, "isadmin")
}

func IsLogin(c *gin.Context) (res bool) {
	username := session.GetSession(c, "username")
	if len(username) > 0 {
		res = true
	} else {
		res = false
	}
	return
}
