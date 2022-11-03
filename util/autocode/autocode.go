package autocode

import (
	"strings"
)

const (
	Unsigned = "unsigned"

	Varchar = "varchar"
	Datetime = "datetime"
	Timestamp = "timestamp"
	Double = "double"
	Float = "float"
	Decimal = "decimal"
	Int = "int"
	Tinyint = "tinyint"
	Smallint = "smallint"
	Mediumint = "mediumint"
	Integer = "integer"
	Bigint = "bigint"
)

const (
	String = "string"
	Time = "time.Time"
	Float64 = "float64"
	Int64 = "int64"
	Uint64 = "uint64"
	Intg = "int"
	Uintg = "uint"
)

func GetDbType(dt,ct string) string {
	switch dt {
	case Varchar:
		return String
	case Datetime:
		return Time
	case Timestamp:
		return Time
	case Double:
		return Float64
	case Float:
		return Float64
	case Decimal:
		return Float64
	case Int:
		if strings.Contains(strings.ToLower(ct),Unsigned) {
			return Uintg
		}
		return Intg
	case Tinyint:
		if strings.Contains(strings.ToLower(ct),Unsigned) {
			return Uintg
		}
		return Intg
	case Smallint:
		if strings.Contains(strings.ToLower(ct),Unsigned) {
			return Uintg
		}
		return Intg
	case Mediumint:
		if strings.Contains(strings.ToLower(ct),Unsigned) {
			return Uint64
		}
		return Int64
	case Integer:
		if strings.Contains(strings.ToLower(ct),Unsigned) {
			return Uint64
		}
		return Int64
	case Bigint:
		if strings.Contains(strings.ToLower(ct),Unsigned) {
			return Uint64
		}
		return Int64
	default:
		return String
	}

}

const DeletedField = "Deleted"
const DeletedType = "mysql.Deleted"

func GetFieldType(field,dt,ct string) string {
	dtype := GetDbType(dt,ct)
	// deleted 处理逻辑
	if field == DeletedField && dtype == Uint64 {
		return DeletedType
	}
	return dtype
}
