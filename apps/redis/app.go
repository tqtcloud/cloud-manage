package redis

import resource "github.com/tqtcloud/cloud-manage/apps/resource"

const (
	AppName = "redis"
)

func NewSet() *Set {
	return &Set{
		Items: []*Redis{},
	}
}

func (s *Set) Add(items ...any) {
	for i := range items {
		s.Items = append(s.Items, items[i].(*Redis))
	}
}

func (s *Set) ToAny() (items []any) {
	for i := range s.Items {
		items = append(items, s.Items[i])
	}
	return
}

func (s *Set) Length() int64 {
	return int64(len(s.Items))
}

func NewDefaultRedis() *Redis {
	return &Redis{
		Resource: resource.NewDefaultResource(resource.TYPE_REDIS),
		Describe: &Describe{},
	}
}
