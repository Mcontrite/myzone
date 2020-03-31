package router

import (
	"html/template"
	adminservice "myzone/controller/admin"
	apiservice "myzone/controller/api"
	webservice "myzone/controller/web"
	"myzone/middleware/auth"
	"myzone/middleware/cros"
	"myzone/middleware/jwt"
	"myzone/middleware/loger"
	"myzone/middleware/xss"
	"myzone/package/setting"
	"net/http"

	limit "github.com/aviddiviner/gin-limit"
	"github.com/gin-gonic/gin"
	sessions "github.com/tommy351/gin-sessions"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cros.Cors())
	r.Use(limit.MaxAllowed(100))
	// 引入session
	store := sessions.NewCookieStore([]byte("secret123"))
	r.Use(sessions.Middleware("my_session", store))
	r.Use(loger.LoggerToFile())
	gin.SetMode(setting.ServerSetting.RunMode)
	// 模板函数
	r.SetFuncMap(template.FuncMap{
		"unescaped":   unescaped,
		"strtime":     StrTime,
		"plus1":       selfPlus,
		"numplusplus": numPlusPlus,
		"strip":       Long2IPString,
		"truncate":    Truncate,
	})
	// 避免404
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", gin.H{})
	})
	r.NoMethod(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", gin.H{})
	})
	r.LoadHTMLGlob("views/*/**/***")
	// 推荐使用绝对路径 相当于简历了软连接--快捷方式
	r.StaticFS("/static", http.Dir("./static"))
	r.StaticFS("/upload", http.Dir("./upload"))
	// 用户前端页面
	web := r.Group("")
	{
		// 首页
		web.GET("/", webservice.SayingIndex)
		web.GET("/articles", webservice.Index)
		web.GET("/sayings", webservice.SayingIndex)
		web.GET("/contactme.html", webservice.ContactMe)
		// 注册页
		web.GET("/register.html", webservice.Register)
		// 登录页
		web.GET("/login.html", webservice.Login)
		// 登出页面
		web.GET("/logout", apiservice.UserLogout)
		// 文章：新建页
		web.GET("/newarticle.html", webservice.NewArticle)
		web.GET("/newsaying.html", webservice.NewSaying)
		// 文章：编辑页
		web.GET("/article/:id/edit.html", webservice.EditArticle)
		web.GET("/saying/:id/edit.html", webservice.EditSaying)
		// 文章：详情页
		web.GET("/article/:id/detail.html", webservice.Article)
		web.GET("/saying/:id/detail.html", webservice.Saying)
		// 高级回复
		web.GET("/article/:id/areply.html", webservice.ArticleAddReply)
		web.GET("/saying/:id/acomment.html", webservice.SayingAddComment)
		// 我的信息
		web.GET("/my.html", webservice.MyInfo)
		// 修改信息
		web.GET("/my_edit.html", webservice.MyEdit)
		// 我的文章说说列表
		web.GET("/my_article.html", webservice.MyArticle)
		web.GET("/my_saying.html", webservice.MySaying)
		// 我的收藏列表
		web.GET("/my_favorite.html", webservice.MyFavorite)
		// 我的回复列表
		web.GET("/my_reply.html", webservice.MyReply)
		web.GET("/my_comment.html", webservice.MyComment)
		// 查看其它用户内容
		web.GET("/user/:id/info.html", webservice.UserInfo)
		web.GET("/user/:id/article.html", webservice.UserArticle)
		web.GET("/user/:id/saying.html", webservice.UserSaying)
		web.GET("/user/:id/reply.html", webservice.UserReply)
		web.GET("/user/:id/comment.html", webservice.UserComment)
		// 前台管理员进行操作的模态框
		web.GET("/mod/article/delete.html", webservice.DeleteMod)
		web.GET("/mod/saying/delete.html", webservice.DeleteMod)
	}
	// 数据操作的接口
	api := r.Group("/api")
	{
		// 检测用户名是否被使用
		api.GET("/checkname", apiservice.CheckNameUsed)
		// 获取某用户
		api.GET("/user", apiservice.GetUser)
		//注册
		api.POST("/user", apiservice.AddUser)
		api.POST("/register", apiservice.AddUser)
		// 登录
		api.POST("/login", apiservice.UserLogin)
		// 登出操作
		api.GET("/logout", apiservice.UserLogout)
		// 发送重设密码的邮件
		api.POST("/password/reset", apiservice.UserResetPassword)
		// 刷新token
		api.GET("/token", apiservice.RefreshToken)
		// 更新用户
		api.PUT("/user/:id", apiservice.EditUser)
		// 删除用户
		api.DELETE("/user/:id", apiservice.DeleteUser)
		// 用户：重设密码
		api.POST("/user/:id/password/reset", apiservice.ResetUserPassword)
		// 用户：重设头像
		api.POST("/user/:id/avatar/reset", apiservice.ResetUserAvatar)
		// 用户：重设用户名
		api.POST("/user/:id/name/reset", apiservice.ResetUserName)
		// 文章：发表
		api.POST("/article", xss.XSS(), apiservice.AddArticle)
		api.POST("/saying", xss.XSS(), apiservice.AddSaying)
		// 文章：发表
		api.POST("/article/:id/favourite", apiservice.Addarticlefavourite)
		// 文章：删除
		api.POST("/article/:id/delete", apiservice.DeleteArticles)
		api.POST("/saying/:id/delete", apiservice.DeleteSayings)
		// 文章：修改
		api.POST("/article/:id/update", apiservice.UpdateArticle)
		api.POST("/saying/:id/update", apiservice.UpdateSaying)
		// 添加评论
		api.POST("/article/:id/reply", apiservice.AddReply)
		api.POST("/saying/:id/comment", apiservice.AddComment)
		// 添加附件
		api.POST("/article/:id/attach/add", apiservice.AddarticleAttach)
		// 删除附件
		api.POST("/article/:id/attach/del", apiservice.DelarticleAttach)
		// 评论的相关操作
		api.POST("/reply/:id/update", apiservice.UpdateReply)
		api.POST("/comment/:id/update", apiservice.UpdateComment)
		// 获取验证码
		api.GET("/capacha", apiservice.GetCapacha)
		api.POST("/capacha", apiservice.VerfiyCaptcha)
		// 上传图片
		api.POST("/image/upload", apiservice.CkeditorUpload)
		api.POST("/attach/upload", apiservice.UploadAttach)
		api.POST("/attach/add", apiservice.UploadAddAttach)
		api.POST("/attach/delete", apiservice.DeleteAttach)
	}
	// 管理员页面
	admin := r.Group("/admin")
	admin.Use(auth.AUTH())
	{
		// 登录展示页
		admin.GET("/login.html", adminservice.AdminLogin)
		// 管理员二次登录验证
		admin.POST("/login", adminservice.AdminLoginCheck)
	}
	admin.Use(jwt.JWT())
	{
		// 后台首页
		admin.GET("/index.html", adminservice.AdminIndex)
		admin.GET("/user/list.html", adminservice.AdminUserList)
		admin.GET("/user/group.html", adminservice.AdminGroupList)
		admin.GET("/user/create.html", adminservice.AdminUserCreate)
		admin.POST("/user/add", adminservice.AdminUserAdd)
		admin.GET("/user/edit.html", adminservice.AdminUserEdit)
		admin.POST("/user/update", adminservice.AdminUserUpdate)
		admin.GET("/article/list.html", adminservice.GetArticleList)
		admin.GET("/saying/list.html", adminservice.GetSayingList)
	}
	return r
}
