package initialize

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gluten/global"
	"os"
)

func Mysql() {
	str := fmt.Sprintf("%s:%s@(%s)/gluten?charset=utf8&parseTime=True&loc=Local",
		global.MYSQL.Username, global.MYSQL.Password, global.MYSQL.Host)
	if db, err := gorm.Open("mysql", str); err != nil {
		fmt.Println("Mysql连接异常")
		fmt.Println(err)
		os.Exit(-1)
	} else {
		global.DB = db
		global.DB.DB().SetMaxIdleConns(10)
		global.DB.DB().SetMaxOpenConns(10)
		global.DB.SingularTable(true)
	}
}
