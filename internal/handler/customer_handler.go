package handler

import (
	"github.com/AbhinanKumar/smart-dispatch/internal/dto"
	"github.com/AbhinanKumar/smart-dispatch/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CustomerHandler struct {
	custService service.CustomerService
}

func NewCustomerHandler(service service.CustomerService) *CustomerHandler {
	return &CustomerHandler{
		custService: service,
	}
}

func (h *CustomerHandler) Create(c *gin.Context) {
	var req dto.CreateCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := h.custService.Create(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errror": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h *CustomerHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid customer id",
		})
		return
	}

	res, err := h.custService.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *CustomerHandler) GetAll(c *gin.Context) {
	res, err := h.custService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch customers",
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *CustomerHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid customer id",
		})
		return
	}

	var req dto.CreateCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.custService.Update(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
