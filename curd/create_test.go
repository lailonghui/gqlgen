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

	for i := 0; i < 10000000; i++ {
		enterpriseInfo := &model.EnterpriseInfo{
			EnterpriseID:    xid.New().String(),
			EnterpriseCode:  string(i),
			EnterpriseName:  "辉哥股份有限公司",
			EnterpriseLevel: i,
			Remark:          fmt.Sprintf("第%d个备注", i),
		}
		c.CreateEnterpriseInfo(enterpriseInfo)
	}
	//enterpriseInfo := &model.EnterpriseInfo{
	//	EnterpriseName: "小灰灰科技有限公司",
	//	EnterpriseID:   xid.New().String(),
	//}
	//err := c.CreateEnterpriseInfo(enterpriseInfo)
	//if err != nil {
	//	fmt.Printf("CreateEnterpriseInfo failed, err=%v\n", err)
	//}
}
