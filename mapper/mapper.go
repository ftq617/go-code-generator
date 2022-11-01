package mapper

import (
	"code/gen/auto"
	"code/gen/util/conf"
	"errors"
)

//获取数据库所有的表
func GetTables(dbName string) (data []auto.TableInfo, err error) {
	if conf.DB == nil {
		return data,errors.New("未找到数据连接")
	}
	err = conf.DB.Raw("select table_name as table_name,table_comment as table_comment,CREATE_TIME create_time from information_schema.tables where table_schema = ? ", dbName).Scan(&data).Error
	return
}

func GetColumns(dbName, tableName string) (data []auto.TableColumnInfo, err error) {
	if conf.DB == nil {
		return data, errors.New("未找到数据连接")
	}
	err = conf.DB.Raw("SELECT COLUMN_NAME column_name,DATA_TYPE data_type,COLUMN_TYPE column_type,COLUMN_COMMENT column_comment FROM INFORMATION_SCHEMA.COLUMNS c WHERE table_name = ? AND table_schema = ?", tableName, dbName).Scan(&data).Error
	return
}
