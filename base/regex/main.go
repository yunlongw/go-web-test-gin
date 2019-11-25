package main

import (
	"fmt"
	"regexp"
)

const text = `My email is ccmouse@gmail.com 

mail@dd.com asdfadsf@qq.com.cn
`

func main() {
	//re := regexp.MustCompile(`[a-zA-Z0-9_.]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`)
	//match := re.FindAllString(text, -1)

	//子匹配
	re := regexp.MustCompile(`([a-zA-Z0-9_.]+)@([a-zA-Z0-9]+)(\.([a-zA-Z0-9.]+))`)
	match := re.FindAllStringSubmatch(text, -1)

	fmt.Println(match)
}
