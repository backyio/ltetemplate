package tables

import (
	"github.com/gobuffalo/helpers/hctx"
	"github.com/gobuffalo/tags/v3"
	"html/template"
	"lte/helpers/html/utils"
)

type (
	// ButtonGroupImp represents BS button group
	ButtonGroupImp struct {
		// BuffaloHelper is interface to helper function
		utils.BuffaloHelper
		// Options is represents initialization paramteres
		tags.Options
		body string
	}

	// ButtonGroupHelper contains pointer to button group implementation
	ButtonGroupHelper struct {
		// ButtonGroupImp is implement core logic
		*ButtonGroupImp
	}
)

// newButtonGroup it returns ButtonGroup
func newButtonGroup(opts tags.Options,) ButtonGroupHelper {
	return ButtonGroupHelper {
		ButtonGroupImp: &ButtonGroupImp {
			Options: opts,
		},
	}
}

// NewButtonGroup implements bootstrap button group
func NewButtonGroup(opts tags.Options, help hctx.HelperContext) (template.HTML, error) {
	if _, ok := opts["var"]; !ok {
		opts["var"] = "b"
	}
	return utils.BuffaloHelperCallback(nil, nil, opts, help, func(model interface{}, data interface{}, opts tags.Options) utils.BuffaloHelper {
		return newButtonGroup(opts)
	})
}

/* build is internal function to build html string
    Options:

*/
func (b *ButtonGroupImp) build() string {
	return ""
}
// String convert struct to string with builder function
func (b *ButtonGroupImp) String() string {
	return b.build()
}

// HTML converts
func (s *ButtonGroupImp) HTML() template.HTML {
	return template.HTML(s.String())
}