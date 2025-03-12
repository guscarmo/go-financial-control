package handlers

import (
	"go-financial-control/config"
	"go-financial-control/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateTransacao(c *gin.Context) {
	var transacao models.Transacao
	if err := c.ShouldBindJSON(&transacao); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := config.DB.Exec("INSERT INTO transacoes (descricao, categoria, valor, tipo, forma_pagamento, observacao, data) VALUES (LOWER($1), LOWER($2), $3, LOWER($4), LOWER($5), LOWER($6), $7)",
		transacao.Description, transacao.Category, transacao.Amount, transacao.Typ, transacao.Payment, transacao.Obs, transacao.Date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir transação: " + err.Error()})
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar transações: " + err.Error()})
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

func GetTransacoes(c *gin.Context) {
	campos := c.Query("campos")
	dateStart := c.Query("date-start")
	dateEnd := c.Query("date-end")

	var startDate, endDate string
	var err error

	if dateStart != "" {
		_, err = time.Parse("2006-01-02", dateStart)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Data inicial inválida"})
			return
		}
		startDate = dateStart
	}

	if dateEnd != "" {
		_, err = time.Parse("2006-01-02", dateEnd)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Data final inválida"})
			return
		}
		endDate = dateEnd
	}

	if campos == "" {
		campos = "id, descricao, categoria, valor, tipo, forma_pagamento, observacao, data"
	}

	query := "SELECT " + campos + " FROM transacoes WHERE 1=1"
	var args []interface{}
	argCount := 1

	if startDate != "" {
		query += " AND data >= $" + strconv.Itoa(argCount)
		args = append(args, startDate)
		argCount++
	}

	if endDate != "" {
		query += " AND data <= $" + strconv.Itoa(argCount)
		args = append(args, endDate)
		argCount++
	}

	rows, err := config.DB.Query(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar transações"})
		return
	}
	defer rows.Close()

	var transacoes []map[string]interface{}
	columns, _ := rows.Columns()
	values := make([]interface{}, len(columns))
	valuePointers := make([]interface{}, len(columns))

	for rows.Next() {
		for i := range columns {
			valuePointers[i] = &values[i]
		}

		if err := rows.Scan(valuePointers...); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar dados"})
			return
		}

		item := make(map[string]interface{})

		for i, col := range columns {
			item[col] = values[i]
		}
		transacoes = append(transacoes, item)
	}

	c.JSON(http.StatusOK, transacoes)
}
