package auth

import (
	"GinProjectTemplate/controllers/auth"
	"github.com/gin-gonic/gin"
)

func RegisterSubRouter(routerGroup *gin.RouterGroup) {
	authGroup := routerGroup.Group("/auth")
	Login(authGroup)
	Logout(authGroup)

}
func Login(authGroup *gin.RouterGroup) {
	authGroup.POST("/login", auth.Login)
}
func Logout(authGroup *gin.RouterGroup) {
	authGroup.GET("/logout", auth.Logout)
}
