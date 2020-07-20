package reflection

import "reflect"

// walk : calls a given function with a string parameter when one is found in an interface
// uses recursion to "walk" through an interface properties
func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Slice, reflect.Array:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walk(v.Interface(), fn)
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walk(res.Interface(), fn)
		}
	}
}

// getValue : Gets the value from an interface using reflect, if its a pointer it will extract the value
func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	// If its a pointer extract the underlying value
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
