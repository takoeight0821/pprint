// Example: Automatically implement Pretty() for a struct using reflection
package examples

import (
	"fmt"
	"reflect"

	"github.com/takoeight0821/pprint"
)

type Person struct {
	Name string
	Age  int
}

// PrettyStruct uses reflection to pretty-print any struct
func PrettyStruct(v any) pprint.Doc {
	rv := reflect.ValueOf(v)
	rt := reflect.TypeOf(v)
	if rv.Kind() != reflect.Struct {
		return pprint.Text(fmt.Sprintf("%v", v))
	}
	var fields []pprint.Doc
	for i := 0; i < rv.NumField(); i++ {
		fname := rt.Field(i).Name
		fval := rv.Field(i).Interface()
		fields = append(fields, pprint.Hsep(
			pprint.Text(fname+":"),
			pprint.Text(fmt.Sprintf("%v", fval)),
		))
	}
	return pprint.Vsep(fields...)
}
