/*
@Time : 2020/10/13 15:45
@Author : Administrator
@Description :
@File : init_service
@Software: GoLand
*/
package service

import (
	"lai.com/GraphQL_Server/db"
	"lai.com/GraphQL_Server/model"
)

//初始化数据表
func InitTables() {
	var con db.Connect
	pg := con.GetConnection()
	if !pg.HasTable(&model.EnterpriseInfo{}) {
		pg.CreateTable(&model.EnterpriseInfo{})
	}

	if !pg.HasTable(&model.VehicleLocationLatest{}) {
		pg.CreateTable(&model.VehicleLocationLatest{})
	}

	pg.AutoMigrate(&model.EnterpriseInfo{}, &model.VehicleLocationLatest{})
}
