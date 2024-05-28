package main

import "fmt"

type inner struct {
	a int 
	b int
}
type outer struct {
	c int 
	inner
}
type hasInner interface {
	getInner() *inner
}
var _ hasInner = &outer{}
func (o *outer) getInner() *inner {
	return &o.inner
}


func newOuter(options ...func(*outer)) outer {
	o := outer{}
	for _, opt := range options {
		opt(&o)
	}
	return o
}
func withC(c int) func(*outer) {
	return func(o *outer) {
		o.c = c
	}
}
func withA[T hasInner](a int) func(T) {
	return func(t T) {
		t.getInner().a = a
	}
}
func withB[T hasInner](b int) func(T) {
	return func(t T) {
		t.getInner().b = b
	}
}

func main() {
	o := newOuter(withC(1), withA[*outer](2), withB[*outer](3))
	fmt.Println("jo")
	fmt.Println(o)
}