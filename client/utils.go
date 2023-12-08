package client

import (
	"fmt"
	"net/url"
	"reflect"
)

// Converts power parameters to url.Values.
// Supported power parameter value types: `string`, `int`, `bool`, `[]string`, `[]int`, `[]bool`
func powerParamsToURLValues(params map[string]interface{}) url.Values {
	qsp := url.Values{}

	for k, v := range params {
		val := reflect.ValueOf(v)
		switch val.Kind() {
		case reflect.String:
			qsp.Add(k, v.(string))
		case reflect.Int:
			qsp.Add(k, fmt.Sprintf("%d", v.(int)))
		case reflect.Bool:
			qsp.Add(k, fmt.Sprintf("%t", v.(bool)))
		case reflect.Slice:
			for i := 0; i < val.Len(); i++ {
				switch val.Index(0).Elem().Kind() {
				case reflect.String:
					qsp.Add(k, val.Index(i).Elem().String())
				case reflect.Int:
					qsp.Add(k, fmt.Sprintf("%d", val.Index(i).Elem().Int()))
				case reflect.Bool:
					qsp.Add(k, fmt.Sprintf("%t", val.Index(i).Elem().Bool()))
				}
			}
		}
	}

	return qsp
}
