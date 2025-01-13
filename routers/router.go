package routers

import (
	"GinProjectTemplate/routers/auth"
	"github.com/gin-gonic/gin"
)

func RegisterRouters(r *gin.Engine) {
	apiGroup := r.Group("/api")
	auth.RegisterSubRouter(apiGroup)
}
