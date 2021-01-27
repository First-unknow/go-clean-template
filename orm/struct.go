package orm

import (
	"strconv"
	"sync"

	"github.com/fatih/structs"
	"github.com/spf13/cast"
)

const (
	typeUUID       = "uuid"
	typeZeroUUID   = "zerouuid"
	typeString     = "string"
	typeInt32      = "int32"
	typeInt64      = "int64"
	typeFloat64    = "float64"
	typeTimeStamp  = "timestamp"
	typeDate       = "date"
	typeZeroString = "zerostring"
	typeZeroInt    = "zeroint"
	typeZeroFloat  = "zerofloat"
)

func GetStructFields(models interface{}) (*structs.Struct, *sync.Map) {
	var ptrColumnMap = new(sync.Map)
	faithOrder := structs.New(models)
	fields := faithOrder.Fields()
	for _, f := range fields {
		tagCol := f.Tag("db")
		if tagCol != "" && tagCol != "-" {

			ptrColumnMap.Store(tagCol, f)
		}
	}

	return faithOrder, ptrColumnMap
}

func GetTableName(models interface{}) string {
	var tablename string
	faith := structs.New(models)

	if f, ok := faith.FieldOk("TableName"); ok {
		tablename = f.Tag("db")
	}

	return tablename
}

func SetFieldFromType(field *structs.Field, v interface{}) error {
	var value string
	var tag = field.Tag("type")

	value = cast.ToString(v)

	switch tag {
	case typeString:
		field.Set(value)
	case typeInt32:
		valInt, err := strconv.Atoi(value)
		if err == nil {
			field.Set(valInt)
		}
	case typeInt64:
		valInt64, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			field.Set(valInt64)
		}
	case typeFloat64:
		f, err := strconv.ParseFloat(value, 64)
		if err == nil {
			field.Set(f)
		}
	}

	return nil
}
