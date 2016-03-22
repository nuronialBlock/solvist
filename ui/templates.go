package ui

import (
	"html/template"
)

type TplCreateNewFormValues struct{}

var TplCreateNewForm = template.Must(template.ParseFiles("layout.gohtml", "createnewformpage.gohtml"))
