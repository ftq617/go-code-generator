package mysql

import (
    "context"

    "github.com/pkg/errors"
    "gorm.io/gorm"

    "{{.ModName}}/pkg/code"
    "{{.ModName}}/{{.Abbr}}/model"
    "{{.ModName}}/{{.Abbr}}/request"
    "{{.ModName}}/pkg/storage/mysql"
)

func new{{.SupStructName}}(db *mysql.DB) *{{.LowStructName}} {
	return &{{.LowStructName}}{
		DB: db,
	}
}

type {{.LowStructName}} struct {
	*mysql.DB
}

// Create 创建
func (r *{{.LowStructName}}) Create(ctx context.Context, data *model.{{.SupStructName}}) error {
    if err := r.With(ctx).Model(data).Create(data).Error; err != nil {
        return errors.WithStack(code.ErrDatabaseException.WithResult(err.Error()))
    }
	return nil
}
// Deleted 根据ID删除
func (r *{{.LowStructName}}) Deleted(ctx context.Context, id uint64) error {
    if err := r.With(ctx).Model(&model.{{.SupStructName}}{}).Where("id = ?", id).
    	Delete(&model.{{.SupStructName}}{}).Error; err != nil {
    	return errors.WithStack(code.ErrDatabaseException.WithResult(err.Error()))
    }
	return nil
}
// DeletedByIds 根据ID批量删除
func (r *{{.LowStructName}}) DeletedByIds(ctx context.Context, ids []uint64) error {
if err := r.With(ctx).Model(&model.{{.SupStructName}}{}).Where("id IN (?)", ids).
    	Delete(&model.{{.SupStructName}}{}).Error; err != nil {
    	return errors.WithStack(code.ErrDatabaseException.WithResult(err.Error()))
    }
	return nil
}

// Update根据id 更新 ，排除零值
func (r *{{.LowStructName}}) Update(ctx context.Context,id uint64,values interface{}) error {
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
func (r *{{.LowStructName}}) Get(ctx context.Context,id uint64, selectQuery...string) (*model.{{.SupStructName}}, error) {
    var obj model.{{.SupStructName}}
	query := r.With(ctx).Model(&model.{{.SupStructName}}{})
	if len(selectQuery) != 0 {
		query = query.Select(selectQuery[0])
	}
	if err := query.Where("id = ?", id).First(&obj).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithStack(code.ErrNotFound)
		}
		return nil, errors.WithStack(code.ErrDatabaseException.WithResult(err.Error()))
	}
    return &obj,nil
}


// List 按条件分页查询
func (r *{{.LowStructName}}) List(ctx context.Context, data *request.Query{{.SupStructName}}Req) ([]*model.{{.SupStructName}}, error) {
    var list []*model.{{.SupStructName}}
    query := r.With(ctx).Model(&model.{{.SupStructName}}{})
   	if data.ID != 0 {
   		query = query.Where("id = ?", data.ID)
   	}

   	if err := data.Build(ctx, query).Find(&list).Error; err != nil {
   		return nil, errors.WithStack(code.ErrDatabaseException.WithResult(err.Error()))
   	}

    return list, nil
}