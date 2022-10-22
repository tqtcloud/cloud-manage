package provider

import (
	"context"

	"github.com/infraboard/mcube/pager"
	"github.com/tqtcloud/cloud-manage/apps/lb"
)

type LoadBalancerOperator interface {
	DescribeLoadBalancer(context.Context, *DescribeRequest) (*lb.LoadBalancer, error)
	PageQueryLoadBalancer(*QueryRequest) pager.Pager
}
