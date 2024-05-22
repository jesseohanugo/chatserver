package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jesseohanugo/chatserver/database/postgres"
)

func DatabaseMiddleware(postgresDB *postgres.PostgresDB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("postgresDB", postgresDB)
		c.Next()
	}
}
