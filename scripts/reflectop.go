package scripts

import "reflect"

type Person struct {
	Name string
	Age uint
	Address string
}

func showtype() {
	Info(reflect.TypeOf(3.3))  // int
	Info(reflect.TypeOf(3.3i))  // complex128
	Info(reflect.TypeOf(Person{}))  // scripts.Person
	Info(reflect.TypeOf(Person{Name: "Jim"})) // scripts.Person
}

func showvalue() {
	Info(reflect.ValueOf(3.3))
	Info(reflect.ValueOf(3.3i))
	Info(reflect.ValueOf(Person{}))
	Info(reflect.ValueOf(Person{Name: "Jim"}))
}

func changevalue() {
	x := 2
	Info(reflect.ValueOf(&x))
}

func ReflectMain() {
	showtype()
	showvalue()
	changevalue()
}