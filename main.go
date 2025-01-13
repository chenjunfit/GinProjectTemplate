package main

import (
	"GinProjectTemplate/config"
	_ "GinProjectTemplate/config"
	"GinProjectTemplate/middlewares/auth"
	"GinProjectTemplate/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(auth.JWTCheck)
	routers.RegisterRouters(r)
	r.Run(config.Port)
}
