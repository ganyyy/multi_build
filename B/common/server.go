package common

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitServer(srv string) {

	fmt.Println("run server:", srv)

	var g = gin.New()
	g.Use(gin.Recovery())
	g.Use(gin.Logger())

	g.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello",
			"server":  srv,
		})
	})

	if err := g.Run(":8080"); err != nil {
		panic(err)
	}
}
