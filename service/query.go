/*
@Time : 2020/10/14 11:16
@Author : Administrator
@Description :
@File : query
@Software: GoLand
*/
package service

import (
	"fmt"
	"lai.com/GraphQL_Server/curd"
	"lai.com/GraphQL_Server/model"
)

var query curd.Query

func init() {
	query.Init()
}

//获取所有的企业信息列表
func GetEnterpriseList(fieldNames []string) (enterpriseList []*model.EnterpriseInfo, err error) {
	enterpriseList, err = query.GetEnterpriseList(fieldNames)
	if err != nil {
		fmt.Printf("create.CreateEnterpriseInfo() failed,err:%v\n", err)
		return
	}

	return enterpriseList, err
}
