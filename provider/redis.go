package provider

import (
	"context"

	"github.com/infraboard/mcube/pager"
	"github.com/tqtcloud/cloud-manage/apps/redis"
)

type RedisOperator interface {
	PageQueryRedis(req *QueryRequest) pager.Pager
	DescribeRedis(context.Context, *DescribeRequest) (*redis.Redis, error)
}
