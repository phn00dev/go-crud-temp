package bootstrap

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/phn00dev/go-crud-temp/pkg/config"
	"github.com/phn00dev/go-crud-temp/pkg/database"
)

func Serve() {
	config.Set()
	configs := config.Get()
	database.Connection()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))
}
