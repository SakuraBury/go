package dal

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"gorm.io/gorm"
)
import "gorm.io/driver/mysql"

var myDb *gorm.DB

func init() {
	host, _ := beego.AppConfig.String("db::host")
	port, _ := beego.AppConfig.String("db::port")
	database, _ := beego.AppConfig.String("db::database")
	user, _ := beego.AppConfig.String("db::user")
	password, _ := beego.AppConfig.String("db::password")

	dbcon := fmt.Sprintf("%s:%s@tcp(%s%s)/%s", user, password, host, port, database)
	logs.Debug("===mysql collection gorm: %s", dbcon)
	db, err := gorm.Open(mysql.Open(dbcon))
	if err != nil {
		panic(err)
	}
	myDb = db
}

func GetMySqlConnection() *gorm.DB {
	return myDb
}
