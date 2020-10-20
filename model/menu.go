package model

import (
	"time"
	"errors"
	"lhc.go.crawler/libs/mysql"
)

type Menu struct {
	Id          int    `json:"id"  form:"id"`
	Mark        string `json:"mark"  form:"mark"`
	Type        string `json:"type"  form:"type"`
	Name        string `json:"name"  form:"name"`
	Url         string `json:"url"  form:"url"`
	Parent      int    `json:"parent" gorm:"default:0"  form:"parent"`
	Icon        string `json:"icon"  form:"icon"`
	Weight      int    `json:"weight"  form:"weight"`
	Status      int    `json:"status"  form:"status"`
	Author      string `json:"author"  `
	CreateTime int64  `json:"create_time"  `
	UpdateTime int64  `json:"update_time"  `
}

type MenuTree struct {
	Id       int         `json:"id"`
	Mark     string      `json:"mark"`
	Type     string      `json:"type"`
	Name     string      `json:"name"`
	Url      string      `json:"url"`
	Parent   int         `json:"parent" gorm:"default:0"`
	Icon     string      `json:"icon"`
	Weight   int         `json:"weight"`
	Status   int         `json:"status"`
	Children []*MenuTree
}

func GetParentMenu(parent int) (menu []*Menu)  {
	mysql.MysqlConnet.Table("yxyy_menu").Where("parent = ?",parent).Order("weight").Order("create_time").Find(&menu)
	return
}

func GetMeunTree(parent int) (menuList []*MenuTree) {
	var menu []Menu
	mysql.MysqlConnet.Table("yxyy_menu").Where("parent = ?",parent).Order("weight").Order("create_time").Find(&menu)
	for _,v := range menu {
		Children := GetMeunTree(v.Id) // 拿到每个父菜单的子菜单
		node := &MenuTree{
			Id:     v.Id,
			Name:   v.Name,
			Url:    v.Url,
			Icon:   v.Icon,
			Mark:   v.Mark,
			Weight: v.Weight,
			Parent: v.Parent,
		}
		node.Children = Children
		menuList = append(menuList,node)
	}
	return
}

func (m *Menu) UpdateData () error {
	if m.Name=="" {
		return errors.New("名称不能为空")
	}
	if m.Url=="" {
		return errors.New("url不能为空")
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