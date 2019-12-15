package config

import (
	"github.com/jessevdk/go-flags"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"local/gin_init/model/db"
)

var ServerConfig = struct {
	Mode   string `default:"release"`
	Port   int32  `default:"62004"`
	Logger struct {
		LogLevel    string `default:"error"`
		LogDir      string `default:"/opt/log"`
		LogFileName string `default:"demo.log"`
		LogFormat   string `default:"text"` // text or json
	}
	Mongo db.MongoConfig
	Redis db.RedisConfig
	Mysql db.MysqlConfig
}{}

var opts struct {
	ConfigFile string `short:"c" required:"true" name:"config file"`
}

func init() {
	// init configs
	if _, err := flags.Parse(&opts); err != nil {
		panic(err)
	}
	loadConfigFile(opts.ConfigFile)
}

func loadConfigFile(filePath string) {
	logrus.Debug("load config %s", filePath)
	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatal("fail to read config", err)
		panic(err)
	}
	if err := viper.Unmarshal(&ServerConfig); err != nil {
		logrus.Fatal("fail to unmarshal config", err)
		panic(err)
	}
}
