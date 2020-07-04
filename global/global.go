package global

import (
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	"gluten/util"
	"gopkg.in/ini.v1"
	"os"
)

var (
	DB    *gorm.DB
	MYSQL *MySQL
)

type MySQL struct {
	Username string
	Password string
	Host     string
}

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println(err)
		os.Exit(-2)
	}
	MYSQL = new(MySQL)
	err = cfg.Section("MySQL").MapTo(MYSQL)

	salt, pwd, iter := getParams()
	MYSQL.Host, _ = util.Decrypt(pwd, iter, MYSQL.Host, []byte(salt))
	MYSQL.Password, _ = util.Decrypt(pwd, iter, MYSQL.Password, []byte(salt))
}

func getParams() (string, string, int) {
	salt := flag.String("salt", "", "")
	password := flag.String("pwd", "", "")
	iter := flag.Int("iter", 0, "")
	flag.Parse()
	if *salt == "" || *password == "" || *iter == 0 {
		fmt.Println("获取加密参数失败")
		os.Exit(-1)
	}
	return *salt, *password, *iter
}
