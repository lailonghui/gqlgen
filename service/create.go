/*
@Time : 2020/10/14 8:54
@Author : Administrator
@Description :
@File : create
@Software: GoLand
*/
package service

import (
	"fmt"
	"lai.com/GraphQL_Server/curd"
	"lai.com/GraphQL_Server/graph/model"
)

var create curd.Create

func init() {
	create.Init()
}

//添加企业信息的service
func AddEnterpriseInfo(obj *model.EnterpriseInfo) (err error) {
	err = create.CreateEnterpriseInfo(obj)
	if err != nil {
		fmt.Printf("create.CreateEnterpriseInfo() failed,err:%v\n", err)
		return
	}
	return
}
