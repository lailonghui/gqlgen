/*
@Time : 2020/10/14 8:59
@Author : Administrator
@Description :
@File : create_test
@Software: GoLand
*/
package curd

import (
	"fmt"
	"github.com/rs/xid"
	"lai.com/GraphQL_Server/model"
	"testing"
)

var c Create

func init() {
	c.Init()
}

func TestCreateEnterpriseInfo(t *testing.T) {
	enterpriseInfo := &model.EnterpriseInfo{
		EnterpriseName: "小灰灰科技有限公司",
		EnterpriseID:   xid.New().String(),
	}
	err := c.CreateEnterpriseInfo(enterpriseInfo)
	if err != nil {
		fmt.Printf("CreateEnterpriseInfo failed, err=%v\n", err)
	}
}
