package auto

import (
	"time"
)

//数据库表信息
type TableInfo struct {
	TableName     string    `json:"tableName" gorm:"column:table_name"`
	TableComment  string    `json:"tableComment" gorm:"column:table_comment"`
	CreateTime    time.Time `json:"createTime" gorm:"column:create_time"`
	Checked       bool      `json:"checked"`
	SupStructName string    `json:"supStructName"`
	LowStructName string    `json:"lowStructName"`
	BaseInfo  string		`json:"baseInfo"`
	AccountInfo  string		`json:"accountInfo"`
	ProjectInfo  string		`json:"projectInfo"`
	HasTime      bool 		`json:"hasTime"`
}

//表列信息
type TableColumnInfo struct {
	ColumnName    string `json:"columnName"`
	DataType      string `json:"dataType"`
	ColumnType string `json:"columnType"`
	ColumnComment string `json:"columnComment"`

	FieldName string `json:"fieldName"`
	FieldType string `json:"fieldType"`
	FieldJson string `json:"fieldJson"`
}
