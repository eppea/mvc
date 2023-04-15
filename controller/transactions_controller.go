package controller

import (
	"mvc/config"
	"mvc/model"
	"mvc/view"
	"net/http"
	"strconv"

	// "github.com/eppea/mvc/model"
	// "github.com/eppea/mvc/view"

	"github.com/labstack/echo/v4"
)

func GetAllTransactions(c echo.Context) error {
	var transactions []model.Transaction
	result := config.DB.Find(&transactions)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, view.Response{Message: "Failed to get transaction"})
	}
	return c.JSON(http.StatusOK, transactions)
}

func GetTransaction(c echo.Context) error {
	// Struct
	// db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 ORDER BY id LIMIT 1;
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, view.Response{Message: "Invalid transaction ID"})
	}

	var transaction model.Transaction
	result := config.DB.Where(&model.Transaction{ID: id}).Find(&transaction)
	// err = config.DB.Raw("SELECT id, description, amount FROM transactions WHERE id = ?", id).Scan(&transaction.ID, &transaction.Description, &transaction.Amount)
	if result.Error != nil {
		// if err == sql.ErrNoRows {
		// 	return c.JSON(http.StatusNotFound, view.Response{Message: "Transaction not found"})
		// }
		return c.JSON(http.StatusInternalServerError, view.Response{Message: "Failed to get transaction"})
	}

	return c.JSON(http.StatusOK, transaction)
}

func CreateTransaction(c echo.Context) error {
	transaction := new(model.Transaction)
	if err := c.Bind(transaction); err != nil {
		return c.JSON(http.StatusBadRequest, view.Response{Message: "Invalid request payload"})
	}
	if transaction.Description == "" && transaction.Amount == 0 {
		return c.JSON(http.StatusBadRequest, view.Response{Message: "Invalid request payload"})
	}

	result := config.DB.Create(&transaction)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, view.Response{Message: "Failed to create transaction"})
	}

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
	// b.Model(&User{}).Where("active = ?", true).Update("name", "hello")
	result := config.DB.Model(&transaction).Where("id = ?", id).Updates(transaction)
	// ("UPDATE transactions SET description = ?, amount = ? WHERE id = ?", transaction.Description, transaction.Amount, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, view.Response{Message: "Failed to update transaction"})
	}

	rowsAffected := result.RowsAffected
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

	result := config.DB.Delete(&model.Transaction{}, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, view.Response{Message: "Failed to delete transaction"})
	}

	rowsAffected := result.RowsAffected
	if rowsAffected == 0 {
		return c.JSON(http.StatusNotFound, view.Response{Message: "Transaction not found"})
	}

	return c.JSON(http.StatusOK, view.Response{Message: "Transaction deleted"})
}
