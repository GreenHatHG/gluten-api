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
	DB     *gorm.DB
	MYSQL  *MySQL
	GITHUB *Github
)

type MySQL struct {
	Username string
	Password string
	Host     string
}

type Github struct {
	ClientID     string
	ClientSecret string
}

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println(err)
		os.Exit(-2)
	}
	MYSQL = new(MySQL)
	err = cfg.Section("MySQL").MapTo(MYSQL)

	GITHUB = new(Github)
	err = cfg.Section("Github").MapTo(GITHUB)

	salt, pwd, iter := GetParams()
	MYSQL.Host, _ = util.Decrypt(pwd, iter, MYSQL.Host, []byte(salt))
	MYSQL.Password, _ = util.Decrypt(pwd, iter, MYSQL.Password, []byte(salt))
	GITHUB.ClientID, _ = util.Decrypt(pwd, iter, GITHUB.ClientID, []byte(salt))
	GITHUB.ClientSecret, _ = util.Decrypt(pwd, iter, GITHUB.ClientSecret, []byte(salt))

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
