package logx

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.Logger
func init(){
	file,err :=  os.Create("./_runtime/logs/access.log")
	if err != nil{
		panic(err)
	}
	wsy := zapcore.AddSync(file)
	Logger = zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),wsy,zapcore.InfoLevel| zapcore.ErrorLevel),
			zap.AddCaller(),
	)
	Logger,_ = zap.NewProduction()

	//defer Logger.Sync()
}
