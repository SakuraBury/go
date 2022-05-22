package model

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"gorm.io/gorm"
	"user/dal"
)

type TUserDTO struct {
	gorm.Model
	Name     string
	Desc     string
	Phone    string
	Email    string
	Password string
}

// CreateUser /*
func (u *TUserDTO) CreateUser() error {
	result := dal.GetMySqlConnection().Create(u)

	if result.Error != nil || result.RowsAffected <= 0 {
		logs.Error("TUserDTO create err. result:%v", result)
		fmt.Errorf("insert user err")
	}
	return nil
}

// GetUserById /*
func (u *TUserDTO) GetUserById(id uint) {
	dal.GetMySqlConnection().First(u, id)
}

// GetUserByPhone /*
func (u *TUserDTO) GetUserByPhone(phone string) {
	dal.GetMySqlConnection().First(u, "Phone = ?", phone)
}

func (u *TUserDTO) GetUserByPhoneAndPassword(phone string, password string) {
	dal.GetMySqlConnection().First(u, "Phone = ? AND Password = ?", phone, password)
}

func (u *TUserDTO) GetUserByName(name string) []TUserDTO {
	var record []TUserDTO
	dal.GetMySqlConnection().First(&record, "Name LIKE ?", "%"+name+"%")
	return record
}
