package pkg

type config struct {
	MysqlDsn string
	Redis struct{
		Addr string
		PassWord string
		DB int
	}
}
var Cfg = &config{
	MysqlDsn: "root:root@tcp(127.0.0.1:3306)/dev_test?charset=utf8mb4&parseTime=True&loc=Local",
	Redis: struct {
		Addr     string
		PassWord string
		DB       int
	}{Addr: "127.0.0.1:6379", PassWord: "", DB: 0},
}
