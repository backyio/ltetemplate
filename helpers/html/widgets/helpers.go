package widgets

import "github.com/gobuffalo/helpers/hctx"

// New returns a map of the widget helpers within this package.
func New() hctx.Map {
	return hctx.Map{
		"Info": Info,
	}
}
