package utils

import (
	"fmt"
	"html/template"
	"reflect"
	"strconv"
)

// {{ map_get m "a" }} // return 1
// {{ map_get m 1 "c" }} // return 4
func MapGet(arg1 interface{}, arg2 ...interface{}) (interface{}, error) {
	arg1Type := reflect.TypeOf(arg1)
	arg1Val := reflect.ValueOf(arg1)

	if arg1Type.Kind() == reflect.Map && len(arg2) > 0 {
		// check whether arg2[0] type equals to arg1 key type
		// if they are different, make conversion
		arg2Val := reflect.ValueOf(arg2[0])
		arg2Type := reflect.TypeOf(arg2[0])
		if arg2Type.Kind() != arg1Type.Key().Kind() {
			// convert arg2Value to string
			var arg2ConvertedVal interface{}
			arg2String := fmt.Sprintf("%v", arg2[0])

			// convert string representation to any other type
			switch arg1Type.Key().Kind() {
			case reflect.Bool:
				arg2ConvertedVal, _ = strconv.ParseBool(arg2String)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				arg2ConvertedVal, _ = strconv.ParseInt(arg2String, 0, 64)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				arg2ConvertedVal, _ = strconv.ParseUint(arg2String, 0, 64)
			case reflect.Float32, reflect.Float64:
				arg2ConvertedVal, _ = strconv.ParseFloat(arg2String, 64)
			case reflect.String:
				arg2ConvertedVal = arg2String
			default:
				arg2ConvertedVal = arg2Val.Interface()
			}
			arg2Val = reflect.ValueOf(arg2ConvertedVal)
		}

		storedVal := arg1Val.MapIndex(arg2Val)

		if storedVal.IsValid() {
			var result interface{}

			switch arg1Type.Elem().Kind() {
			case reflect.Bool:
				result = storedVal.Bool()
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				result = storedVal.Int()
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				result = storedVal.Uint()
			case reflect.Float32, reflect.Float64:
				result = storedVal.Float()
			case reflect.String:
				result = storedVal.String()
			default:
				result = storedVal.Interface()
			}

			// if there is more keys, handle this recursively
			if len(arg2) > 1 {
				return MapGet(result, arg2[1:]...)
			}
			return result, nil
		}
		return nil, nil

	}
	return nil, nil
}

func RegisterTplFunc() template.FuncMap {
	tplFuncMap := make(template.FuncMap)
	tplFuncMap["map_get"] = MapGet
	return tplFuncMap
}
