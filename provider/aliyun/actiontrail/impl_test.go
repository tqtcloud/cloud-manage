package actiontrail_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/infraboard/mcube/logger/zap"
	"github.com/tqtcloud/cloud-manage/apps/lb"
	"github.com/tqtcloud/cloud-manage/provider"
	"github.com/tqtcloud/cloud-manage/provider/aliyun"
)

var (
	operator provider.EventOperator
	ctx      = context.Background()
)

func TestPageQueryEvent(t *testing.T) {
	req := provider.NewQueryEventRequest()
	req.StartTime = time.Now().Add(-24 * 10 * time.Hour)
	pager := operator.PageQueryEvent(req)

	for pager.Next() {
		set := lb.NewLoadBalancerSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		fmt.Println(set)
	}
}

func init() {
	zap.DevelopmentSetup()

	err := aliyun.LoadOperatorFromEnv()
	if err != nil {
		panic(err)
	}

	operator = aliyun.O().EventOperator()
}
