package tables

import (
	"github.com/gobuffalo/helpers/hctx"
	"github.com/gobuffalo/tags/v3"
	"html/template"
	"reflect"
)

type (

	SimpleTableImp struct {
		tableHelper
		tags.Options
		data       interface{}
		id         string
		name       string
		reflection reflect.Value
		Errors     tags.Errors
	}

	simpleTableHelper struct {
		*SimpleTableImp
	}
)

// NewSimpleTable implements data table renderer
func NewSimpleTable(opts tags.Options, help hctx.HelperContext) (template.HTML, error) {
	return mainTableHelper(nil, opts, help, func(data interface{}, opts tags.Options) tableHelper {
		return newSimpleTable(data, opts)
	})
}

// NewSimpleTableFor implements data table renderer
func NewSimpleTableFor(data interface{}, opts tags.Options, help hctx.HelperContext) (template.HTML, error) {
	return mainTableHelper(data, opts, help, func(data interface{}, opts tags.Options) tableHelper {
		return newSimpleTable(data, opts)
	})
}

// newSimpleTable create simple table
func newSimpleTable(data interface{}, opts tags.Options) simpleTableHelper {
	return simpleTableHelper {
		SimpleTableImp: &SimpleTableImp {
			Options: opts,
			data: data,
		},
	}
}