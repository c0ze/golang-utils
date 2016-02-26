package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func GetStructStr(lead string, strct interface{}) string {
	str := lead
	tempintslice := []int{0}
	ielements := reflect.TypeOf(strct).Elem().NumField()
	for i := 0; i < ielements; i++ {
		tempintslice[0] = i
		f := reflect.TypeOf(strct).Elem().FieldByIndex(tempintslice)
		value := reflect.Indirect(reflect.ValueOf(strct)).FieldByName(f.Name)
		str += fmt.Sprintf(" %v: %v\n", f.Name, getStringVal(value))
	}
	return str
}

func getStringVal(value reflect.Value) string {
	var retVal string
	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		retVal = strconv.FormatInt(value.Int(), 10)
	case reflect.Int64, reflect.Uint64:
		retVal = strconv.FormatUint(value.Uint(), 10)
	case reflect.String:
		retVal = value.String()
	}
	return retVal
}

func StripColon(macAddr string) string {
	return strings.ToUpper(strings.Join(strings.Split(macAddr, ":"), ""))
}
