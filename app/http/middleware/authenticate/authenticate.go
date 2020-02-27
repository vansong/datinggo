package authenticate

import (
	"github.com/gin-gonic/gin"
)

func SetUp() gin.HandlerFunc {

	return func(c *gin.Context) {
		// TODO: verify login status && pass login info to context

		c.Set("loggedin", 3)

		c.Next()
	}
}