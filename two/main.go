package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	final := make([]string, 9)
	fmt.Println("Enter some worlds *separated by spaces*")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	arr := scanner.Text()
	fmt.Scanln(&arr)
	for _, v := range strings.Split(arr, " ") {
		for _, v2 := range v {
			v3, err := strconv.Atoi(string(v2))
			if err == nil {
				final[v3] = v
			}
		}
	}
	fmt.Println(strings.Join(delete_empty(final), " "))
}

//Delete empty slice
func delete_empty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
