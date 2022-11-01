package response

import (
	"{{.ModName}}/{{.Abbr}}/model"
)

type List{{.SupStructName}}Res struct {
	model.Pagination
	List []*model.{{.SupStructName}} `json:"list"`
}