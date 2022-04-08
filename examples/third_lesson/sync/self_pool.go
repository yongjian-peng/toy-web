package main

import (
	"fmt"
	"sync"
)

func main() {
	// 建立对象
	var pipe = &sync.Pool{New: func() interface{} { return "Hello Beijing11" }}

	// 准备放入字符串
	val := "Hello,World"

	// 放入
	pipe.Put(val)

	first := pipe.Get().(string)

	fmt.Println(first)

	// 再取就没有了，会自动调用 NEW
	second := pipe.Get().(string)

	fmt.Println(second)

	three := pipe.Get().(string)

	fmt.Println(three)
}
