package pkg

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql()*gorm.DB  {
	db, err := gorm.Open(mysql.Open(Cfg.MysqlDsn), &gorm.Config{})
	if err != nil{
		panic(err)
	}
	return db
}
