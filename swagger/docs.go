package swagger

import (
	"github.com/go-openapi/spec"
	"github.com/tqtcloud/cloud-manage/version"
)

func Docs(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "Cloud Management",
			Description: "云管平台",
			Contact: &spec.ContactInfo{
				ContactInfoProps: spec.ContactInfoProps{
					Name:  "john",
					Email: "john@doe.rp",
					URL:   "http://johndoe.org",
				},
			},
			License: &spec.License{
				LicenseProps: spec.LicenseProps{
					Name: "MIT",
					URL:  "http://mit.org",
				},
			},
			Version: version.Short(),
		},
	}
	swo.Tags = []spec.Tag{
		//{TagProps: spec.TagProps{
		//	Name:        "books",
		//	Description: "Managing books"},
		//},
		{TagProps: spec.TagProps{
			Name:        "dict",
			Description: "数据字典"},
		},
		{TagProps: spec.TagProps{
			Name:        "host",
			Description: "云主机"},
		},
		{TagProps: spec.TagProps{
			Name:        "resource",
			Description: "资源"},
		},
		{TagProps: spec.TagProps{
			Name:        "secret",
			Description: "秘钥管理"},
		},
		{TagProps: spec.TagProps{
			Name:        "task",
			Description: "任务管理"},
		},
	}
}
