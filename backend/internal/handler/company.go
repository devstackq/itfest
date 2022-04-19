package handler

import (
	"net/http"

	"bimbo/internal/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCompany(c *gin.Context) {
	var (
		company model.Company
		err     error
	)

	err = c.ShouldBindJSON(&company)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "Input error", nil)
		return
	}
	err = h.services.CompanyServiceInterface.Create(company)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success create Company", "OK", nil)
}

func (h *Handler) GetListCompany(c *gin.Context) {
	var (
		result []model.Company
		err    error
	)

	result, err = h.services.CompanyServiceInterface.GetList()
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success return list Comapny", "OK", result)
}
