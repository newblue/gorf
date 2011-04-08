package pkg2

import (
	"fmt"
	"package1"
)
//a comment
func Bar() {
	fmt.Println("pkg2::Bar")
}

type T struct {
	A, b int
	//comment here too
}

func (t *T) Foo() {
	package1.Foo()
}

func (t *T) Bar() {
	Bar()
}

func Baz() {
     fmt.Println("pkg2::Baz")
}