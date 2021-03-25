package main

import (
	"errors"
	"fmt"
	"testing"
)

func Fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return Fibonacci(n-2) + Fibonacci(n-1)
}

func TestHello_Say(t *testing.T) {
	a := Fibonacci(10)
	t.Log(a)
}

type strA struct {
	A int
}

func fp(lst []*strA) {
	for i, _ := range lst {
		lst[i].A = i
	}
}

func TestHello_Says(t *testing.T) {
	lst := make([]*strA, 0, 3)
	for i := 0; i < 3; i++ {
		lst = append(lst, &strA{})
	}
	fp(lst)
	t.Log(lst[0].A, lst[1].A, lst[2].A)
}
func pt(a interface{}) {
	switch a.(type) {
	case []interface{}:
		fmt.Println("array interface", a)
	default:
		fmt.Println(a)
	}
}
func Bar() error {
	return errors.New("a")
}
func Foo() (err error) {
	if err = Bar(); err != nil {
		return
	}
	return
}

func TestGetDateInt64(t *testing.T) {
	err := Foo()
	fmt.Println(err)
	//fmt.Println(a)
	//fmt.Println(a...)
}
