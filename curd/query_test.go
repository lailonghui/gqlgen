/*
@Time : 2020/10/14 10:04
@Author : Administrator
@Description :
@File : query_test
@Software: GoLand
*/
package curd

import (
	"fmt"
	"testing"
)

var q Query

func init() {
	q.Init()
}

func TestGetEnterpriseList(t *testing.T) {

	obj, err := q.GetEnterpriseList([]string{"id"})
	if err != nil {
		fmt.Printf("GetLabelList() failed,err:%v\n", obj)
		return
	}
	for _, label := range obj {
		fmt.Println(label)
	}
}
