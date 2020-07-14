package util

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
	"time"
)

var Logger *logrus.Logger

func init() {
	Logger = NewLogger("logrus.log")
}

func NewLogger(fileName string) *logrus.Logger {
	if err := CreatePath("log"); err != nil {
		log.Fatal("创建文件夹失败", err)
	}

	//写入文件
	src, err := os.OpenFile(path.Join("log", fileName), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
	}

	//实例化
	logger := logrus.New()
	//设置输出
	logger.Out = src
	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	// 设置日志分割
	logWriter, errRota := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",
		// 生成软链，指向最新日志文件
		//rotatelogs.WithLinkName(fileName),
		// 设置最大保存时间(30天)
		rotatelogs.WithMaxAge(30*24*time.Hour),
		// 设置日志切割时间间隔(2天)
		rotatelogs.WithRotationTime(2*24*time.Hour),
	)
	if errRota != nil {
		log.Fatalln("设置日志分割失败：", errRota)
	}
	log.SetOutput(logWriter)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return logger
}

func CreatePath(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(path, 0777)
			return err
		}
		err = os.Mkdir(path, os.ModeAppend)
		return err
	}
	return nil
}
