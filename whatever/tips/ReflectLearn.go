package tips

import (
	"fmt"
	"reflect"
)


// 在两个结构体之间给同名字段赋值，要求字段需要是被导出且完全一致
// 参数 src： 源结构体
//     dst： 目的结构体
//     fields: 指定字段赋值
// 默认将源结构体的所有同名同类型的字段赋值给目的结构体，如果需要指定字段，可以作为fileds的参数传入，
//type Student struct {
//	Name   string
//	Age    int
//	Gender string
//}
// 如果只想copy Student中的 Name和Gender字段，可以
// needCopyFields := []string{"Gender", "Name"}
// CopyFields(&src, &dst, needCopyFields...)
func CopyFields(src interface{}, dst interface{}, fields ...string) (err error) {
	var errRet error
	srcVal := reflect.ValueOf(src).Elem()
	dstVal := reflect.ValueOf(dst).Elem()

	srcFields := map[string]bool{}
	if len(fields) == 0 {
		for i := 0; i < srcVal.NumField(); i++ {
			srcFields[srcVal.Type().Field(i).Name] = true
		}
	} else {
		for _, v := range fields {
			srcFields[v] = true
		}
	}

	dstFields := map[string]bool{}
	for i := 0; i < dstVal.NumField(); i++ {
		dstFields[srcVal.Type().Field(i).Name] = true
	}

	for field := range srcFields {
		if _, ok := dstFields[field]; ok {
			f := dstVal.FieldByName(field)
			val := srcVal.FieldByName(field)
			if f.IsValid() && val.IsValid() {
				// if f and val both have name field named "field"
				if f.CanSet() {
					if f.Kind() == val.Kind() {
						f.Set(val)
					} else {
						fmt.Println(fmt.Sprintf("%s in src and dst are different type\n", field))
						if errRet == nil {
							errRet = errors.New(fmt.Sprintf("%s in src and dst are different type\n", field))
						}
					}
				} else {
					fmt.Println("dst is not an address")
					if errRet == nil {
						errRet = errors.New("dst is not an address")
					}
				}
			} else {
				fmt.Println(fmt.Sprintf("%s cannot be found in src struct\n", field))
			}

		} else {
			// if field in src cannot be found in dst, just skip, no output
			//fmt.Println(fmt.Sprintf("%s cannot be found in dst struct\n", field))
		}
	}
	return errRet
}
