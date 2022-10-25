package all

import (
	_ "github.com/tqtcloud/cloud-manage/apps/bill/impl"
	// 注册所有GRPC服务模块, 暴露给框架GRPC服务器加载, 注意 导入有先后顺序
	//_ "github.com/tqtcloud/cloud-manage/apps/book/impl"
	_ "github.com/tqtcloud/cloud-manage/apps/host/impl"
	_ "github.com/tqtcloud/cloud-manage/apps/rds/impl"
	_ "github.com/tqtcloud/cloud-manage/apps/resource/impl"
	_ "github.com/tqtcloud/cloud-manage/apps/secret/impl"
	_ "github.com/tqtcloud/cloud-manage/apps/task/impl"
)
