// // remove_slice.go
// package main1

// import (
// 	"fmt"
// 	"os"
// 	"bufio"
// )

// func main() {
// 	// s := []string{"M", "N", "O", "P", "Q", "R"}
// 	// res := RemoveStringSlice(s, 2, 4)
// 	// fmt.Println(res) // [M N Q R]



// 	mapF("1.go","xxxx")
// }

// func RemoveStringSlice(slice []string, start, end int) []string {
// 	result := make([]string, len(slice)-(end-start)-1)
// 	at := copy(result, slice[:start])
// 	copy(result[at:], slice[end+1:])
// 	return result
// }





// func mapF(file string, contens string) (map[string]int){
// 	inputFile ,inputError := os.Open(file)
// 	if inputError != nil{
// 		fmt.Println("open file error")
// 		return nil
// 	}
// 	defer inputFile.Close()


// 	retmap := make(map[string]int)

// 	scanner := bufio.NewScanner(inputFile)  //  f(inputFile)

// 	for scanner.Scan(){
// 		fmt.Println(scanner.Text())
// 	}

	


// 	return retmap



// }
