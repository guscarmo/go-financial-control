package routes

import (
	"go-financial-control/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	//  := gin.Default()

	router.POST("/transacoes", handlers.CreateTransacao)
	router.GET("/transacoes", handlers.ListTransacoes)
	router.GET("/transacoes/resumo", handlers.GetResumo)

	router.POST("/categorias", handlers.AddCategoria)
	router.GET("/categorias", handlers.GetCategorias)

	// return router
}
