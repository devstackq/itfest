package handler

import (
	"net/http"

	"bimbo/internal/model"
	pdfgenerator "bimbo/internal/pdfGenerator"

	"github.com/gin-gonic/gin"
)

// 1.4 return client from Db by templID -> json {title: key}
// 1.5 client:  /create pdf doc -> getFields By TempId -> map[jsonKey]title
func (h *Handler) CreateDocument(c *gin.Context) {
	var (
		// doc  model.Document
		err      error
		tmplData model.Templ1
	)

	err = c.ShouldBindJSON(&tmplData)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "Input error", nil)
		return
	}
	pdfgenerator.GeneratePdf(tmplData)

	r := pdfgenerator.NewRequestPdf("")
	// html template path
	templatePath := "./internal/templates/sample.html"

	// path for download pdf
	outputPath := "./fileServer/example.pdf"

	if err := r.ParseTemplate(templatePath, tmplData); err == nil {
		ok, err := r.GeneratePDF(outputPath)
		if !ok {
			responseWithStatus(c, http.StatusInternalServerError, err.Error(), "Error", nil)
			return
		}
	} else {
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "Error", nil)
		return
	}

	// add history
	// add access

	responseWithStatus(c, http.StatusOK, "success create Document", "OK", nil)
}
