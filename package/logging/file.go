package logging

import (
	"fmt"
	"myzone/package/setting"
	"time"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.ServerSetting.RuntimeRootPath, setting.ServerSetting.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.ServerSetting.LogSaveName,
		time.Now().Format(setting.ServerSetting.TimeFormat),
		setting.ServerSetting.LogFileExt,
	)
}
