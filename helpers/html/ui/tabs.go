package ui

import (
	"github.com/gobuffalo/helpers/hctx"
	"github.com/gobuffalo/tags/v3"
	"html/template"
	"lte/helpers/html/utils"
)

type (
	uiTab struct {
		ID      string
		Title   string
		Content string
	}
	uiTabsImp struct {
		tags.Options
		utils.BuffaloHelper
		ID    string
		Title string
	}

	uiTabsHelper struct {
		*uiTabsImp
	}
)

// NewUITabs ui tabs with options
func NewUITabs(opts tags.Options, help hctx.HelperContext) (template.HTML, error) {
	return utils.BuffaloHelperCallback(nil, nil, opts, help, func(model interface{}, data interface{}, opts tags.Options) utils.BuffaloHelper {
		return newUITabs(opts)
	})
}

// newUITabs is create new tabs collection
func newUITabs(opts tags.Options) utils.BuffaloHelper {
	return uiTabsHelper{
		uiTabsImp: &uiTabsImp{
			Options: opts,
		},
	}
}
