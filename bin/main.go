package main

import (
	"bin/types"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	if err := g.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		fmt.Printf("set trusted proxies failed, err:%v\n", err)
	}
	types.RegisterRouter(g)
	if err := g.Run(":8080"); err != nil {
		fmt.Printf("service failed, err:%v\n", err)
	}
}
