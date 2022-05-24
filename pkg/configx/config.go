package configx

import (
	"github.com/spf13/viper"
	"os"
	"sync"
)

type config struct {
	MysqlDsn string
	MysqlPre string
	Redis struct{
		Addr string
		PassWord string
		DB int
	}
	App struct{
		Port string
	}
	Log struct{
		FileName string
		MaxSize int
		MaxBackup int
		MaxAge int
	}
}
var once sync.Once
func getConf()*config{
	cfg := config{}
	once.Do(func() {

		path,err := os.Getwd()
		if err != nil{
			panic(err)
		}
		vp := viper.New()
		vp.AddConfigPath(path + "/config")
		vp.SetConfigName("config")
		vp.SetConfigType("yaml")
		if err = vp.ReadInConfig();err!=nil{
			panic(err)
		}


		cfg.MysqlDsn = vp.GetString("database.mysql.dsn")
		cfg.MysqlPre = vp.GetString("database.mysql.tablePrefix")
		cfg.Redis.Addr = vp.GetString("database.redis.addr")
		cfg.Redis.PassWord = vp.GetString("database.redis.password")
		cfg.Redis.DB = vp.GetInt("database.redis.db")

		cfg.App.Port = vp.GetString("app.port")

		cfg.Log.FileName = vp.GetString("log.fileName")
		cfg.Log.MaxAge = vp.GetInt("log.maxAge")
		cfg.Log.MaxBackup = vp.GetInt("log.maxBackup")
		cfg.Log.MaxSize = vp.GetInt("log.maxSize")

	})
	return &cfg
}
var Cfg = getConf()

