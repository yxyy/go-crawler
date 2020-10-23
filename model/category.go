package model

import (
	"errors"
	"lhc.go.crawler/libs/mysql"
	"time"
)

type Category struct {
	Id int `json:"id" form:"id"`
	Pid int `json:"pid" form:"pid"`
	Title string `json:"title" form:"title"`
	Mark string `json:"mark" form:"mark"`
	Description string `json:"description" form:"description"`
	Sort string `json:"sort" form:"sort"`
	//Level string `json:"id" form:"level"`
	Author   string `json:"author"`
	CreateTime int64  `json:"create_time"  `
	UpdateTime int64  `json:"update_time"  `
}

func NewCategory() *Category {
	return &Category{}
}

func GetParentCategory(parent int) (category []*Category)  {
	mysql.MysqlConnet.Table("yxyy_category").Where("pid = ?",parent).Order("sort").Order("create_time").Find(&category)
	return
}

func (this *Category) GetList()(data []*Category,total int,err error)  {
	Db :=mysql.MysqlConnet.Model(&this)
	if this.Pid!=0 {
		Db = Db.Where("pid = ?",this.Pid)
	}
	if this.Title !="" {
		Db = Db.Where("title like  ?","%"+this.Title+"%")
	}
	if err:= Db.Find(&data).Error;err!=nil{
		return nil,0,err
	}
	Db.Count(&total)
	return
}

func (this *Category) GetOneCategory() error {
	if this.Id ==0 {
		return errors.New("id 不能为空")
	}
	if err:= mysql.MysqlConnet.Model(&this).First(&this).Error;err!=nil{
			return err
	}
	return nil
}


func (m *Category) UpdateData () error {
	if m.Title=="" {
		return errors.New("名称不能为空")
	}
	m.UpdateTime=time.Now().Unix()
	if m.Id !=0 {
		if err:=mysql.MysqlConnet.Model(&m).Where("id = ?",m.Id).Update(&m).Error;err!=nil{
			return err
		}
	}else {
		m.CreateTime=time.Now().Unix()
		if err:=mysql.MysqlConnet.Model(&m).Create(&m).Error;err!=nil{
			return err
		}
	}
	return nil
}