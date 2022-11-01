package store

import (
	"context"

	"{{.ModName}}/{{.Abbr}}/model"
	"{{.ModName}}/{{.Abbr}}/request"
)

type {{.SupStructName}} interface {
	Create(ctx context.Context, data *model.{{.SupStructName}}) error
	Update(ctx context.Context, id uint64, values interface{}) error
	Deleted(ctx context.Context, id uint64) error
	DeletedByIds(ctx context.Context, ids []uint64) error
	Get(ctx context.Context, id uint64, selectQuery...string) (*model.{{.SupStructName}}, error)
	List(ctx context.Context, req *request.Query{{.SupStructName}}Req) ([]*model.{{.SupStructName}}, error)
}
