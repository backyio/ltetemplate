package tables

import (
	"github.com/gobuffalo/helpers/hctx"
	"github.com/gobuffalo/tags/v3"
	"html/template"
	"lte/helpers/html/utils"
)

const (
	CDataTable    = 0
	CDateTableFor = 1
)

type (
	JsDataTableImp struct {
		tags.Options
		utils.BuffaloHelper
		Model interface{}
		Data  interface{}
		ID    string
		Title string
		Type  int
	}

	JsDataTableHelper struct {
		*JsDataTableImp
	}
)

func (s *JsDataTableImp) Append(body string) {
}

func (s *JsDataTableImp) HTML() template.HTML {
	return template.HTML("")
}

// NewDataTable create new data table for simple row based rendering
func NewDataTable(opts tags.Options, help hctx.HelperContext) (template.HTML, error) {
	if _, ok := opts["var"]; !ok {
		opts["var"] = "dt"
	}
	return utils.BuffaloHelperCallback(nil, nil, opts, help, func(model interface{}, data interface{}, opts tags.Options) utils.BuffaloHelper {
		return newDataTable(opts)
	})
}

// NewDataTable implements data table render
func NewDataTableFor(model interface{}, data interface{}, opts tags.Options, help hctx.HelperContext) (template.HTML, error) {
	if _, ok := opts["var"]; !ok {
		opts["var"] = "dt"
	}
	return utils.BuffaloHelperCallback(model, data, opts, help, func(model interface{}, data interface{}, opts tags.Options) utils.BuffaloHelper {
		return newDataTable(opts)
	})
}

// newDataTable is create new data table
func newDataTable(opts tags.Options) utils.BuffaloHelper {
	return JsDataTableHelper{
		JsDataTableImp: &JsDataTableImp{
			Options: opts,
			Type:    CDataTable,
			Model:   nil,
			Data:    nil,
		},
	}
}

// newDataTable is create new data table
func newDataTableFor(model interface{}, data interface{}, opts tags.Options) utils.BuffaloHelper {
	return JsDataTableHelper{
		JsDataTableImp: &JsDataTableImp{
			Model:   model,
			Data:    data,
			Type:    CDateTableFor,
			Options: opts,
		},
	}
}
