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
	"lai.com/GraphQL_Server/model"
)

type Create struct {
	db *gorm.DB
}

func (c *Create) Init() {
	var con db.Connect
	c.db = con.GetConnection()
	//开启debug模式
	c.db = c.db.Debug()
}

//新增企业信息
func (c *Create) CreateEnterpriseInfo(obj *model.EnterpriseInfo) (err error) {
	err = c.db.Create(obj).Error
	return
}
