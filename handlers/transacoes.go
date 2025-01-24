package handlers

import (
	"go-financial-control/config"
	"go-financial-control/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTransacao(c *gin.Context) {
	var transacao models.Transacao
	if err := c.ShouldBindJSON(&transacao); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := config.DB.Exec("INSERT INTO transacoes (descricao, categoria, valor, tipo, forma_pagamento, observacao, data) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		transacao.Description, transacao.Category, transacao.Amount, transacao.Typ, transacao.Payment, transacao.Obs, transacao.Date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir transação"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Transação criada com sucesso"})
}

func ListTransacoes(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, descricao, categoria, valor, tipo, forma_pagamento, observacao, data FROM transacoes")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar transações"})
		return
	}
	defer rows.Close()

	var transacoes []models.Transacao
	for rows.Next() {
		var t models.Transacao
		if err := rows.Scan(&t.ID, &t.Description, &t.Category, &t.Amount, &t.Typ, &t.Payment, &t.Obs, &t.Date); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar transações"})
			return
		}
		transacoes = append(transacoes, t)
	}

	c.JSON(http.StatusOK, transacoes)
}

func GetResumo(c *gin.Context) {
	var totalGanhos, totalCustos float64
	err := config.DB.QueryRow("SELECT COALESCE(SUM(valor), 0) FROM transacoes WHERE tipo = 'ganho'").Scan(&totalGanhos)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao calcular ganhos"})
		return
	}

	err = config.DB.QueryRow("SELECT COALESCE(SUM(valor), 0) FROM transacoes WHERE tipo = 'custo'").Scan(&totalCustos)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao calcular custos"})
		return
	}

	resumo := gin.H{
		"total_ganhos": totalGanhos,
		"total_custos": totalCustos,
		"saldo":        totalGanhos - totalCustos,
	}

	c.JSON(http.StatusOK, resumo)
}
