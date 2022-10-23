package dict

import (
	"github.com/tqtcloud/cloud-manage/apps/resource"
	"github.com/tqtcloud/cloud-manage/apps/secret"
	"github.com/tqtcloud/cloud-manage/utils"

	ali_region "github.com/tqtcloud/cloud-manage/provider/aliyun/region"
	//aws_region "github.com/tqtcloud/cloud-manage/provider/aws/region"
	//hw_region "github.com/tqtcloud/cloud-manage/provider/huawei/region"
	//tx_region "github.com/tqtcloud/cloud-manage/provider/txyun/region"
)

const (
	AppName = "dict"
)

var (
	CrendentialTypes = []utils.EnumDescribe{
		{Value: secret.TYPE_API_KEY.String(), Describe: "API凭证"},
		{Value: secret.TYPE_PASSWORD.String(), Describe: "用户名密码"},
	}

	Vendors = []utils.EnumDescribe{
		{Value: resource.VENDOR_ALIYUN.String(), Describe: "阿里云"},
		{Value: resource.VENDOR_TENCENT.String(), Describe: "腾讯云"},
		{Value: resource.VENDOR_HUAWEI.String(), Describe: "华为云"},
		{Value: resource.VENDOR_AMAZON.String(), Describe: "AWS"},
		{Value: resource.VENDOR_VSPHERE.String(), Describe: "Vsphere"},
	}

	ResourceTypes = map[string][]utils.EnumDescribe{
		resource.VENDOR_ALIYUN.String(): {
			{Name: "主机", Value: resource.TYPE_HOST.String(), Describe: "ECS", Meta: utils.ParamType("region")},
			{Name: "关系型数据库", Value: resource.TYPE_RDS.String(), Describe: "RDS", Meta: utils.ParamType("region")},
			{Name: "月账单", Value: resource.TYPE_BILL.String(), Describe: "月账单", Meta: utils.ParamType("month")},
		},
		resource.VENDOR_TENCENT.String(): {
			{Name: "主机", Value: resource.TYPE_HOST.String(), Describe: "CVM", Meta: utils.ParamType("region")},
			{Name: "关系型数据库", Value: resource.TYPE_RDS.String(), Describe: "CDB", Meta: utils.ParamType("region")},
			{Name: "月账单", Value: resource.TYPE_BILL.String(), Describe: "月账单", Meta: utils.ParamType("month")},
		},
		resource.VENDOR_HUAWEI.String(): {
			{Name: "主机", Value: resource.TYPE_HOST.String(), Describe: "ECS", Meta: utils.ParamType("region")},
			{Name: "关系型数据库", Value: resource.TYPE_RDS.String(), Describe: "RDS", Meta: utils.ParamType("region")},
			{Name: "月账单", Value: resource.TYPE_BILL.String(), Describe: "月账单", Meta: utils.ParamType("month")},
		},
		resource.VENDOR_AMAZON.String(): {
			{Name: "主机", Value: resource.TYPE_HOST.String(), Describe: "EC2", Meta: utils.ParamType("region")},
		},
		resource.VENDOR_VSPHERE.String(): {
			{Name: "主机", Value: resource.TYPE_HOST.String(), Describe: "VM"},
		},
	}

	Regions = map[string][]utils.EnumDescribe{
		resource.VENDOR_ALIYUN.String(): ali_region.Regions,
		//resource.VENDOR_TENCENT.String(): tx_region.Regions,
		//resource.VENDOR_HUAWEI.String():  hw_region.Regions,
		//resource.VENDOR_AMAZON.String():  aws_region.Regions,
		//resource.VENDOR_VSPHERE.String(): hw_region.Regions,
	}
)
