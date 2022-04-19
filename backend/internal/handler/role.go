package handler

import (
	"net/http"

	"bimbo/internal/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateRole(c *gin.Context) {
	var (
		role model.Role
		err  error
	)

	err = c.ShouldBindJSON(&role)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "Input error", nil)
		return
	}
	err = h.services.RoleServiceInterface.Create(role)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success create Departament", "OK", nil)
}

func (h *Handler) GetListRole(c *gin.Context) {
	var (
		result []model.Role
		err    error
	)

	result, err = h.services.RoleServiceInterface.GetList()
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success return list Departament", "OK", result)
}
