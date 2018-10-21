// map_func.go
package main

import "fmt"

func main() {
	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		5: func() int { return 50 },
	}

	fmt.Println(mf)
	
	mf1 := map[int] string{ 1:"xx",2:"cc",3:"bb"}
	fmt.Println(mf1)
}

// Output:  map[1:0x10903be0 5:0x10903ba0 2:0x10903bc0]
