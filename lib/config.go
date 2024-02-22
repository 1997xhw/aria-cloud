package lib

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var Cfg *ini.File

type ServerConfig struct {
	HTTPPort        int
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	Type            string
	User            string
	Password        string
	Host            string
	DbName          string
	TablePrefix     string
	RedisHost       string
	RedisIndex      string
	UploadLocation  string
	AccessKeyID     string
	AccessKeySecret string
	Endpoint        string
	BucketName      string
}

// LoadServerConfig 加载服务端配置
func LoadServerConfig() ServerConfig {
	var err error
	Cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatal(2, "Fail to parse 'conf/app.ini': %v", err)
	}
	app, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatal(2, "Fail to get section 'app': %v", err)
	}

	//server配置节点读取
	server, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatal(2, "Fail to get section 'server': %v", err)
	}

	//database配置节点读取
	database, err := Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	//redis 配置节点读取
	redis, err := Cfg.GetSection("redis")
	if err != nil {
		log.Fatal(2, "Fail to get section 'redis': %v", err)
	}

	//阿里云oss配置
	oss, err := Cfg.GetSection("oss")
	if err != nil {
		log.Fatal(2, "Fail to get section 'oss': %v", err)
	}

	Config := ServerConfig{
		HTTPPort:        server.Key("HTTP_PORT").MustInt(),
		ReadTimeout:     time.Duration(server.Key("READ_TIMEOUT").MustInt(60)) * time.Second,
		WriteTimeout:    time.Duration(server.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second,
		Type:            database.Key("TYPE").MustString(""),
		User:            database.Key("USER").MustString(""),
		Password:        database.Key("PASSWORD").MustString(""),
		Host:            database.Key("HOST").MustString(""),
		DbName:          database.Key("NAME").MustString(""),
		TablePrefix:     database.Key("TABLE_PREFIX").MustString(""),
		RedisHost:       redis.Key("HOST").MustString(""),
		RedisIndex:      redis.Key("INDEX").MustString(""),
		AccessKeyID:     oss.Key("ACCESS_KEY_ID").MustString(""),
		AccessKeySecret: oss.Key("ACCESS_KEY_SECRET").MustString(""),
		Endpoint:        oss.Key("END_POINT").MustString(""),
		BucketName:      oss.Key("BUCKET_NAME").MustString(""),
		UploadLocation:  app.Key("LOCATION").MustString(""),
	}

	return Config

}
