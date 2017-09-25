package main

import (
	"fmt"
)

//type Stringer interface{
//	String()string
//}

type myStr []string

func (myStr myStr) String() string {
	var out string
	for i, v := range myStr {
		out += fmt.Sprintf("Index %v: Value %v\n", i, v)
	}
	return out
}

func main() {
	s := myStr{"A", "B", "C"}
	fmt.Println(s)
	//s.pretty()
}
