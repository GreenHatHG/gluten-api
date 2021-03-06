package global

import (
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/ini.v1"
	"os"
)

var (
	DB        *gorm.DB
	MongoDB   *mongo.Database
	MYSQL     *DataBase
	MONGO     *DataBase
	GITHUB    *Github
	JwtConfig *JWT
)

type DataBase struct {
	Username string
	Password string
	Host     string
}

type Github struct {
	ClientID     string
	ClientSecret string
}

type JWT struct {
	Secret string
	Exp    int
}

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println(err)
		os.Exit(-2)
	}

	MYSQL = new(DataBase)
	_ = cfg.Section("MySQL").MapTo(MYSQL)
	MONGO = new(DataBase)
	_ = cfg.Section("Mongo").MapTo(MONGO)

	GITHUB = new(Github)
	_ = cfg.Section("Github").MapTo(GITHUB)

	JwtConfig = new(JWT)
	_ = cfg.Section("JWT").MapTo(JwtConfig)

	salt, pwd, iter := GetParams()
	MYSQL.Host, _ = Decrypt(pwd, iter, MYSQL.Host, []byte(salt))
	MYSQL.Password, _ = Decrypt(pwd, iter, MYSQL.Password, []byte(salt))
	MONGO.Host, _ = Decrypt(pwd, iter, MONGO.Host, []byte(salt))
	MONGO.Password, _ = Decrypt(pwd, iter, MONGO.Password, []byte(salt))
	GITHUB.ClientID, _ = Decrypt(pwd, iter, GITHUB.ClientID, []byte(salt))
	GITHUB.ClientSecret, _ = Decrypt(pwd, iter, GITHUB.ClientSecret, []byte(salt))
	JwtConfig.Secret, _ = Decrypt(pwd, iter, JwtConfig.Secret, []byte(salt))

	//s := ""
	//data, _ := Encrypt(pwd, iter, s, []byte(salt))
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
