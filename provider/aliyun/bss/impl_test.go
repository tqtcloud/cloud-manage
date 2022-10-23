package bss_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/infraboard/mcube/logger/zap"
	"github.com/tqtcloud/cloud-manage/apps/bill"
	"github.com/tqtcloud/cloud-manage/apps/order"
	"github.com/tqtcloud/cloud-manage/provider"

	"github.com/tqtcloud/cloud-manage/provider/aliyun"
)

var (
	operator provider.BillOperator
)

func TestQueryBill(t *testing.T) {
	req := provider.NewQueryBillRequest()
	req.Month = "2022-05"

	pager := operator.PageQueryBill(req)
	for pager.Next() {
		set := bill.NewBillSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			t.Logf(err.Error())
			return
		}
		for i := range set.Items {
			fmt.Println(set.Items[i].ToJsonString())
		}
	}
}

func TestQuerySummary(t *testing.T) {
	req := provider.NewQueryBillSummaryRequeset()
	req.Month = "2022-06"
	operator.QuerySummary(context.TODO(), req)
}

func TestQueryOrder(t *testing.T) {
	req := provider.NewQueryOrderRequest()
	req.StartTime = time.Now().Add(-10 * time.Hour)
	pager := operator.PageQueryOrder(req)
	for pager.Next() {
		set := order.NewOrderSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		for i := range set.Items {
			fmt.Println(set.Items[i].ToJsonString())
		}
	}
}

func init() {
	zap.DevelopmentSetup()

	err := aliyun.LoadOperatorFromEnv()
	if err != nil {
		panic(err)
	}
	operator = aliyun.O().BillOperator()
}
