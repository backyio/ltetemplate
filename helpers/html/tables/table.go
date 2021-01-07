package tables

import (
	"github.com/gobuffalo/helpers/hctx"
	"github.com/gobuffalo/tags/v3"
	"html/template"
	"lte/helpers/html/utils"
)

type (

	SimpleTableImp struct {
		utils.BuffaloHelper
		tags.Options
		Data       interface{}
		ID         string
		Name       string
	}

	SimpleTableHelper struct {
		*SimpleTableImp
	}
)

func (s *SimpleTableImp) build() string {
	return ""
}


func (s *SimpleTableImp) Append(body string) {
	s.AddRow(body)
}

func (s *SimpleTableImp) HTML() template.HTML {
	return template.HTML(s.String())
}

func (s *SimpleTableImp) AddRow(body string) {

}

func (s *SimpleTableImp) String() string {
	return s.build()
}

// NewSimpleTable implements data table renderer
func NewSimpleTable(opts tags.Options, help hctx.HelperContext) (template.HTML, error) {
	if _, ok := opts["var"]; !ok {
		opts["var"] = "st"
	}
	return utils.BuffaloHelperCallback(nil, nil, opts, help, func(model interface{}, data interface{}, opts tags.Options) utils.BuffaloHelper {
		return newSimpleTable(data, opts)
	})
}

// NewSimpleTableFor implements data table renderer
func NewSimpleTableFor(data interface{}, opts tags.Options, help hctx.HelperContext) (template.HTML, error) {
	if _, ok := opts["var"]; !ok {
		opts["var"] = "st"
	}
	return utils.BuffaloHelperCallback(nil, nil, opts, help, func(model interface{}, data interface{}, opts tags.Options) utils.BuffaloHelper {
		return newSimpleTable(data, opts)
	})
}

// newSimpleTable create simple table
func newSimpleTable(data interface{}, opts tags.Options) utils.BuffaloHelper {
	return SimpleTableHelper {
		SimpleTableImp: &SimpleTableImp {
			Options: opts,
			Data: data,
		},
	}
}