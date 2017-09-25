package main

import "fmt"

func anytype(a interface{}) {
	fmt.Println(a)
}

func main() {
	anytype(1)
	anytype("hello")
	anytype(4.5)
	anytype([]string{"a", "b"})
	anytype(map[string]string{"Name":"Satya", "id": "KLMN"})
}
