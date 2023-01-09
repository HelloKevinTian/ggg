package some

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

// HH ...
func HH(argv []int) {
	fmt.Println("func HH", argv)
}

// ReflectHandler ...
type ReflectHandler struct {
	Path string `json:"path" id:"1008"`
}

// RPrint ...
func (r ReflectHandler) RPrint(prefix string) {
	fmt.Printf("ReflectHandler.RPrint prefix[%s] Path[%s]", prefix, r.Path)
}

// TestReflect ...
func TestReflect() {
	//反射常量
	fmt.Println(reflect.Uint8, reflect.Struct, reflect.Func)

	var a int
	a = 1024
	typeOfA := reflect.TypeOf(a)
	valueOfA := reflect.ValueOf(a)
	fmt.Println(typeOfA.Name(), typeOfA.Kind())
	fmt.Println(valueOfA.Kind(), valueOfA.Type(), valueOfA.Int(), valueOfA.Interface(), valueOfA.Interface().(int))

	ht := reflect.TypeOf(HH)
	hv := reflect.ValueOf(HH)
	fmt.Println(ht.Name(), ht.Kind(), hv.Interface())

	// 声明一个结构体
	rh := ReflectHandler{Path: "/doLogin"}
	st := reflect.TypeOf(rh)
	sv := reflect.ValueOf(rh)
	fmt.Println(st.Name(), st.Kind(), sv.Interface()) //结构体类型的名称、类型、值

	//遍历结构体的属性
	for i := 0; i < st.NumField(); i++ {
		fmt.Println(st.Field(i).Index, st.Field(i).Type, st.Field(i).Name, st.Field(i).Tag)
		// 从tag中取出需要的tag
		fmt.Println(st.Field(i).Tag.Get("json"), st.Field(i).Tag.Get("id"))
	}

	//遍历结构体的方法
	for i := 0; i < st.NumMethod(); i++ {
		fmt.Println(st.Method(i).Name)
	}

	//动态调用结构体的成员方法
	for i := 0; i < sv.NumMethod(); i++ {
		args := []reflect.Value{reflect.ValueOf("前缀")}
		fmt.Println(sv.Method(i).Call(args))
	}

}

// App ...
type App struct {
	w http.ResponseWriter
}

// Say ...
func (a *App) Say(ss string) {
	fmt.Fprintf(a.w, "Say called"+ss)
}

// InvokeRouter ...
func InvokeRouter() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/public/") {
			//匹配静态文件服务
		} else {
			app := &App{w}
			rValue := reflect.ValueOf(app)
			rType := reflect.TypeOf(app)
			path := strings.Split(r.URL.Path, "/")
			controlName := path[1]
			method, exist := rType.MethodByName(controlName)
			if exist {
				args := []reflect.Value{rValue, reflect.ValueOf("asdf")}
				method.Func.Call(args)
			} else {
				fmt.Fprintf(w, "method %s not found", r.URL.Path)
			}
		}
	})

	fmt.Println("ListenAndServe :8080")
	http.ListenAndServe(":8080", nil)
}

//-------------测试反射性能-----------------

// UserReflect ...
type UserReflect struct {
}

// GetName ...
func (u UserReflect) GetName() string {
	return "joe"
}

// TestCallFunc ...
func TestCallFunc(u UserReflect) {
	u.GetName()
}

// TestCallReflect ...
func TestCallReflect(m reflect.Method, arg []reflect.Value) {
	m.Func.Call(arg)
}
