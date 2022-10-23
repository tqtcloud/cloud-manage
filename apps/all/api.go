package all

import (
	_ "github.com/tqtcloud/cloud-manage/apps/bill/api"
	_ "github.com/tqtcloud/cloud-manage/apps/dict/api"
	_ "github.com/tqtcloud/cloud-manage/apps/host/api"
	_ "github.com/tqtcloud/cloud-manage/apps/resource/api"
	_ "github.com/tqtcloud/cloud-manage/apps/secret/api"
	_ "github.com/tqtcloud/cloud-manage/apps/task/api"
	// 注册所有HTTP服务模块, 暴露给框架HTTP服务器加载
	_ "github.com/tqtcloud/cloud-manage/apps/book/api"
)
