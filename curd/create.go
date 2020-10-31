/*
@Time : 2020/10/14 8:57
@Author : Administrator
@Description :
@File : create
@Software: GoLand
*/
package curd

import (
	"github.com/jinzhu/gorm"
	"lai.com/GraphQL_Server/db"
	"lai.com/GraphQL_Server/graph/model"
)

type Create struct {
	db *gorm.DB
}

func (c *Create) Init() {
	var con db.Connect
	c.db = con.GetConnection()
	//开启debug模式
	//c.db = c.db.Debug()
}

//新增企业信息
func (c *Create) CreateEnterpriseInfo(obj *model.EnterpriseInfo) (err error) {
	err = c.db.Create(obj).Error
	return
}

//新增最新位置信息
func (c *Create) CreateLocationLatest(obj model.VehicleLocationLatest) (err error) {
	err = c.db.Create(&obj).Error
	return
}
