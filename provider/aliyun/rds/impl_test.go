package rds_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/mcube/logger/zap"
	"github.com/tqtcloud/cloud-manage/apps/rds"
	"github.com/tqtcloud/cloud-manage/provider"
	"github.com/tqtcloud/cloud-manage/provider/aliyun"
)

var (
	operator provider.RdsOperator
)

func TestQuery(t *testing.T) {
	req := provider.NewQueryRequest()
	pager := operator.PageQueryRds(req)

	for pager.Next() {
		set := rds.NewSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		fmt.Println(set)
	}
}

func TestDescribe(t *testing.T) {
	req := provider.NewDescribeRequest("xxx")
	ins, err := operator.DescribeRds(context.TODO(), req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func init() {
	zap.DevelopmentSetup()
	err := aliyun.LoadOperatorFromEnv()
	if err != nil {
		panic(err)
	}
	operator = aliyun.O().RdsOperator()
}
