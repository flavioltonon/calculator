package controller

import (
	"encoding/json"
	"net/http"

	"calculator/application/http/dto"
	"calculator/domain/usecases"
	"calculator/domain/valueobject"
	"calculator/infrastructure/response"
)

type Controller struct {
	calculator usecases.Calculator
}

func New(calculator usecases.Calculator) *Controller {
	return &Controller{
		calculator: calculator,
	}
}

func (c *Controller) Add(w http.ResponseWriter, r *http.Request) {
	var operands dto.Operands

	if err := json.NewDecoder(r.Body).Decode(&operands); err != nil {
		response.JSON(w, http.StatusBadRequest, response.Error(err))
		return
	}

	result := c.calculator.Add(
		valueobject.Number(operands.X),
		valueobject.Number(operands.Y),
	)

	response.JSON(w, http.StatusOK, response.Success(dto.NewCalculation(result)))
}

func (c *Controller) Subtract(w http.ResponseWriter, r *http.Request) {
	var operands dto.Operands

	if err := json.NewDecoder(r.Body).Decode(&operands); err != nil {
		response.JSON(w, http.StatusBadRequest, response.Error(err))
		return
	}

	result := c.calculator.Subtract(
		valueobject.Number(operands.X),
		valueobject.Number(operands.Y),
	)

	response.JSON(w, http.StatusOK, response.Success(dto.NewCalculation(result)))
}

func (c *Controller) Multiply(w http.ResponseWriter, r *http.Request) {
	var operands dto.Operands

	if err := json.NewDecoder(r.Body).Decode(&operands); err != nil {
		response.JSON(w, http.StatusBadRequest, response.Error(err))
		return
	}

	result := c.calculator.Multiply(
		valueobject.Number(operands.X),
		valueobject.Number(operands.Y),
	)

	response.JSON(w, http.StatusOK, response.Success(dto.NewCalculation(result)))
}

func (c *Controller) Divide(w http.ResponseWriter, r *http.Request) {
	var operands dto.Operands

	if err := json.NewDecoder(r.Body).Decode(&operands); err != nil {
		response.JSON(w, http.StatusBadRequest, response.Error(err))
		return
	}

	result := c.calculator.Divide(
		valueobject.Number(operands.X),
		valueobject.Number(operands.Y),
	)

	response.JSON(w, http.StatusOK, response.Success(dto.NewCalculation(result)))
}
