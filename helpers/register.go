package helpers

import (
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/helpers"
	"github.com/gobuffalo/helpers/forms"
	"github.com/gobuffalo/helpers/hctx"
	HtmlTables "lte/helpers/html/tables"
	"lte/helpers/html/ui"
	"lte/helpers/html/widgets"
	JSTables "lte/helpers/js/tables"
	"github.com/gobuffalo/helpers/forms/bootstrap"
)

// hWidgets contains lte widget helpers
var (
	hWidgets = widgets.New()
	hHtmlTables  = HtmlTables.New()
	hUI      = ui.New()
	hJSTables = JSTables.New()
)

// lteALL is function for collecting all available lte helper function
var lteALL = func() hctx.Map {
	return hctx.Merge(
		hWidgets,
		hHtmlTables,
		hUI,
		hJSTables,
	)
}

// DefaultHelpers returns merge LTE and default helpers structure
func DefaultHelpers() render.Helpers {
	hAll := helpers.ALL()
	lAll := lteALL();
	allHelpers := hctx.Merge( hAll, lAll)
	h := render.Helpers(allHelpers)
	h[forms.FormKey] = bootstrap.Form
	h[forms.FormForKey] = bootstrap.FormFor
	h["form_for"] = bootstrap.FormFor
	return h
}