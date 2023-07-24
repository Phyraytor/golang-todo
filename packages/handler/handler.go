package handler

import (
	"github.com/Phyraytor/golang-todo/packages/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	//router.Use(cors.Default())
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	api := router.Group("/api")
	{
		api.POST("/sign-up", h.signUp)
		api.GET("/foodstuff", h.getAllFoodstuff)
		api.GET("/foodstuff/:id", h.getFoodstuffById)
		api.POST("/foodstuff", h.createFoodstuff)
		api.PUT("/foodstuff/:id", h.updateFoodstuff)
		api.DELETE("/foodstuff/:id", h.deleteFoodstuff)
	}
	return router
}
