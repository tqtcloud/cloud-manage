package impl

import (
	"github.com/tqtcloud/cloud-manage/apps/resource"
	"github.com/tqtcloud/cloud-manage/apps/secret"
)

func InjectBaseFromSecret(b *resource.Meta, s *secret.Secret) {
	// 补充管理信息
	b.CredentialId = s.Id
	b.Domain = s.Data.Domain
	b.Namespace = s.Data.Namespace
}
