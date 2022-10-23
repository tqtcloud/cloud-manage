package dds_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/infraboard/mcube/logger/zap"
	"github.com/tqtcloud/cloud-manage/apps/mongodb"
	"github.com/tqtcloud/cloud-manage/provider"
	"github.com/tqtcloud/cloud-manage/provider/aliyun"
)

var (
	operator provider.MongoOperator
)

func TestQuery(t *testing.T) {
	req := provider.NewQueryRequest()
	pager := operator.PageQueryMongo(req)

	for pager.Next() {
		set := mongodb.NewMongoDBSet()
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
	operator = aliyun.O().MongoOperator()
}
