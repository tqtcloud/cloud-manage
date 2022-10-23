package dns_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/infraboard/mcube/logger/zap"
	"github.com/tqtcloud/cloud-manage/apps/dns"
	"github.com/tqtcloud/cloud-manage/provider"

	"github.com/tqtcloud/cloud-manage/provider/aliyun"
)

var (
	operator provider.DnsOperator
)

func TestQueryDomain(t *testing.T) {
	req := provider.NewQueryDomainRequest()
	pager := operator.QueryDomain(req)
	for pager.Next() {
		set := dns.NewDomainSet()
		if err := pager.Scan(context.Background(), set); err != nil {
			panic(err)
		}
		fmt.Println(set)
	}
}

func TestQueryRecord(t *testing.T) {
	req := provider.NewQueryRecordRequest(os.Getenv("AL_DNS_DOMAIN"))
	pager := operator.QueryRecord(req)
	for pager.Next() {
		set := dns.NewRecordSet()
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
	operator = aliyun.O().DnsOperator()
}
