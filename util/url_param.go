package util

import (
	"reflect"
	"strconv"
	"net/url"
	"errors"
)

func UrlParameterUnMarshal(params url.Values, holder interface{}) error {
	tp := reflect.TypeOf(holder)
	if tp.Kind() != reflect.Ptr {
		return errors.New("holder must be pointer")
	}
	tp = reflect.TypeOf(holder).Elem()
	val := reflect.ValueOf(holder).Elem()

	for i := 0; i < val.NumField(); i++ {
		ft := tp.Field(i)
		tag := ft.Tag.Get("url")
		if len(tag) == 0 {
			if ft.Anonymous {
				// 如果是匿名域，递归处理
				fv := val.Field(i)
				if fv.CanAddr() {
					UrlParameterUnMarshal(params, fv.Addr().Interface())
				}
			}
			continue
		}
		require := ft.Tag.Get("require")
		fv := val.Field(i)
		x := params.Get(tag)
		if len(x) == 0 {
			if require == "true" {
				return errors.New(tag)
			}
			continue
		}

		positive := ft.Tag.Get("positive")
		if fv.Kind() == reflect.String {
			fv.SetString(x)
		} else if fv.Kind() == reflect.Int64 || fv.Kind() == reflect.Int {
			y, err := strconv.ParseInt(x, 10, 64)
			if err != nil {
				return err
			}
			if y <= 0 && positive == "true" {
				return errors.New(tag)
			}
			fv.SetInt(y)
		} else if fv.Kind() == reflect.Float64 {
			y, err := strconv.ParseFloat(x, 64)
			if err != nil {
				return err
			}
			fv.SetFloat(y)
		} else if fv.Kind() == reflect.Bool {
			y, err := strconv.ParseBool(x)
			if err != nil {
				return err
			}
			fv.SetBool(y)
		} else {
			panic(errors.New("url marshal does not support kind except string, int64, float64, int, bool"))
		}
	}
	return nil
}
