package main

import "fmt"

type T1 struct {
	name string
}

type T2 struct {
	name string
}

func main() {
	vs := []interface{}{T2(T1{"foo"}), string(322), []byte("abł")}

	for _, v := range vs {
		fmt.Printf("%v %T\n", v, v)
	}
}
