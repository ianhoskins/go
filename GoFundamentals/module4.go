package main

import (
	"fmt"
	"os"
	"reflect"
)

var name1 string = "Hoskins"

func main() {

	name := "Ian"
	course := "Go Fundamentals"
	module := 1
	ptr := &module
	const pi = 3.1415

	fmt.Println("Name is", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Course is", course, "and is of type", reflect.TypeOf(course))
	fmt.Println("Module is", module, "and is of type", reflect.TypeOf(module))
	fmt.Println("Const PI", pi, "and is of type", reflect.TypeOf(module))
	fmt.Println("OS Environment", os.Environ)
	fmt.Println("Memory Address of *module* is", &module)
	fmt.Println("Value of pointer *ptr* is", *ptr)
	fmt.Println("Value of pointer name1 is", name1)

	changeCourse1(course)
	fmt.Println("1-You are watching course", course)

	changeCourse2(&course)
	fmt.Println("2-You are watching course", course)
}

func changeCourse1(course string) string {
	course = "Kicking Ass"

	fmt.Println("\n1Func-Trying to change your course to", course)

	return course
}

func changeCourse2(course *string) string {
	*course = "Kicking Ass"

	fmt.Println("\n2Func-Trying to change your course to", *course)

	return *course
}
