package db

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

const (
	MgoDBDisposition string = "Disposition"
)

var (
	mongoSession map[string]*mgo.Session
	mongo        map[string]*mgo.DialInfo
)

type MongoConfig struct {
	UrlDisposition string
	//UrlXcity stringutil
}

type MgoDB struct {
	*mgo.Database
}

func InitMongo(mongoCfg *MongoConfig) {
	mongoSession = make(map[string]*mgo.Session)
	mongo = make(map[string]*mgo.DialInfo)

	//mode := mgo.Mode(conf.Configure.MongodbMode)
	if mongoCfg.UrlDisposition != "" {
		var err error
		url := mongoCfg.UrlDisposition
		mongo[MgoDBDisposition], err = mgo.ParseURL(url)
		if err != nil {
			panic(err)
		}
		mongoSession[MgoDBDisposition], err = mgo.Dial(url)
		if err != nil {
			logrus.Fatal("mongo fail connect to ", url, err.Error())
			panic(err.Error())
		}
		//mongoSession.SetMode(mode, false)
		mongoSession[MgoDBDisposition].SetSafe(&mgo.Safe{})

		//logrus.Info("mongo connected to ", url)
	}

	//if mongoCfg.UrlXcity != "" {
	//	var err errors
	//	url := mongoCfg.UrlXcity
	//	mongo[MgoDBXcity], err = mgo.ParseURL(url)
	//	if err != nil {
	//		panic(err)
	//	}
	//	mongoSession[MgoDBXcity], err = mgo.Dial(url)
	//	if err != nil {
	//		logrus.Fatal("mongo fail connect to ", url, err.Error())
	//		panic(err.Error())
	//	}
	//	//mongoSession.SetMode(mode, false)
	//	mongoSession[MgoDBXcity].SetSafe(&mgo.Safe{})
	//
	//	logrus.Info("mongo connected to ", url)
	//}

	logrus.Info("mongo connect success")
}

func (m *MgoDB) Close() {
	m.Session.Close()
}

// MongoDB 需要在调用之后，手动关闭Session.Close()
// 为什么使用Clone https://www.mongodb.com/blog/post/running-mongodb-queries-concurrently-with-go
func MongoDB(db string) *MgoDB {
	if db == "" {
		db = MgoDBDisposition
	}
	clone := mongoSession[db].Clone()

	internal := clone.DB(mongo[db].Database)
	return &MgoDB{internal}
}
