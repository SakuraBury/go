package dao

import (
	"open-center/dal/config"
	"open-center/dal/model"
)

func addUser(do *model.TUserDO) {

	sql := "INSERT INTO user (id,name,password extended_info) VALUES (%s,%s,%s,%s);"

	_, err := config.GetMySqlConnection().Exec(sql, do.Id, do.Name, do.Password, do.ExtendedInfo)
	if err != nil {
		panic(err)
	}
}
