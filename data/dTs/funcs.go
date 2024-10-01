package dTs

import "reflect"

func GetInters(slice interface{}) []Inter {
	val := reflect.ValueOf(slice)

	inters := make([]Inter, 0, val.Len())

	for i := 0; i < val.Len(); i++ {
		v := val.Index(i).Interface()

		if v == nil {
			continue
		}

		if item, ok := v.(Inter); ok {
			inters = append(inters, item)
		}
	}

	return inters
}
