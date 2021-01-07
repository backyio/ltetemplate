package actions

import (
    "net/http"
    
	"github.com/gobuffalo/buffalo"
)

// ExamplesBootstrap default implementation.
func ExamplesBootstrap(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("examples/bootstrap.plush.html"))
}


// ExamplesWidgets default implementation.
func ExamplesWidgets(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("examples/widgets.plush.html"))
}


// ExamplesTables default implementation.
func ExamplesTables(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("examples/tables.plush.html"))
}

