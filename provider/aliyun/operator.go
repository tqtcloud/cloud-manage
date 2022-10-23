package aliyun

import (
	"github.com/caarlos0/env/v6"
	"github.com/tqtcloud/cloud-manage/provider"
	"github.com/tqtcloud/cloud-manage/provider/aliyun/actiontrail"
	"github.com/tqtcloud/cloud-manage/provider/aliyun/bss"
	"github.com/tqtcloud/cloud-manage/provider/aliyun/cms"
	"github.com/tqtcloud/cloud-manage/provider/aliyun/connectivity"
	"github.com/tqtcloud/cloud-manage/provider/aliyun/dds"
	"github.com/tqtcloud/cloud-manage/provider/aliyun/dns"
	"github.com/tqtcloud/cloud-manage/provider/aliyun/ecs"
	"github.com/tqtcloud/cloud-manage/provider/aliyun/oss"
	"github.com/tqtcloud/cloud-manage/provider/aliyun/rds"
	"github.com/tqtcloud/cloud-manage/provider/aliyun/redis"
	"github.com/tqtcloud/cloud-manage/provider/aliyun/slb"
)

var (
	operator *Operator
)

func O() *Operator {
	if operator == nil {
		panic("please load config first")
	}
	return operator
}

func LoadOperatorFromEnv() error {
	conf := &connectivity.AliCloudClient{}
	if err := env.Parse(conf); err != nil {
		return err
	}
	op, err := NewOperator(conf.AccessKey, conf.AccessSecret, conf.Region)
	if err != nil {
		return err
	}
	operator = op

	return nil
}

func NewOperator(ak, sk, region string) (*Operator, error) {
	client := connectivity.NewAliCloudClient(ak, sk, region)

	account, err := client.Account()
	if err != nil {
		return nil, err
	}

	return &Operator{
		client:  client,
		account: account,
	}, nil
}

type Operator struct {
	client  *connectivity.AliCloudClient
	account string
}

func (o *Operator) HostOperator() provider.HostOperator {
	c, err := o.client.EcsClient()
	if err != nil {
		panic(err)
	}
	op := ecs.NewEcsOperator(c)
	op.WithAccountId(o.account)
	return op
}

func (o *Operator) BillOperator() provider.BillOperator {
	c, err := o.client.BssClient()
	if err != nil {
		panic(err)
	}
	return bss.NewBssOperator(c)
}

func (o *Operator) EventOperator() provider.EventOperator {
	c, err := o.client.ActionTrailClient()
	if err != nil {
		panic(err)
	}
	return actiontrail.NewOperator(c)
}

func (o *Operator) CmsOperator() provider.CmsOperator {
	c, err := o.client.CmsClient()
	if err != nil {
		panic(err)
	}
	return cms.NewCmsOperator(c)
}

func (o *Operator) RdsOperator() provider.RdsOperator {
	c, err := o.client.RdsClient()
	if err != nil {
		panic(err)
	}
	return rds.NewRdsOperator(c)
}

func (o *Operator) MongoOperator() provider.MongoOperator {
	c, err := o.client.MongoDBClient()
	if err != nil {
		panic(err)
	}
	return dds.NewOperator(c)
}

func (o *Operator) RedisOperator() provider.RedisOperator {
	c, err := o.client.RedisClient()
	if err != nil {
		panic(err)
	}
	return redis.NewRedisOperator(c)
}

func (o *Operator) OssOperator() provider.OssOperator {
	c, err := o.client.OssClient()
	if err != nil {
		panic(err)
	}
	return oss.NewOssOperator(c)
}

func (o *Operator) LbOperator() provider.LoadBalancerOperator {
	c, err := o.client.SLBClient()
	if err != nil {
		panic(err)
	}
	return slb.NewSLBOperator(c)
}

func (o *Operator) DnsOperator() provider.DnsOperator {
	c, err := o.client.DomainClient()
	if err != nil {
		panic(err)
	}
	d, err := o.client.DnsClient()
	if err != nil {
		panic(err)
	}

	op := dns.NewDomainOperator(c, d)
	return op
}
