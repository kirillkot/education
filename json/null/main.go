package main

import (
	"encoding/json"
	"fmt"
	"runtime"
)

type bar struct {
	Name string
	Info json.RawMessage
}

func main() {
	fmt.Println("Hello, playground", runtime.Version())
	foo := bar{Name: "Gopher"}
	fmt.Println("Before marshal ", foo)
	buf, err := json.Marshal(foo)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Marshal json", string(buf))

	var remoteFoo bar
	err = json.Unmarshal(buf, &remoteFoo)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Unmarshal Foo", foo)
	fmt.Println("Unmarshal RemoteFoo", remoteFoo)

}
