package db

import (
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var (
	RedisClient *redis.Client
)

type RedisConfig struct {
	MasterName    string   `default:""`
	SentinelAddrs []string `default:{"localhost:26379"}`
	Password      string   `default:""`
	DB            int      `default:0`
}

// InitRedis init redis connection and client
func InitRedis(conf *RedisConfig) {
	//logrus.Info(conf)

	//failOverOpt := redis.FailoverOptions{
	//	MasterName:    conf.MasterName,
	//	SentinelAddrs: conf.SentinelAddrs,
	//	Password:      conf.Password,
	//	DB:            conf.DB,
	//}
	//client := redis.NewFailoverClient(&failOverOpt)
	client := redis.NewClient(&redis.Options{
		Addr:     conf.SentinelAddrs[0],
		Password: conf.Password,
		DB:       conf.DB,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		logrus.Fatal("redis connect demo return ", pong, " ", err.Error())
		panic(err)
	}
	logrus.Info("redis connect success")
	RedisClient = client
}
