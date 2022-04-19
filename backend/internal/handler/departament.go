package handler

import (
	"net/http"

	"bimbo/internal/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateDepartament(c *gin.Context) {
	var (
		dep model.Departament
		err error
	)

	err = c.ShouldBindJSON(&dep)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "Input error", nil)
		return
	}
	err = h.services.DepartamentServiceInterface.Create(dep)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success create Departament", "OK", nil)
}

func (h *Handler) GetListDepartament(c *gin.Context) {
	var (
		result []model.Departament
		err    error
	)

	result, err = h.services.DepartamentServiceInterface.GetList()
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success return list Departament", "OK", result)
}
