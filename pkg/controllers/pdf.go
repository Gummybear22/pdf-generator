package controllers

import (
	"bytes"
	"os"
	"strings"
	"text/template"

	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"

	"github.com/gofiber/fiber/v2"
)

// initialize buffer to store html with data
var buf bytes.Buffer

// integrate to other pdftest function
func HtmlTest(c *fiber.Ctx) error {
	// declare the html template
	var temp *template.Template
	temp = template.Must(template.ParseFiles("pdf-templates/sample1.html"))

	// declare data struct to be rendered in template
	type data struct {
		Pnum string `json:"pnum"`
	}

	//Get data from data source

	// initialize data
	Data := data{
		Pnum: "12345678912345678",
	}

	// execute the template and data, store result in buffer
	err := temp.Execute(&buf, Data)
	if err != nil {
		return err
	}

	return c.SendString(buf.String())
}

func PdfTest(c *fiber.Ctx) error {
	// pdf generator
	r, err := wkhtml.NewPDFGenerator()
	if err != nil {
		return err
	}

	// global options for pdf
	r.PageSize.Set(wkhtml.PageSizeLetter)

	// no need to read file
	// read template html using absolute path
	b, err := os.ReadFile("/Users/g.tan/Projects/Pdf-Generator/Go_Template/pdf-templates/sample.html")
	if err != nil {
		return err
	}

	// use buffer instead
	// convert html to string
	str := string(b)

	// r.AddPage(wkhtml.NewPageReader(strings.NewReader(str)))

	// set page 1
	page1 := wkhtml.NewPageReader(strings.NewReader(str))

	// set page 1 options
	page1.Allow.Set("/Users/g.tan/Projects/Pdf-Generator/Go_Template/pdf-templates")
	page1.EnableLocalFileAccess.Set(true)

	// add page
	r.AddPage(page1)

	// Create PDF document in internal buffer
	err = r.Create()
	if err != nil {
		return err
	}

	//Your Pdf Name
	err = r.WriteFile("./test.pdf")
	if err != nil {
		return err
	}

	return c.SendString("success")
}
