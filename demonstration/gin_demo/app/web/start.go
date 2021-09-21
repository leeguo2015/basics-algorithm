package web

import (
	"github.com/gin-gonic/gin"
	"go-integrated/demonstration/gin_demo/app/web/url"
	"net/http"
)

func StartServer()  {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "",
			"data": 123,
		})
	})
	//404 Handler.
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound,"the incorrect API route.")
	})

	ApiHandler := r.Group("/api")
	url.ApiInclude(ApiHandler)
	_ = r.Run(":8000")
}
