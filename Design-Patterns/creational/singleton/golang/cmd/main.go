package main

import (
	"fmt"
	"singleton/singleton"
)

func main() {
	// we cannot access the singleton type here, we have to use the singleton method
	s_1 := singleton.GetSingletonInstance()
	s_2 := singleton.GetSingletonInstance()

	s_1.SetVal(5)

	fmt.Println("s1 => ", s_1.Val())
	fmt.Println("s2 => ", s_2.Val())
}
