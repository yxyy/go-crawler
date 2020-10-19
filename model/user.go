package model

type User struct {
	UserName string `json:"username" form:"username"`
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
