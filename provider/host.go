package provider

import (
	"context"

	"github.com/infraboard/mcube/pager"
	"github.com/tqtcloud/cloud-manage/apps/disk"
	"github.com/tqtcloud/cloud-manage/apps/host"
)

type HostOperator interface {
	PageQueryHost(req *QueryRequest) pager.Pager
	PageQueryDisk(req *QueryRequest) pager.Pager
	PageQueryEip(req *QueryRequest) pager.Pager
	DescribeHost(ctx context.Context, req *DescribeRequest) (*host.Host, error)
	DescribeDisk(ctx context.Context, req *DescribeRequest) (*disk.Disk, error)
}
