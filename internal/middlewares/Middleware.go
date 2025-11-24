package middlewares

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) { c.Next() }

}

/*
seharusnya dipanggil setelah authentication
baca data dari database untuk dicompare sama role
dipanggal dalam AuthMiddleware diatas
*/
func AuthorizationMiddleware() {

}

// overrider agar content-type selalu vnd.api+jeson
func DefaultHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/vnd.api+json")
		c.Next()
	}
}
