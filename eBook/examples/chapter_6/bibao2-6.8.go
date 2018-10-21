package main 

import "fmt"

func f() (ret int){
	defer addret(&ret)
	defer func() {
		ret++
	}()  // 此处生命加自调用
	return 1
}
func addret(ret *int){
	*ret++
}
func main() {
	fmt.Println(f())
}
