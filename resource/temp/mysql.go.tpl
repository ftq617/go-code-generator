package mysql

import (
    "context"

    "github.com/pkg/errors"
    "gorm.io/gorm"
    codex "template/pkg/code"
    "template/pkg/storage"

    "{{.ModName}}/internal/code"
    "{{.ModName}}/{{.Abbr}}/model"
    "{{.ModName}}/{{.Abbr}}/request"
)

func new{{.SupStructName}}(db *storage.DB) *{{.LowStructName}} {
	return &{{.LowStructName}}{
		DB: db,
	}
}

type {{.LowStructName}} struct {
	*storage.DB
}

// Create 创建
func (r *{{.LowStructName}}) Create(ctx context.Context, data *model.{{.SupStructName}}) (string, error) {
    if err := r.With(ctx).Model(data).Create(data).Error; err != nil {
        return "", errors.WithStack(code.ErrDatabaseException.WithResult(err.Error()))
    }
	return data.PK(), nil
}
// Delete 根据ID删除
func (r *{{.LowStructName}}) Delete(ctx context.Context, id string) error {
    if err := r.With(ctx).Model(&model.{{.SupStructName}}{}).Where("id = ?", id).
    	Delete(&model.{{.SupStructName}}{}).Error; err != nil {
    	return errors.WithStack(code.ErrDatabaseException.WithResult(err.Error()))
    }
	return nil
}
// DeleteByIds 根据ID批量删除
func (r *{{.LowStructName}}) DeleteByIds(ctx context.Context, ids []string) error {
if err := r.With(ctx).Model(&model.{{.SupStructName}}{}).Where("id IN (?)", ids).
    	Delete(&model.{{.SupStructName}}{}).Error; err != nil {
    	return errors.WithStack(code.ErrDatabaseException.WithResult(err.Error()))
    }
	return nil
}

// Update 根据id 更新 ，排除零值
func (r *{{.LowStructName}}) Update(ctx context.Context,id string,values interface{}) error {
	query := r.With(ctx).Model(&model.{{.SupStructName}}{}).
   		Where("id = ?", id).Updates(values)
   	if err := query.Error; err != nil {
   		return errors.WithStack(code.ErrDatabaseException.WithResult(err.Error()))
   	}
   	if query.RowsAffected == 0 {
   		return errors.WithStack(code.ErrNoRowsAffected)
   	}
	return nil
}

// Get (id int64)  根据id获取model
func (r *{{.LowStructName}}) Get(ctx context.Context,id string, selectQuery ...string) (*model.{{.SupStructName}}, error) {
    var obj model.{{.SupStructName}}
	query := r.With(ctx).Model(&model.{{.SupStructName}}{})
	if len(selectQuery) != 0 {
		query = query.Select(selectQuery[0])
	}
	if err := query.Where("id = ?", id).First(&obj).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithStack(codex.ErrNotFound.WithResult("{{.SupStructName}} not found, id :" + id))
		}
		return nil, errors.WithStack(code.ErrDatabaseException.WithResult(err.Error()))
	}
    return &obj,nil
}


// List 按条件分页查询
func (r *{{.LowStructName}}) List(ctx context.Context, data *request.Query{{.SupStructName}}Req) ([]*model.{{.SupStructName}}, error) {
    var list []*model.{{.SupStructName}}
    query := r.With(ctx).Model(&model.{{.SupStructName}}{})
    if data.Select != "" {
    		query = query.Select(data.Select)
    }
   	if data.ID != 0 {
   		query = query.Where("id = ?", data.ID)
   	}

   	if err := data.Build(ctx, query).Find(&list).Error; err != nil {
   		return nil, errors.WithStack(code.ErrDatabaseException.WithResult(err.Error()))
   	}

    return list, nil
}