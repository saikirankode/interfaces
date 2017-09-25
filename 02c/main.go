package main

import "fmt"

type myI interface {
	add()
	subtract()
}

type myS struct {
	a, b int
}

func (m myS) subtract() {
	fmt.Println(m.a - m.b)
}

func main() {
	var a myI
	d:= myS{1,2}
	a = d
	a.add()
	a.subtract()

}

//.\main.go:21: cannot use d (type myS) as type myI in assignment:
//myS does not implement myI (missing add method)

