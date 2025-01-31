package handlers

import (
	"go-financial-control/config"
	"go-financial-control/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategorias(c *gin.Context) {
	rows, err := config.DB.Query("SELECT ID, categoria FROM categorias")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar categorias"})
		return
	}
	defer rows.Close()

	var categorias []models.Categoria
	for rows.Next() {
		var ctg models.Categoria
		if err := rows.Scan(&ctg.ID, &ctg.Category); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar categorias"})
			return
		}
		categorias = append(categorias, ctg)
	}

	c.JSON(http.StatusOK, categorias)
}
