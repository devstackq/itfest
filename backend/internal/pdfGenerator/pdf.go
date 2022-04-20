package pdfgenerator

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"

	"bimbo/internal/model"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/jung-kurt/gofpdf"
)

func GeneratePdf(data model.Templ1) {
	// libs -> config/ tag ?
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)
	// pdf.SetCellMargin(20)

	v := reflect.ValueOf(data)

	values := make([]interface{}, v.NumField())

	for i := 1; i < v.NumField()-1; i++ {
		// pdf.Cell(10, float64(i)*10, v.Field(i).Interface())
		values[i] = v.Field(i).Interface()
	}

	for i, v := range values {
		switch s := v.(type) {
		case model.Company:
			pdf.Cell(1, float64(i+1)*10, s.Name)
		case model.User:
			pdf.Cell(1, float64(i+1)*10, s.FullName)
		case model.Departament:
			pdf.Cell(1, float64(i+1)*10, s.Name)
		case model.Status:
			pdf.Cell(1, float64(i+1)*10, s.Name)
		case string:
			pdf.Cell(1, float64(i+1)*10, s)
		}
	}

	// pdf.Cell(10, 50, "Hello, 2222")

	err := pdf.OutputFileAndClose("hello.pdf")
	log.Println(err, 32)
}

// pdf requestpdf struct
type RequestPdf struct {
	body string
}

// new request to pdf function
func NewRequestPdf(body string) *RequestPdf {
	return &RequestPdf{
		body: body,
	}
}

// parsing template function
func (r *RequestPdf) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}

// generate pdf function
func (r *RequestPdf) GeneratePDF(pdfPath string) (bool, error) {
	t := time.Now().Unix()
	// write whole the body

	if _, err := os.Stat("cloneTemplate/"); os.IsNotExist(err) {
		errDir := os.Mkdir("cloneTemplate/", 0o777)
		if errDir != nil {
			return false, errDir
		}
	}
	err1 := ioutil.WriteFile("cloneTemplate/"+strconv.FormatInt(int64(t), 10)+".html", []byte(r.body), 0o644)
	if err1 != nil {
		return false, err1
	}

	f, err := os.Open("cloneTemplate/" + strconv.FormatInt(int64(t), 10) + ".html")
	if f != nil {
		defer f.Close()
	}
	if err != nil {
		return false, err
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		// log.Fatal(err, 123)
		return false, err
	}
	// /usr/local/go/src

	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		return false, err
	}

	err = pdfg.WriteFile(pdfPath)
	if err != nil {
		return false, err
	}

	dir, err := os.Getwd()
	if err != nil {
		return false, err
	}

	defer os.RemoveAll(dir + "/cloneTemplate")

	return true, nil
}
