package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file, err := ioutil.ReadFile("test.txt")

	if err != nil {
		return
	}

	str := string(file)
	fmt.Println(str)
}