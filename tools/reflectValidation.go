package tools

import (
	"fmt"
	"reflect"

	"github.com/sirupsen/logrus"
)

func Validate(I interface{}) { //TODO shouldn't use logger
	defer func() {
		if r := recover(); r != nil {
			logrus.Error("Probably reflect private field")
		}
	}()
	refVal := reflect.ValueOf(I)
	typeOfRef := refVal.Type()

	for i := 0; i < refVal.NumField(); i++ {
		refValue := refVal.Field(i)
		if isZero(refValue) {
			logrus.Warn(fmt.Sprintf("%s value of %s is Zero!", typeOfRef.Field(i).Name, typeOfRef.Name()))
		}
	}
}

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Func, reflect.Map, reflect.Slice:
		return v.IsNil()
	case reflect.Array:
		z := true
		for i := 0; i < v.Len(); i++ {
			z = z && isZero(v.Index(i))
		}
		return z
		// case reflect.Struct: //! panic on time.Time type
		// 	z := true
		// 	for i := 0; i < v.NumField(); i++ {
		// 		z = z && isZero(v.Field(i))
		// 	}
		// 	return z
	}
	// Compare other types directly:
	z := reflect.Zero(v.Type())
	return v.Interface() == z.Interface()
}

//Get K is V for given tag (example: models.Config{} "db")
func DbTagInfo(I interface{}, tag string) {

	refVal := reflect.ValueOf(I)
	typeOfRef := refVal.Type()

	for i := 0; i < typeOfRef.NumField(); i++ {
		refType := typeOfRef.Field(i)
		v, ok := refType.Tag.Lookup(tag)
		if !ok {
			logrus.Warn(fmt.Sprintf("%s value of %s no DB tag", typeOfRef.Field(i).Name, typeOfRef.Name()))
			continue
		}
		logrus.Info(fmt.Sprintf("Field `%s` is `%s`", typeOfRef.Field(i).Name, v))
	}
}
