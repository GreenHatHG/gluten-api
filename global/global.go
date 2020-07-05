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
	MINI  *Mini
)

type MySQL struct {
	Username string
	Password string
	Host     string
}

type Mini struct {
	AppId  string
	Secret string
}

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println(err)
		os.Exit(-2)
	}
	MYSQL = new(MySQL)
	err = cfg.Section("MySQL").MapTo(MYSQL)

	MINI = new(Mini)
	err = cfg.Section("Mini").MapTo(MINI)

	salt, pwd, iter := GetParams()
	MYSQL.Host, _ = util.Decrypt(pwd, iter, MYSQL.Host, []byte(salt))
	MYSQL.Password, _ = util.Decrypt(pwd, iter, MYSQL.Password, []byte(salt))
	MINI.Secret, _ = util.Decrypt(pwd, iter, MINI.Secret, []byte(salt))
	MINI.AppId, _ = util.Decrypt(pwd, iter, MINI.AppId, []byte(salt))

	//s := ""
	//data, _ := util.Encrypt(pwd, iter, s, []byte(salt))
	//fmt.Println("加密后：", data)
}

func GetParams() (string, string, int) {
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
