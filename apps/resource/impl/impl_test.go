package impl_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/mcube/logger/zap"
	"github.com/tqtcloud/cloud-manage/apps/host"
	"github.com/tqtcloud/cloud-manage/provider"

	"github.com/tqtcloud/cloud-manage/provider/aliyun"
)

var (
	operator provider.HostOperator
)

func TestPageQueryHost(t *testing.T) {
	req := provider.NewQueryRequest()
	pager := operator.PageQueryHost(req)
	for pager.Next() {
		set := host.NewHostSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		for i := range set.Items {
			fmt.Println(set.Items[i])
		}
	}
}

func init() {
	zap.DevelopmentSetup()
	err := aliyun.LoadOperatorFromEnv()
	if err != nil {
		panic(err)
	}
	operator = aliyun.O().HostOperator()
}
