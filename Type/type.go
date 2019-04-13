package Type

import (
	"strings"
	"reflect"
)
//数据库操作所需
type DBOperation interface {
	TableName() string
	NewInterface() DBOperation
	//Count() TableType
}
//json转换的时候需要用到
type IHandler interface {
}
//数据库操作需要，不能用基本类型
type TableType interface {
	TableName() string
	Type() string
	Value() interface{}
	String() string
	Int() int64
	Float() float64
	IsNil() bool
	SetValue(i interface{})
	Name() string
}

func IsTabelType(type_ reflect.Type) bool {
	return strings.HasPrefix(type_.String(), "*Type.") || strings.HasPrefix(type_.String(), "Type.")
}

