package widgets

import (
	"github.com/gobuffalo/helpers/hctx"
	"github.com/gobuffalo/tags/v3"
	"html/template"
)

// htmler generates HTML source
type htmler interface {
	HTML() template.HTML
}

type widgetHelper interface {
	Append(...tags.Body)
	htmler
}

func mainWidgetHelper(opts tags.Options, help hctx.HelperContext, fn func(opts tags.Options) widgetHelper) (template.HTML, error) {
	if opts == nil {
		opts = tags.Options{}
	}
	hn := "f"
	if n, ok := opts["var"]; ok {
		hn = n.(string)
		delete(opts, "var")
	}
	widget := fn(opts)
	help.Set(hn, widget)
	s, err := help.Block()
	if err != nil {
		return "", err
	}
	widget.Append(s)
	return widget.HTML(), nil
}
