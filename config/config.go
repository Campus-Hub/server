package config

import (
	"os"

	"github.com/Campus-Hub/server/pkg/logger"

	"github.com/golang/glog"
	"github.com/jinzhu/configor"
)

// Setup Configuration
// FIXME configor cannot read the relative path
func Setup() {
	var appPath, err = os.Getwd()
	if err != nil {
		logger.Logger.Error("Configuration Setup Error: ", err)
	}
	err = configor.New(&configor.Config{
		//Debug:                true,
		//Verbose:              true,
		ErrorOnUnmatchedKeys: true,
	}).Load(&Config, appPath+"/config/config.toml")
	if err != nil {
		glog.Error("Configuration Reading Error:", err)
	}
}

// Config Global configurations Variable
var Config = struct {
	AppMode string `default:"debug"`
	AppName string `default:"campus-hub-server"`

	Server struct {
		ReadTimeout  int
		WriteTimeout int
	}

	DB struct {
		Name     string
		User     string
		Password string
		Port     uint
		Type     string
	}

	Redis struct {
		Host        string
		Port        int
		Password    string
		DbName      int
		MaxIdle     int
		MaxActive   int
		IdleTimeout int
	}

	RabbitMQ struct {
		RabbitMQ         string
		RabbitMQUser     string
		RabbitMQPassWord string
		RabbitMQHost     string
		RabbitMQPort     string
	}

	Wechat struct {
		AppID  string
		Secret string
	}

	SMS struct{}

	OSS struct {
		EndPoint        string
		AccessKeyId     string
		AccessKeySecret string
		BucketName      string
		MainDirectory   int
		PartSize        int
	}
}{}
