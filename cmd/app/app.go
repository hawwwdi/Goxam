package main

import (
	//"github.com/gin-gonic/gin"
	"github.com/hawwwdi/Goxam/internal/models/start"
)

func main() {
	go 	start.Start()
	/*r := gin.Default()
	r.LoadHTMLGlob("views/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200,"index.tmpl","dfa")
	})
	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")*/
}
