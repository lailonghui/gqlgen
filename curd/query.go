/*
@Time : 2020/10/14 9:00
@Author : Administrator
@Description :
@File : query
@Software: GoLand
*/
package curd

import (
	"github.com/jinzhu/gorm"
	"lai.com/GraphQL_Server/db"
	"lai.com/GraphQL_Server/model"
)

type Query struct {
	db *gorm.DB
}

func (q *Query) Init() {
	var con db.Connect
	q.db = con.GetConnection()
	//开启debug模式
	q.db = q.db.Debug()
}

func (q *Query) GetEnterpriseList(fieldNames []string) (obj []*model.EnterpriseInfo, err error) {
	err = q.db.Select(fieldNames).Find(&obj).Error
	return
}
