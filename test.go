package main

import (
	"reflect"
	"strconv"
)

func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}
