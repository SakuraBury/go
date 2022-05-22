package config

import (
	"database/sql"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

var myDbCollection *sql.DB

func init() {
	host, _ := beego.AppConfig.String("db::host")
	port, _ := beego.AppConfig.String("db::port")
	database, _ := beego.AppConfig.String("db::database")
	user, _ := beego.AppConfig.String("db::user")
	password, _ := beego.AppConfig.String("db::password")

	dbcon := fmt.Sprintf("%s:%s@tcp(%s%s)/%s", user, password, host, port, database)

	fmt.Println(dbcon)
	db, err := sql.Open("mysql", dbcon)
	if err != nil {
		fmt.Printf("open mysql %s err\n", dbcon)
		panic(err)
	}
	myDbCollection = db
}

func GetMySqlConnection() *sql.DB {
	return myDbCollection
}
