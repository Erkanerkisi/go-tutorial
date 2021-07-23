package main

import "fmt"

func main() {
	arr := make([]string, 0, 3)
	fmt.Println(arr)
	arr = append(arr, "asd", "123","21q3","11")
	fmt.Println(arr)
}
