package provider

import (
	"context"

	"github.com/infraboard/mcube/pager"
	"github.com/tqtcloud/cloud-manage/apps/rds"
)

type RdsOperator interface {
	PageQueryRds(req *QueryRequest) pager.Pager
	DescribeRds(ctx context.Context, req *DescribeRequest) (*rds.Rds, error)
}
