package controller

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/eppea/MVC/model"
	"github.com/eppea/MVC/view"

	"github.com/labstack/echo/v4"
)

var db *sql.DB

func GetAllTransactions(c echo.Context) error {
	rows, err := db.Query("SELECT id, description, amount FROM transactions")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, view.Response{Message: "Failed to get transactions"})
	}
	defer rows.Close()

	transactions := []model.Transaction{}
	for rows.Next() {
		var transaction model.Transaction
		err := rows.Scan(&transaction.ID, &transaction.Description, &transaction.Amount)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, view.Response{Message: "Failed to get transactions"})
		}
		transactions = append(transactions, transaction)
	}

	return c.JSON(http.StatusOK, transactions)
}

func GetTransaction(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, view.Response{Message: "Invalid transaction ID"})
	}

	var transaction model.Transaction
	err = db.QueryRow("SELECT id, description, amount FROM transactions WHERE id = ?", id).Scan(&transaction.ID, &transaction.Description, &transaction.Amount)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, view.Response{Message: "Transaction not found"})
		}
		return c.JSON(http.StatusInternalServerError, view.Response{Message: "Failed to get transaction"})
	}

	return c.JSON(http.StatusOK, transaction)
}

func CreateTransaction(c echo.Context) error {
	transaction := new(model.Transaction)
	if err := c.Bind(transaction); err != nil {
		return c.JSON(http.StatusBadRequest, view.Response{Message: "Invalid request payload"})
	}

	result, err := db.Exec("INSERT INTO transactions (description, amount) VALUES (?, ?)", transaction.Description, transaction.Amount)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, view.Response{Message: "Failed to create transaction"})
	}

	id, _ := result.LastInsertId()
	transaction.ID = int(id)

	return c.JSON(http.StatusCreated, transaction)
}

func UpdateTransaction(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, view.Response{Message: "Invalid transaction ID"})
	}

	transaction := new(model.Transaction)
	if err := c.Bind(transaction); err != nil {
		return c.JSON(http.StatusBadRequest, view.Response{Message: "Invalid request payload"})
	}

	result, err := db.Exec("UPDATE transactions SET description = ?, amount = ? WHERE id = ?", transaction.Description, transaction.Amount, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, view.Response{Message: "Failed to update transaction"})
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return c.JSON(http.StatusNotFound, view.Response{Message: "Transaction not found"})
	}

	return c.JSON(http.StatusOK, transaction)
}

func DeleteTransaction(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, view.Response{Message: "Invalid transaction ID"})
	}

	result, err := db.Exec("DELETE FROM transactions WHERE id = ?", id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, view.Response{Message: "Failed to delete transaction"})
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return c.JSON(http.StatusNotFound, view.Response{Message: "Transaction not found"})
	}

	return c.JSON(http.StatusOK, view.Response{Message: "Transaction deleted"})
}
