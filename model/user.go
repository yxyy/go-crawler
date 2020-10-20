package model

import (
	"crypto/md5"
	"fmt"
	"lhc.go.crawler/libs/mysql"
)

type User struct {
	Id int 	`json:"id"`
	Account string `json:"account" form:"account"`
	NickName string `json:"username" form:"username"`
	PassWord string `json:"password" form:"password"`
	Slat string `json:"slat"`
	Status int `json:"status"`
	GroupId int `json:"group_id"`
	Phone int	`json:"phone"`
	Expiration int64 `json:"expiration"`
	LoginNum int	`json:"login_num"`
	LastLoginTime int64 `json:"last_login_time"`
	CreateTime int64 `json:"create_time"`
	UpdateTime int64 `json:"update_time"`
}

func (u *User) MatchAccount() bool {
	if u.Account=="" {
		return false
	}
	if err := mysql.MysqlConnet.Model(&u).First(&u).Error;err!=nil{
		return false
	}
	return true
}

func (u *User) MatchPassWord() bool {
	if u.Id==0 {
		return false
	}
	if u.PassWord=="" {
		return false
	}
	if u.Slat=="" {
		return false
	}
	password := fmt.Sprint(md5.Sum([]byte(u.PassWord+u.Slat)))
	if err:=mysql.MysqlConnet.Model(&u).Where("id = ?",u.Id).Where("password = ?",password).Find(&u).Error;err!=nil{
		return false
	}
	return true
}
