package helpers

import (
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/helpers"
	"github.com/gobuffalo/helpers/forms"
	"github.com/gobuffalo/helpers/forms/bootstrap"
	"github.com/gobuffalo/helpers/hctx"
	bs "lte/helpers/html/bootstrap"
	HtmlTables "lte/helpers/html/tables"
	"lte/helpers/html/ui"
	"lte/helpers/html/widgets"
	JSTables "lte/helpers/js/tables"
)

// hWidgets contains lte widget helpers
var (
	hWidgets = widgets.New()
	hHtmlTables  = HtmlTables.New()
	hUI      = ui.New()
	hJSTables = JSTables.New()
	hBootStrap = bs.New()
)

// lteALL is function for collecting all available lte helper function
var lteALL = func() hctx.Map {
	return hctx.Merge(
		hWidgets,
		hHtmlTables,
		hUI,
		hJSTables,
		hBootStrap,
	)
}

// DefaultHelpers returns merge LTE and default helpers structure
func DefaultHelpers() render.Helpers {
	hAll := helpers.ALL()
	lAll := lteALL();
	allHelpers := hctx.Merge( hAll, lAll)
	allHelpers[forms.FormKey] = bootstrap.Form
	allHelpers[forms.FormForKey] = bootstrap.FormFor
	allHelpers["form_for"] = bootstrap.FormFor
	return render.Helpers(allHelpers)
}