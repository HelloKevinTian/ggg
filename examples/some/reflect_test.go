package some

import (
	"reflect"
	"testing"
)

var u UserReflect = UserReflect{}
var reflectT = reflect.TypeOf(u)
var reflectV = reflect.ValueOf(u)
var m, _ = reflectT.MethodByName("GetName")

func BenchmarkCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TestCallFunc(u)
	}
}

func BenchmarkReflectCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TestCallReflect(m, []reflect.Value{reflectV})
	}
}
