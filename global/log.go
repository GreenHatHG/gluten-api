package global

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
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

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func GinLoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		NewLogger("gin-request.log").WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
			"response":     blw.body.String(),
		}).Info()
	}
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
