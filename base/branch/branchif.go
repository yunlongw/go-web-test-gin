package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "base/branch/abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}
