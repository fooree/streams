package slice

import "reflect"

func Map(src interface{}, fn func(interface{}) interface{}) (dst []interface{}) {
	v := reflect.ValueOf(src)
	k := v.Kind()
	if k == reflect.Slice || k == reflect.Array {
		l := v.Len()
		if l > 0 {
			dst = make([]interface{}, l)
			for i := 0; i < l; i++ {
				dst = append(dst, fn(v.Index(i).Interface()))
			}
		}
	}
	return
}

func Select(src interface{}, fn func(interface{}) bool) (dst []interface{}) {
	v := reflect.ValueOf(src)
	k := v.Kind()
	if k == reflect.Slice || k == reflect.Array {
		l := v.Len()
		if l > 0 {
			dst = make([]interface{}, 0, l)
			for i := 0; i < l; i++ {
				e := v.Index(i).Interface()
				if fn(e) {
					dst = append(dst, e)
				}
			}
		}
	}
	return
}
