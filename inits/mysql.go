package inits

import (
	"goWebDemo/dao"

	"goWebDemo/conf"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/pmars/beego/logs"
)

func initMysqlEngine() {
	logs.Debug("Init Mysql Engine Now...")
	// 初始化Mysql
	dao.Engine = initOneMysqlEngine(
		conf.Config.Mysql.Conn,
		conf.Config.Mysql.MaxActive,
		conf.Config.Mysql.MaxIdle)
	logs.Debug("Init Mysql Engine Done!!!")
}

func initOneMysqlEngine(conn string, maxActive, maxIdle int) *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", conn)
	if err != nil {
		panic(err)
	}
	engine.ShowSQL(true)
	engine.SetMaxOpenConns(maxActive)
	engine.SetMaxIdleConns(maxIdle)
	if err = engine.Ping(); err != nil {
		logs.Info("mysql error:%v", err)
		panic(err)
	}

	return engine
}
