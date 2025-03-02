package sFuncs

import "reflect"

// IsZero checks if a struct is the zero of its type
func IsZero(a any) bool {
	// We get the value of the underlying struct, by dereferencing the pointer
	v := reflect.ValueOf(a).Elem()

	// Iterate over each struct field and check if its value is the zero of type
	for i := range v.NumField() {
		if !v.Field(i).IsZero() {
			return false
		}
	}
	return true
}
