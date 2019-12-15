package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	SqlDb *sqlx.DB
)

type MysqlConfig struct {
	UrlDisposition string
}

func InitMysql(conf MysqlConfig) error {
	//logrus.Info(conf.UrlDisposition)
	db, err := sqlx.Connect("mysql", conf.UrlDisposition)
	if err != nil {
		logrus.Error("mysql fail connect to ", conf.UrlDisposition, err.Error())
		//panic(err)
		return err
	}
	logrus.Info("mysql connect success")
	db.SetConnMaxLifetime(time.Second * 60)
	SqlDb = db
	return nil
}
