package model

import (
    {{ if .HasTime}}"time"{{ end }}
    "{{.ModName}}/pkg/storage/mysql"
)

// {{.SupStructName}} {{.TableName}}  {{.TableComment}}
type {{.SupStructName}} struct {
    {{ if .BaseInfo}}mysql.Base{{ end }}{{if .AccountInfo }}mysql.AccountUser{{ end }}{{ if .ProjectInfo }}mysql.ProjectAndOrganization{{ end }}
    {{- range .Fields}}
    {{.FieldName}} {{.FieldType}} `json:"{{.FieldJson}}" gorm:"column:{{.FieldJson}}"` //{{.ColumnComment}}{{ end }}
}

{{ if .TableName }}
func ({{.SupStructName}}) TableName() string {
  return "{{.TableName}}"
}
{{ end }}