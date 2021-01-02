package tables

import (
	"github.com/gobuffalo/helpers/hctx"
	"github.com/gobuffalo/tags/v3"
	"html/template"
)

// htmler generates HTML source
type htmler interface {
	HTML() template.HTML
}

type tableHelper interface {
	Append(...tags.Body)
	htmler
}

func mainTableHelper(data interface{}, opts tags.Options, help hctx.HelperContext, fn func(data interface{}, opts tags.Options) tableHelper) (template.HTML, error) {
	if opts == nil {
		opts = tags.Options{}
	}
	hn := "f"
	if n, ok := opts["var"]; ok {
		hn = n.(string)
		delete(opts, "var")
	}
	table := fn(data, opts)
	help.Set(hn, table)
	s, err := help.Block()
	if err != nil {
		return "", err
	}
	table.Append(s)
	return table.HTML(), nil
}