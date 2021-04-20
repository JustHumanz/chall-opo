package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var arr string
	fmt.Println("Enter array *separated by commas*")
	fmt.Scanln(&arr)
	finnal := 0
	tmp := 0
	for _, v := range strings.Split(arr, ",") {
		v2, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		if v2 >= 0 && tmp >= 0 {
			finnal = v2 + tmp
		}
		tmp = v2
	}
	fmt.Println(finnal)
}
