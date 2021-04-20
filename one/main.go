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
	isEven := true
	for _, v := range strings.Split(arr, ",") {
		v2, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		if v2 != 0 && v2 > 0 {
			isEven = false
		}

	}

	if isEven {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}
