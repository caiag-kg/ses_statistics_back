package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"ses_back/internal/models"
)

// GetEvents retrieves all events from the database.
func GetEvents(ctx *gin.Context) {
	data, err := DBselect("")
	if err != nil {
		ctx.IndentedJSON(http.StatusNoContent, gin.H{"error": err.Error()})
	}
	ctx.IndentedJSON(http.StatusOK, data)
}

// getConditionsPrompt generates a SQL query string based on the provided filters.
func getConditionsPrompt(filters *models.DBFilterModel) string {
	var query string

	if filters.Search.Column != "" && filters.Search.Value != "" && filters.Search.Condition == "" {
		query += fmt.Sprintf(" WHERE %s ILIKE %s", filters.Search.Column, "'%"+filters.Search.Value+"%'")
	}

	if filters.Search.Column != "" && filters.Search.Value != "" && filters.Search.Condition != "" {
		query += fmt.Sprintf(
			" WHERE %s %s %s",
			filters.Search.Column,
			filters.Search.Condition,
			filters.Search.Value,
		)
	}

	if filters.SortBy.Column != "" {
		query += " ORDER BY " + filters.SortBy.Column + " " + filters.SortBy.Order
	}

	if filters.Limit > 0 {
		query += fmt.Sprintf(" LIMIT %d", filters.Limit)
	}

	if filters.Offset > 0 {
		query += fmt.Sprintf(" OFFSET %d", filters.Offset)
	}

	return query
}

// getDataByFilter retrieves data from the database based on the provided filters.
func getDataByFilter(filters *models.DBFilterModel) []models.Bulls {
	var conditions = getConditionsPrompt(filters)
	var data, err = DBselect(conditions)
	if err != nil {
		log.Panic(err)
	}
	return data
}

// GetEvByFilter handles HTTP requests to retrieve events by filter.
func GetEvByFilter(ctx *gin.Context) {
	// Input:
	var inputData models.DBFilterModel
	if err := ctx.BindJSON(&inputData); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := getDataByFilter(&inputData)

	// Output:
	ctx.JSON(http.StatusOK, gin.H{"data": response})
}