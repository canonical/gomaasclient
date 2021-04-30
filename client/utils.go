package client

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

// BUG(onwsk8r) JSON does not use unexported fields, and may not convert field names to lower case
// BUG(onwsk8r) ToQSP converts bools to string("true") or string("false"), API wants 0 or 1 <- Verify this

// ToQSP represents a struct as query string parameters.
// Given a struct value as input, it will add each field to a url.Values to be
// printed in the form field=val. Elements of array values are handled via Add().
//
// The behavior of field names is similar to that of the json package: names are
// converted to lower case, the value of a json struct tag will be used if present,
// and the field will be excluded if it has a json tag of "-". Field values are
// printed with fmt.Sprint() to create a string representation; this will work
// properly for simple data types such as int and string.
//
// The function will panic if the input is not a struct, including on a pointer
// to a struct. As this function relies heavily on reflection, it may panic under
// other circumstances as well. It is only expected to work with simple structs
// that represent query string parameters as Go types; other use cases are not
// supported.
func ToQSP(t interface{}) url.Values {
	st := reflect.TypeOf(t)
	sv := reflect.ValueOf(t)
	if st.Kind() == reflect.Ptr {
		st = st.Elem()
		sv = sv.Elem()
	}
	qsp := url.Values{}

	for i := 0; i < st.NumField(); i++ {
		// Get the name of the QSP
		key := st.Field(i).Name
		if tag, ok := st.Field(i).Tag.Lookup("json"); ok {
			if tag == "-" {
				continue
			}
			key = tag
		}
		key = strings.ToLower(key)

		// Parse out the values
		field := sv.Field(i)
		if field.Kind() == reflect.Array || field.Kind() == reflect.Slice {
			for j := 0; j < field.Len(); j++ {
				qsp.Add(key, fmt.Sprint(field.Index(j)))
			}
		} else {
			qsp.Set(key, fmt.Sprint(field))
		}
	}
	return qsp
}
