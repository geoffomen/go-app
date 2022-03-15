package ginimp

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	// cors "github.com/rs/cors/wrapper/gin"
)

func enableCors() gin.HandlerFunc {
	op := cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:     []string{"*"},
		AllowCredentials:   true,
		OptionsPassthrough: true,
	}
	cr := cors.New(op)

	return func(c *gin.Context) {
		cr.HandlerFunc(c.Writer, c.Request)
		if !op.OptionsPassthrough && c.Request.Method == http.MethodOptions && c.GetHeader("Access-Control-Request-Method") != "" {
			// Abort processing next Gin middlewares.
			c.AbortWithStatus(http.StatusOK)
		}
	}
}
