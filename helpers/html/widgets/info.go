package widgets

import (
	"github.com/gobuffalo/helpers/hctx"
	"github.com/gobuffalo/tags/v3"
	"html/template"
)

// Info implements admin lte info box
func Info(opts tags.Options, help hctx.HelperContext) (template.HTML, error) {
	if opts == nil {
		opts = tags.Options{}
	}
	hn := "f"
	if n, ok := opts["var"]; ok {
		hn = n.(string)
		delete(opts, "var")
	}
	help.Set(hn, nil)
	s, err := help.Block()
	if err != nil {
		return "", err
	}
	return template.HTML(s), nil
}