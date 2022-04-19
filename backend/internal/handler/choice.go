package handler

import (
	"net/http"
	"strconv"

	"bimbo/internal/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTemplateChoice(c *gin.Context) {
	var (
		ch     []model.Choice
		err    error
		tmplId int
	)

	tmplId, err = strconv.Atoi(c.Param("tmplId"))
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "input error", nil)
		return
	}

	err = c.ShouldBindJSON(&ch)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "Input error", nil)
		return
	}
	err = h.services.ChoiceServiceInterface.Create(ch, tmplId)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success create Choice", "OK", nil)
}

func (h *Handler) GetListTemplateChoiceBiID(c *gin.Context) {
	var (
		result []model.Choice
		err    error
		tmplId int
	)
	tmplId, err = strconv.Atoi(c.Param("tmplId"))
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "input error", nil)
		return
	}

	result, err = h.services.ChoiceServiceInterface.GetList(tmplId)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	responseWithStatus(c, http.StatusOK, "success return list Choice", "OK", result)
}
