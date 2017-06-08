package utils

import (
	_ "fmt"
)

func Interface2Float(v interface{}) float64 {
	if v == nil {
		return 0
	}
	return v.(float64)
}

func Interface2Int(v interface{}) int64 {
	if v == nil {
		return 0
	}
	return v.(int64)
}

func Interface2String(v interface{}) string {
	if v == nil {
		return ""
	}
	return v.(string)
}
