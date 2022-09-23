package main

import (
	"cookieAndsession/session"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/signin", session.Signin)
	r.POST("/signin", session.Signin)
	r.GET("/welcome", session.Welcome)
	r.GET("/refresh", session.Refresh)
	r.GET("/logout", session.Logout)
	r.Run(":8080")
}
