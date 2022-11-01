package service

import (
    "context"

	"github.com/pkg/errors"
	"go.uber.org/zap"


	"{{.ModName}}/{{.Abbr}}/store"
    "{{.ModName}}/{{.Abbr}}/model"
	"{{.ModName}}/{{.Abbr}}/request"
	"{{.ModName}}/{{.Abbr}}/response"
	"{{.ModName}}/pkg/logger"
)

type {{.SupStructName}}Srv interface {
	Create(ctx context.Context, data *request.Create{{.SupStructName}}Req) error
	Update(ctx context.Context, id uint64, data *request.Update{{.SupStructName}}Req) error
	Deleted(ctx context.Context, id uint64) error
	DeletedByIds(ctx context.Context, ids []uint64) error
	Get(ctx context.Context, id uint64, selectQuery...string) (*model.{{.SupStructName}}, error)
	List(ctx context.Context, req *request.Query{{.SupStructName}}Req) (*response.List{{.SupStructName}}Res, error)
}

func New{{.SupStructName}}Srv() {{.SupStructName}}Srv {
	return &{{.LowStructName}}Srv{}
}

type {{.LowStructName}}Srv struct {
}

func (a {{.LowStructName}}Srv) Create(ctx context.Context, data *request.Create{{.SupStructName}}Req) error {
	if err := store.Client().{{.SupStructName}}().Create(ctx, data.{{.SupStructName}}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (a {{.LowStructName}}Srv) Update(ctx context.Context, id uint64, data *request.Update{{.SupStructName}}Req) error {
	if err := store.Client().{{.SupStructName}}().Update(ctx, id, data.{{.SupStructName}}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (a {{.LowStructName}}Srv) Deleted(ctx context.Context, id uint64) error {
	if err := store.Client().{{.SupStructName}}().Deleted(ctx, id); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (a {{.LowStructName}}Srv) DeletedByIds(ctx context.Context, ids []uint64) error {
	if err := store.Client().{{.SupStructName}}().DeletedByIds(ctx, ids); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (a {{.LowStructName}}Srv) Get(ctx context.Context, id uint64, selectQuery...string) (*model.{{.SupStructName}}, error) {
	{{.LowStructName}},err := store.Client().{{.SupStructName}}().Get(ctx, id, selectQuery...)
	 if err != nil {
		return nil,errors.WithStack(err)
	}
	return {{.LowStructName}},nil
}


func (a {{.LowStructName}}Srv) List(ctx context.Context, req *request.Query{{.SupStructName}}Req)  (*response.List{{.SupStructName}}Res, error) {
	{{.LowStructName}}s, err := store.Client().{{.SupStructName}}().List(ctx, req)
	if err != nil {
		logger.From(ctx).Error("The database failed to query the {{.LowStructName}} list",
			zap.Any("param", req), zap.Error(err))
		return nil, err
	}

	results := response.List{{.SupStructName}}Res{
		Pagination: model.Pagination{
			PageNum:  req.PageNum,
			PageSize: req.PageSize,
			Total:    req.Total,
		}}
	results.List = {{.LowStructName}}s
	return &results, nil
}