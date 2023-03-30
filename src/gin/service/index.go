package service

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"test/src/gin/model"
)

func GetError(err error, form *model.Info) string {
	//断言.判断是否为err类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		formType := reflect.TypeOf(form)
		for _, v := range errs {
			if StructField, ok := formType.Elem().FieldByName(v.Field()); ok {
				return StructField.Tag.Get("msg")
			}
		}

	}
	return err.Error()
}
