/**
  * @author wellsjiang
  * @date 2018/8/6
  */

package wellgo

import (
	"reflect"
	"encoding/json"
	"fmt"
	"strings"
)

/**
 * 将src中的值填充到dstValue中
 * src暂时为simple json解析的格式数据
 */
func AssignJsonTo(src interface{}, dstVal reflect.Value, tagName string) bool {
	sv := reflect.ValueOf(src)
	if !dstVal.IsValid() || !sv.IsValid() {
		logger.Warn("src or dstVal is invalid")
		return false
	}

	if dstVal.Kind() == reflect.Ptr {
		//初始化空指针
		if dstVal.IsNil() && dstVal.CanSet() {
			dstVal.Set(reflect.New(dstVal.Type().Elem()))
		}
		dstVal = dstVal.Elem()
	}

	// 判断可否赋值，小写字母开头的字段、常量等不可赋值
	if !dstVal.CanSet() {
		logger.Warn("dstVal can not set")
		return false
	}

	switch dstVal.Kind() {
	case reflect.Bool:
		dstVal.Set(reflect.ValueOf(src))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v, _ := src.(json.Number).Int64()
		dstVal.SetInt(v)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v, _ := src.(json.Number).Int64()
		dstVal.SetUint(uint64(v))
	case reflect.String:
		dstVal.Set(reflect.ValueOf(src))
	case reflect.Slice:
		fmt.Println("AssignJsonTo not support type reflect.Slice")
		logger.Error("AssignJsonTo not support type reflect.Slice")
		return false
	case reflect.Map:
		fmt.Println("AssignJsonTo not support type reflect.Map")
		logger.Error("AssignJsonTo not support type reflect.Map")
		return false
	case reflect.Interface:
		dstVal.Set(reflect.ValueOf(src))
	case reflect.Struct:
		if sv.Kind() != reflect.Map || sv.Type().Key().Kind() != reflect.String {
			logger.Warn("AssignJsonTo src type only support map and key is only to be string")
			return false
		}

		success := false
		for i := 0; i < dstVal.NumField(); i++ {
			fv := dstVal.Field(i)
			if fv.IsValid() == false || fv.CanSet() == false {
				continue
			}

			ft := dstVal.Type().Field(i)
			name := ft.Name
			strs := strings.Split(ft.Tag.Get(tagName), ",")
			if strs[0] == "-" { //处理ignore的标志
				continue
			}

			if len(strs[0]) > 0 {
				name = strs[0]
			}

			fsv := sv.MapIndex(reflect.ValueOf(name))
			if fsv.IsValid() {
				if fv.Kind() == reflect.Ptr && fv.IsNil() {
					pv := reflect.New(fv.Type().Elem())
					if AssignJsonTo(fsv.Interface(), pv, tagName) {
						fv.Set(pv)
						success = true
					}
				} else {
					if AssignJsonTo(fsv.Interface(), fv, tagName) {
						success = true
					}
				}
			} else if ft.Anonymous {
				//尝试对匿名字段进行递归赋值，跟JSON的处理原则保持一致
				if fv.Kind() == reflect.Ptr && fv.IsNil() {
					pv := reflect.New(fv.Type().Elem())
					if AssignJsonTo(src, pv, tagName) {
						fv.Set(pv)
						success = true
					}
				} else {
					if AssignJsonTo(src, fv, tagName) {
						success = true
					}
				}
			}
		}
		return success
	default:
		return false
	}

	return true
}

// assertion function
// default return system error
func Assert(expression bool, we ...WException) {
	if len(we) > 1 {
		panic("wellgo.Assert only allow 1 WException")
	}
	if !expression {
		var (
			e WException
		)
		if len(we) == 1 {
			e = we[0]
		} else {
			e = NewWException(ErrSystemError.Error(), GetErrorCode(ErrSystemError))
		}
		if e.Code == 0 {
			e.Code = GetErrorCode(ErrSystemError)
		}
		panic(e)
	}
}
