package route

import (
	"github.com/gin-gonic/gin"
	"github.com/xumingcheng/gin_application/controller"
	"github.com/xumingcheng/gin_application/middleware"
)

func CollectRoute(r *gin.Engine)*gin.Engine{
	r.Use(middleware.CORSMiddleware(),middleware.RecoverMiddleware())
	r.POST("/api/auth/register",controller.Register)
	r.POST("/api/auth/login",controller.Login)
	r.GET("/api/auth/info",middleware.AuthMiddleware(),controller.Info)
	return r
}
