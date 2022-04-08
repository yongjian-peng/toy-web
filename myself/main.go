package main

import (
	"fmt"
	"net/http"

	"../pkg/web"
)

func main() {
	handler := web.NewHandlerBasedOnTree().(*web.HandlerBasedOnTree)

	postNode := handler.forest[http.MethodPost]

	err := handler.Route(http.MethodPost, "/user", func(c *web.Context) {})

	fmt.Println(err)

	n := postNode.children[0]

	fmt.Printf(n)
}
