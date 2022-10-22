package impl

import (
	"context"

	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"
	"github.com/tqtcloud/cloud-manage/apps/rds"
)

func (s *service) SyncRDS(ctx context.Context, h *rds.Rds) (
	*rds.Rds, error) {
	h.Resource.Meta.Id = xid.New().String()
	h.Resource.Meta.SyncAt = ftime.Now().Timestamp()

	if err := s.save(ctx, h); err != nil {
		return nil, err
	}
	return h, nil
}
