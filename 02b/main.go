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

func (m myS) add() {
	fmt.Println(m.a - m.b)
}

func main() {
	var a myI
	a = myS{a: 1, b: 2}
	a.add()
	a.subtract()

}
