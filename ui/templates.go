// Copyright 2016 The Solvist Author(s). All rights reserved.

package ui

import (
	"html/template"
)

// TplCreateNewFormValues passes the new form values.
type TplCreateNewFormValues struct{}

// TplCreateNewForm generates new form for user.
var TplCreateNewForm = template.Must(template.ParseFiles("layout.gohtml", "createnewformpage.gohtml"))
