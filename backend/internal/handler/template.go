package handler

import (
	"net/http"

	"bimbo/internal/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTemplateCategory(c *gin.Context) {
	var (
		dep model.Template
		err error
	)

	err = c.ShouldBindJSON(&dep)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "Input error", nil)
		return
	}
	err = h.services.TemplateServiceInterface.Create(dep)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success create Template", "OK", nil)
}

func (h *Handler) GetListTemplateCategory(c *gin.Context) {
	var (
		result []model.Template
		err    error
	)

	result, err = h.services.TemplateServiceInterface.GetList()
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success return list Template", "OK", result)
}
