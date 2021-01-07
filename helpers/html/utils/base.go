package utils

import (
	"github.com/gobuffalo/helpers/hctx"
	"github.com/gobuffalo/tags/v3"
	"html/template"
)

type BuffaloHelper interface {
	Append(body string)
	HTML() template.HTML
}

// BuffaloHelperCallback is the callback function for object generator
func BuffaloHelperCallback(model interface{}, data interface{}, opts tags.Options, help hctx.HelperContext, fn func(model interface{}, data interface{}, opts tags.Options) BuffaloHelper) (template.HTML, error) {
	if opts == nil {
		opts = tags.Options{}
	}
	hn := "f"
	if n, ok := opts["var"]; ok {
		hn = n.(string)
		delete(opts, "var")
	}
	ui := fn(model, data, opts)
	help.Set(hn, ui)
	if help.HasBlock() {
		s, err := help.Block()
		if err != nil {
			return "", err
		}
		ui.Append(s)
	}
	return ui.HTML(), nil
}
