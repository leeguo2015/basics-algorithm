package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Get(c *gin.Context) {
	fmt.Println(c.Param("userId"))
	c.String(http.StatusOK, "User!")
}

func Post(c *gin.Context) {
	c.String(http.StatusOK, "User!")
}
