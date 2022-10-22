package impl

import (
	"database/sql"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/tqtcloud/cloud-manage/apps/bill"
	"github.com/tqtcloud/cloud-manage/apps/host"
	"github.com/tqtcloud/cloud-manage/apps/rds"
	"github.com/tqtcloud/cloud-manage/apps/secret"
	"github.com/tqtcloud/cloud-manage/apps/task"
	"github.com/tqtcloud/cloud-manage/conf"
	"google.golang.org/grpc"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	db  *sql.DB
	log logger.Logger
	task.UnimplementedServiceServer

	secret secret.ServiceServer
	host   host.ServiceServer
	bill   bill.ServiceServer
	rds    rds.ServiceServer
}

func (s *service) Config() error {
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}

	s.log = zap.L().Named(s.Name())
	s.db = db

	s.secret = app.GetGrpcApp(secret.AppName).(secret.ServiceServer)
	s.host = app.GetGrpcApp(host.AppName).(host.ServiceServer)
	s.bill = app.GetGrpcApp(bill.AppName).(bill.ServiceServer)
	s.rds = app.GetGrpcApp(rds.AppName).(rds.ServiceServer)
	return nil
}

func (s *service) Name() string {
	return task.AppName
}

func (s *service) Registry(server *grpc.Server) {
	task.RegisterServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
