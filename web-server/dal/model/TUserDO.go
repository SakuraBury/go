package model

import "time"

type TUserDO struct {
	Id           string
	Name         string
	Password     string
	ExtendedInfo string
	CreateTime   time.Time
}
