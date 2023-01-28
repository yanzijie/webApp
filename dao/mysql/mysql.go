package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/yanzijie/webApp/settings"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init(cfg *settings.MySqlConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	//dsn := "root:123456@tcp(127.0.0.1:3306)/testDB?charset=utf8mb4&parseTime=True"
	//使用connect连接数据库
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect mysql failed", zap.Error(err))
		return err
	}
	db.SetMaxOpenConns(cfg.MaxOpenConn)
	db.SetMaxIdleConns(cfg.MaxIdleConn)
	return
}

func Close() {
	_ = db.Close()
}
