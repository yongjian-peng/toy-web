package main

import (
	"fmt"
	web "geektime/toy-web/pkg"
	"net/http"
)

func main() {
	handler := web.NewHandlerBasedOnTree().(*HandlerBasedOnTree)

	handler.Route(http.MethodPost, "/user", func(c *Context) {})

	// 开始做断言，这个时候我们应该确认，在根节点之下只有一个user节点
	fmt.Println(len(handler.root.children))

	// n := handler.root.children[0]
	// assert.NotNil(t, n)
	// assert.Equal(t, "user", n.path)
	// assert.NotNil(t, n.handler)
	// assert.Empty(t, n.children)

	// // 我们只有
	// //  user -> profile
	// handler.Route(http.MethodPost, "/user/profile", func(c *Context) {})
	// assert.Equal(t, 1, len(n.children))
	// profileNode := n.children[0]
	// assert.NotNil(t, profileNode)
	// assert.Equal(t, "profile", profileNode.path)
	// assert.NotNil(t, profileNode.handler)
	// assert.Empty(t, profileNode.children)

	// // 试试重复
	// handler.Route(http.MethodPost, "/user", func(c *Context) {})
	// n = handler.root.children[0]
	// assert.NotNil(t, n)
	// assert.Equal(t, "user", n.path)
	// assert.NotNil(t, n.handler)
	// // 有profile节点
	// assert.Equal(t, 1, len(n.children))

	// // 给 user 再加一个节点
	// handler.Route(http.MethodPost, "/user/home", func(c *Context) {})
	// assert.Equal(t, 2, len(n.children))
	// homeNode := n.children[1]
	// assert.NotNil(t, homeNode)
	// assert.Equal(t, "home", homeNode.path)
	// assert.NotNil(t, homeNode.handler)
	// assert.Empty(t, homeNode.children)

	// // 添加 /order/detail
	// handler.Route(http.MethodPost, "/order/detail", func(c *Context) {})
	// assert.Equal(t, 2, len(handler.root.children))
	// orderNode := handler.root.children[1]
	// assert.NotNil(t, orderNode)
	// assert.Equal(t, "order", orderNode.path)
}
