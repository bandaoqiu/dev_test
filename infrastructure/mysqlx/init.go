package mysqlx

import (
	"dev_test/pkg/configx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitMysql()*gorm.DB  {
	config := &gorm.Config{
		SkipDefaultTransaction:                   false,
		NamingStrategy:                           schema.NamingStrategy{
			TablePrefix:   configx.Cfg.MysqlPre,
			SingularTable: true,
		},
		Logger:                                   logger.Default.LogMode(logger.Info),
		PrepareStmt:                              false,
	}
	db, err := gorm.Open(mysql.Open(configx.Cfg.MysqlDsn), config)
	if err != nil{
		panic(err)
	}
	sqldb,_ := db.DB()

	sqldb.SetMaxIdleConns(100)
	sqldb.SetMaxOpenConns(80)
	return db
}
