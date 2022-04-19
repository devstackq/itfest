package handler

import (
	"net/http"

	"bimbo/internal/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePosition(c *gin.Context) {
	var (
		pos model.Position
		err error
	)

	err = c.ShouldBindJSON(&pos)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "Input error", nil)
		return
	}
	err = h.services.PositionServiceInterface.Create(pos)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success create Departament", "OK", nil)
}

func (h *Handler) GetListPosition(c *gin.Context) {
	var (
		result []model.Position
		err    error
	)

	result, err = h.services.PositionServiceInterface.GetList()
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success return list Departament", "OK", result)
}
