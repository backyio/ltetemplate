package ui

import "github.com/gobuffalo/helpers/hctx"

// New returns a map of theui helpers within this package.
func New() hctx.Map {
	return hctx.Map{
		"UITabs": NewUITabs,
	}
}
