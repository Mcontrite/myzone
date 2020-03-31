package setting

import (
	"log"
	"strings"
	"time"

	"github.com/go-ini/ini"
)

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Redis struct {
	RedisHost        string
	RedisPassword    string
	RedisMaxidle     int
	RedisMaxActive   int
	RedisIdleTimeout time.Duration
}

var RedisSetting = &Redis{}

type Image struct {
	ImageSavePath   string
	ImageMaxSize    int
	ImageAlloweXts  string
	ImageAllowExts  []string
	RuntimeRootPath string
}

var ImageSetting = &Image{}

type Server struct {
	RunMode        string
	HttpPort       int
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	LogPath        string
	PageSize       int
	JwtSecret      string
	ViewUrl        string
	UploadUrl      string
	UploadPath     string
	LogoMobileUrl  string
	LogoPcUrl      string
	LogoWaterUrl   string
	Sitepre        string
	Siteurl        string
	Sitename       string
	Siteseoword    string
	Timezone       string
	Lang           string
	Runlevel       int
	RunlevelReason string

	CookieDomain string
	CookiePath   string

	ReplylistPagesize     int
	CacheArticleListPages int
	SessionDelayUpdate    int
	UploadImageWidth      int
	OrderDefault          string
	AttachDirSaveRule     string

	UpdateViewsOn  int
	UserCreateOn   int
	UserResetpwOn  int
	CdnOn          int
	UrlRewriteOn   int
	DisabledPlugin int
	Version        string
	StaticVersion  string
	Installed      int

	RuntimeRootPath string
	LogSavePath     string
	LogSaveName     string
	TimeFormat      string
	LogFileExt      string
}

var ServerSetting = &Server{}

var cfg *ini.File

func init() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)
	mapTo("image", ImageSetting)
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	RedisSetting.RedisIdleTimeout = RedisSetting.RedisIdleTimeout * time.Second
	ImageSetting.ImageAllowExts = strings.Split(ImageSetting.ImageAlloweXts, ",")
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}

func UpdateItemValue(section, key, value string) {
	cfg.Section(section).Key(key).SetValue(value)
	cfg.SaveTo("conf/app.ini")
}
