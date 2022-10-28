package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/pager"
	"github.com/tqtcloud/cloud-manage/apps/bill"
	"github.com/tqtcloud/cloud-manage/apps/resource"
	"github.com/tqtcloud/cloud-manage/apps/secret"
	"github.com/tqtcloud/cloud-manage/apps/task"
	"github.com/tqtcloud/cloud-manage/provider"

	"github.com/tqtcloud/cloud-manage/provider/aliyun"
	//"github.com/tqtcloud/cloud-manage/provider/huawei"
	//"github.com/tqtcloud/cloud-manage/provider/txyun"
)

func (s *service) syncBill(ctx context.Context, secretIns *secret.Secret, t *task.Task, cb SyncTaskCallback) {
	var (
		pager pager.Pager
	)

	// 处理任务状态
	t.Run()
	defer s.syncBillDown(ctx, t, cb)

	secret := secretIns.Data
	req := provider.NewQueryBillRequestWithRate(secret.RequestRate)
	req.Month = t.Data.Params["month"]

	switch secret.Vendor {
	case resource.VENDOR_ALIYUN:
		s.log.Debugf("sync aliyun bill ...")
		op, err := aliyun.NewOperator(secret.ApiKey, secret.ApiSecret, t.Data.Region)
		if err != nil {
			s.log.Errorf("aliyun.NewOperator： %s",err.Error())
			t.Failed(err.Error())
			return
		}
		pager = op.BillOperator().PageQueryBill(req)
	//case resource.VENDOR_TENCENT:
	//	s.log.Debugf("sync txyun bill ...")
	//	op, err := txyun.NewOperator(secret.ApiKey, secret.ApiSecret, t.Data.Region)
	//	if err != nil {
	//		t.Failed(err.Error())
	//		return
	//	}
	//	pager = op.BillOperator().PageQueryBill(req)
	//case resource.VENDOR_HUAWEI:
	//	s.log.Debugf("sync hwyun bill ...")
	//	op, err := huawei.NewOperator(secret.ApiKey, secret.ApiSecret, t.Data.Region)
	//	if err != nil {
	//		t.Failed(err.Error())
	//		return
	//	}
	//	pager = op.BillOperator().PageQueryBill(req)
	default:
		t.Failed(fmt.Sprintf("unsuport bill syncing vendor %s", secret.Vendor))
		return
	}

	// 分页查询数据
	if pager != nil {
		for pager.Next() {
			set := bill.NewBillSet()
			if err := pager.Scan(ctx, set); err != nil {
				t.Failed(fmt.Sprintf("sync error, %s", err))
				return
			}
			for i := range set.Items {
				ins := set.Items[i]
				ins.TaskId = t.Id
				// 打印所有的实例信息
				//s.log.Info(ins)

				// 账单时间没有的，采用请求年月
				if ins.Month == "" {
					ins.Month =  t.Data.Params["month"]
				}
				s.doSyncBill(ctx, ins, t)
			}
		}
	}

}

// doSyncBill 月底账单数据入库
func (s *service) doSyncBill(ctx context.Context, ins *bill.Bill, t *task.Task) {
	h, err := s.bill.SyncBill(ctx, ins)
	//s.log.Debug("doSyncBill ",ins.Year,ins.Month)

	var detail *task.Record
	if err != nil {
		s.log.Warnf("save bill error, %s", err)
		detail = task.NewSyncFailedRecord(t.Id, ins.InstanceId, ins.InstanceName, err.Error())
	} else {
		s.log.Debugf("save bill %s to db", h.ShortDesc())
		detail = task.NewSyncSucceedRecord(t.Id, ins.InstanceId, ins.InstanceName)
	}

	t.AddDetail(detail)
	if err := s.insertTaskDetail(ctx, detail); err != nil {
		s.log.Errorf("update detail error, %s", err)
	}
}

// syncBillDown 任务同步完后的回调函数
func (s *service) syncBillDown(ctx context.Context, t *task.Task, cb SyncTaskCallback) {
	t.Completed()
	cb(t)

	s.log.Debugf("task status: %s", t.Status)
	// 调用bill服务保存数据, 由于账单对象没有更新逻辑
	// 任务同步成功, 确认当前同步版本为正确版本, 删除之前的成本
	// 任务同步失败, 删除当前同步的版本
	if t.Status.Stage.Equal(task.Stage_SUCCESS) {
		resp, err := s.bill.ConfirmBill(ctx, bill.NewConfirmBillRequest(t.Id))
		if err != nil {
			s.log.Errorf("confirm bill error, %s", err)
		} else {
			s.log.Debugf("confirm bill success, total: %d bill", resp.GetTotal())
		}
	} else {
		resp, err := s.bill.DeleteBill(ctx, bill.NewDeleteBillRequest(t.Id))
		if err != nil {
			s.log.Errorf("delete bill error, %s", err)
		} else {
			s.log.Debugf("delete bill success, total: %d bill", resp.Total)
		}
	}
}
