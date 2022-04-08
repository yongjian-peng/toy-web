package main

import (
	"fmt"
	"sync"
)

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			return &user{}
		}}

	// Get 返回的是 interface{}，所以需要类型断言
	u := pool.Get().(*user)

	fmt.Printf("u.Name, u.Email")
	// defer 还回去
	defer pool.Put(u)
	fmt.Printf(u.Name, u.Email)
	// 紧接着重置 u 这个对象
	u.Reset("Tom", "my_email@qq.com")

	fmt.Printf(u.Name, u.Email)

	// 下边就是使用 u 来完成你的业务逻辑
}

type user struct {
	Name  string
	Email string
}

// 一般来说，复用对象都要求我们取出来之后，
// 重置里面的字段
func (u *user) Reset(name string, email string) {
	u.Email = email
	u.Name = name
}
