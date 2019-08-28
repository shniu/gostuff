package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

// @author shniu
// http://www.flysnow.org/2017/05/06/go-in-action-go-log.html  Good!!!
// @link https://stackoverflow.com/questions/19965795/go-golang-write-log-to-file?rq=1
// @link https://yourbasic.org/golang/log-to-file/

// const (
//	Ldate         = 1 << iota     //日期示例： 2009/01/23
//	Ltime                         //时间示例: 01:23:23
//	Lmicroseconds                 //毫秒示例: 01:23:23.123123.
//	Llongfile                     //绝对路径和行号: /a/b/c/d.go:23
//	Lshortfile                    //文件和行号: d.go:23.
//	LUTC                          //日期时间转为0时区的
//	LstdFlags     = Ldate | Ltime //Go提供的标准抬头信息
//)

//var (
//	Debug   *log.Logger
//	Info    *log.Logger
//	Warning *log.Logger
//	Error   *log.Logger
//)
//func LogInit() {
//	//log.SetPrefix("[kvs] ")
//	//log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
//
//	Debug = log.New(os.Stdout, "[Debug] ", log.Ldate|log.Ltime|log.Lshortfile)
//	Info = log.New(os.Stdout, "[Info] ", log.Ldate|log.Ltime|log.Lshortfile)
//	Warning = log.New(os.Stdout, "[Warning] ", log.Ldate|log.Ltime|log.Lshortfile)
//	Error = log.New(io.MultiWriter(os.Stderr), "[Error] ", log.Ldate|log.Ltime|log.Lshortfile)
//
//	Debug.Println("Debug logger initialized")
//	Info.Println("Info logger initialized")
//	Warning.Println("Warning logger initialized")
//	Error.Println("Error logger initialized")
//}

var (
	Logger = logrus.New()
)

func init() {
	Logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	Logger.SetOutput(os.Stdout)
	Logger.SetLevel(logrus.DebugLevel)
}
