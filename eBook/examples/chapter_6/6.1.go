


package main

import "fmt"
import "strings"




func (st *string) Pop() int {
    v := 0
    for ix := len(st) - 1; ix >= 0; ix-- {
        if v = st[ix]; v != 0 {
            st[ix] = 0
            return v
        }
    }
}    


func main(){
	Pop()
}