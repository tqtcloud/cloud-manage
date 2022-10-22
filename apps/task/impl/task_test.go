package impl_test

import (
	"context"
	"testing"
	"time"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/stretchr/testify/assert"
	"github.com/tqtcloud/cloud-manage/apps/resource"
	"github.com/tqtcloud/cloud-manage/apps/task"
	"github.com/tqtcloud/cloud-manage/conf"

	_ "github.com/tqtcloud/cloud-manage/apps/all"
)

var (
	svc task.ServiceServer
)

func TestSyncHost(t *testing.T) {
	should := assert.New(t)

	req := task.NewCreateTaskRequst()
	req.Type = task.Type_RESOURCE_SYNC
	req.ResourceType = resource.TYPE_HOST
	req.CredentialId = "c5pcffua0bro7e7a05j0"
	req.Region = "ap-shanghai"

	ins, err := svc.CreatTask(context.Background(), req)
	if should.NoError(err) {
		t.Log(ins.Status)
		time.Sleep(10 * time.Second)
	}
}

func TestSyncBill(t *testing.T) {
	should := assert.New(t)

	req := task.NewCreateTaskRequst()
	req.Type = task.Type_RESOURCE_SYNC
	req.ResourceType = resource.TYPE_BILL
	req.CredentialId = "c5pcffua0bro7e7a05j0"
	req.Params["month"] = "2022-04"
	ins, err := svc.CreatTask(context.Background(), req)
	if should.NoError(err) {
		t.Log(ins.Status)
		time.Sleep(10 * time.Second)
	}
}

func TestSyncRds(t *testing.T) {
	should := assert.New(t)

	req := task.NewCreateTaskRequst()
	req.Type = task.Type_RESOURCE_SYNC
	req.ResourceType = resource.TYPE_RDS
	req.CredentialId = "c9rse891eqlk2nabpb10"
	req.Region = "ap-shanghai"

	ins, err := svc.CreatTask(context.Background(), req)
	if should.NoError(err) {
		t.Log(ins.Status)
		time.Sleep(10 * time.Second)
	}
}

func init() {
	zap.DevelopmentSetup()
	if err := conf.LoadConfigFromEnv(); err != nil {
		panic(err)
	}
	if err := app.InitAllApp(); err != nil {
		panic(err)
	}
	svc = app.GetGrpcApp(task.AppName).(task.ServiceServer)
}
