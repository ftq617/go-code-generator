package store

import (
	"context"

	"{{.ModName}}/{{.Abbr}}/model"
	"{{.ModName}}/{{.Abbr}}/request"
)

type {{.SupStructName}} interface {
	Create(ctx context.Context, data *model.{{.SupStructName}}) (string, error)
	Update(ctx context.Context, id string, values interface{}) error
	Delete(ctx context.Context, id string) error
	DeleteByIds(ctx context.Context, ids []string) error
	Get(ctx context.Context, id string, selectQuery ...string) (*model.{{.SupStructName}}, error)
	List(ctx context.Context, req *request.Query{{.SupStructName}}Req) ([]*model.{{.SupStructName}}, error)
}
