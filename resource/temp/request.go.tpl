package request

import (
	"{{.ModName}}/{{.Abbr}}/model"
)

type Query{{.SupStructName}}Req struct {
	model.ListQuery
	model.{{.SupStructName}}
}

type Create{{.SupStructName}}Req struct {
	model.{{.SupStructName}}
}

type Update{{.SupStructName}}Req struct {
	model.{{.SupStructName}}
}