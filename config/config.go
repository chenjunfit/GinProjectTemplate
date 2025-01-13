package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	TimeFormat = "2006-01-02 15:04:02"
)

var (
	Port          string
	JwtSignKey    string
	JwtExpireTime int64
	AdminUserName string
	AdminPassword string
)

type ReturnData struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func NewReturnData() ReturnData {
	return ReturnData{
		Status: 200,
		Data:   make(map[string]interface{}),
	}
}

func init() {
	viper.SetDefault("LOG_LEVEL", "debug")
	viper.SetDefault("PORT", ":8080")
	viper.SetDefault("JWT_SIGN_KEY", "tiantianmoyu")
	viper.SetDefault("JWT_EXPIRE_TIME", 120)
	//admin,adminpassword
	viper.SetDefault("ADMIN_USER_NAME", "21232F297A57A5A743894A0E4A801FC3")
	viper.SetDefault("ADMIN_PASSWORD", "E3274BE5C857FB42AB72D786E281B4B8")
	viper.AutomaticEnv()
	logLevel := viper.GetString("LOG_LEVEL")
	Port = viper.GetString("PORT")
	JwtSignKey = viper.GetString("JWT_SIGN_KEY")
	JwtExpireTime = viper.GetInt64("JWT_EXPIRE_TIME")
	AdminUserName = viper.GetString("ADMIN_USER_NAME")
	AdminPassword = viper.GetString("ADMIN_PASSWORD")
	initLogConfig(logLevel)
}
func initLogConfig(level string) {
	if level == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.DebugLevel)
	}
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: TimeFormat})
}
