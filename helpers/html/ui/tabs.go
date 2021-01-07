package ui

import (
	"fmt"
	"github.com/gobuffalo/helpers/hctx"
	"github.com/gobuffalo/tags/v3"
	"html/template"
	"lte/helpers/html/utils"
)

type (
	// CardTabImp represents one page on tab
	CardTabImp struct {
		tags.Options
		Active string
		Disabled string
		ID string
		Title string
		Body string
		Header string
	}

	// CardTabsImp represents BS tabs
	CardTabsImp struct {
		// BuffaloHelper is interface to helper function
		utils.BuffaloHelper
		// Options is represents initialization paramteres
		tags.Options
		tabs []*CardTabImp
	}

	// CardTabHelper contains pointer to tab implementation
	CardTabHelper struct {
		// CardTabsImp is implement core logic
		*CardTabsImp
	}
)

// newTab it returns CardTabHelper
func newCardTabs(opts tags.Options) CardTabHelper {
	return CardTabHelper{
		CardTabsImp: &CardTabsImp {
			Options: opts,
		},
	}
}

// NewCardTabs implements admin tle card tabs
func NewCardTabs(opts tags.Options, help hctx.HelperContext) (template.HTML, error) {
	if _, ok := opts["var"]; !ok {
		opts["var"] = "t"
	}
	return utils.BuffaloHelperCallback(nil, nil, opts, help, func(model interface{}, data interface{}, opts tags.Options) utils.BuffaloHelper {
		return newCardTabs(opts)
	})
}

const (
	// cCardTabItem represents tab item html source
	cCardTabItem = `<li class="{{class}}"><a class="{{title_class}}" data-toggle="{{toggle}}" href="#{{id}}">{{title}}</a></li>`
)

// UICardTabPage
func UICardTabPage(v string, opts tags.Options, help hctx.HelperContext) (template.HTML, error) {
	if len(v) == 0 {
		return template.HTML(""), utils.CErrorNoVariableParameter
	}
	id, idErr := utils.GetValueAsString("id", "", opts)
	if idErr != nil {
		return template.HTML(""), idErr
	}
	title, tErr := utils.GetValueAsString("title", "", opts)
	if tErr != nil {
		return template.HTML(""), tErr
	}
	if len(title) == 0 && help.HasBlock(){
		bS, err := help.Block()
		if err != nil {
			return template.HTML(""), err
		}
		title = bS
	}
	if len(title) == 0 {
		return template.HTML(""), fmt.Errorf("Missing id field or title")
	}
	o := help.Value(v)
	if t, ok := o.(CardTabHelper); ok {
		if t.CardTabsImp == nil {
			return "", utils.CErrorEmptyInstanceVariable
		}
		page, pErr := t.page(opts, id, title)
		if pErr != nil {
			return template.HTML(""), pErr
		}

		OptionsTabItem := map[string]string {
			"class": `nav-item {{class+}}`,
			"class+": "",
			"id": page.ID,
			"title": page.Title,
			"title_class": `nav-link {{active}} {{disabled}}`,
			"toggle": "tab",
			"active": page.Active,
			"disabled": page.Disabled,
		}
		h := utils.Transform(cCardTabItem, OptionsTabItem, func(name string, value string) string {
			switch name {
			case "active":
				if value == "true" {
					return "active"
				}
				return ""
			case "disabled":
				if value == "true" {
					return "disabled"
				}
				return ""
			}
			return value
		})
		page.Header = h
	}
	return "", nil
}

// UICardTabPageContent
func UICardTabPageContent(v string, opts tags.Options, help hctx.HelperContext) (h template.HTML, e error) {
	if len(v) == 0 {
		return template.HTML(""), utils.CErrorNoVariableParameter
	}
	id, err := utils.GetValueAsString("id", "", opts)
	if err != nil {
		e = err
		return
	}
	o := help.Value(v)
	if t, ok := o.(CardTabHelper); ok {
		if t.CardTabsImp == nil {
			return "", nil
		}
		var s string
		if help.HasBlock() {
			s, e = help.Block()
			if e != nil {
				return template.HTML(""), e
			}
		}
		e = t.pageContent(id, opts, s)
	}
	return
}

const (
	// cCardTabContentItem represents html source of tab content item
	cCardTabContentItem = `<div class="{{class}}" id="{{id}}">{{content}}</div>`
)
// pageContent is render tab page body
func (b *CardTabsImp) pageContent(id string , opts tags.Options, c string) (e error) {
	for _, v := range b.tabs {
		if v.ID == id {
			optionsContentItem := map[string]string {
				"class": `tab-pane {{class+}} {{fade}} {{active}} {{disabled}}`,
				"class+": "",
				"fade": "fade",
				"id": v.ID,
				"active": v.Active,
				"disabled": v.Disabled,
				"content": c,
			}
			utils.LoadOptions(opts, optionsContentItem)
			optionsContentItem["disabled"] = v.Disabled
			optionsContentItem["active"] = v.Active

			v.Body = utils.Transform(cCardTabContentItem, optionsContentItem, func(name string, value string) string {
				switch name {
				case "active":
					if value == "true" {
						return "active show"
					}
					return ""
				case "disabled":
					if value == "true" {
						return "disabled"
					}
					return ""
				}
				return value
			})
			return
		}
	}
	return utils.CErrorIDNotExists
}

// page is adding new tab page to tabs
func (b *CardTabsImp) page(opts tags.Options, id string, title string) (*CardTabImp, error) {
	sActive := "false"
	sDisabled := "false"

	if _, ok := opts["active"]; ok {
		if b, okt := opts["active"].(bool); okt {
			sActive = fmt.Sprintf("%t", b)
		}
	}

	if _, ok := opts["disabled"]; ok {
		if b, okt := opts["disabled"].(bool); okt {
			sDisabled = fmt.Sprintf("%t", b)
		}
	}

	t := &CardTabImp{
		Options: opts,
		Active: sActive,
		Disabled: sDisabled,
		ID: id,
		Title: title,
	}
	for _, v := range b.tabs {
		if v.ID == id {
			return nil, utils.CErrorIDAlreadyExists
		}
	}
	b.tabs = append(b.tabs, t)
	return t, nil
}

const (
	// cCarTabHeaderFrame is html source of tab frame
	cCarTabHeaderFrame = `<div class="card"><div class="{{card_class}}"><h3 class="card-title p-3">{{title}}</h3><ul class="{{class}}">{{tabs}}</ul></div>`

	// cTabContentFrame is html source of tab
	cTabContentFrame = `<div class="card-body"><div class="{{class}}">{{tabs}}</div></div>`
)
func (b *CardTabsImp) build() string {
	tabHeader := ""
	tabBody := ""

	for _, v := range b.tabs {
		tabHeader += v.Header
		tabBody += v.Body
	}

	OptionsTabContentFrame := map[string]string {
		"class": "tab-content {{class+}}",
		"class+": "",
		"tabs": tabBody,
	}
	tabBody = utils.Transform(cTabContentFrame, OptionsTabContentFrame, nil)


	OptionsTabFrame := map[string]string {
		"class": "nav nav-pills {{class+}}",
		"class+": "m1-auto p-1",
		"card_class": "card-header {{cardclass+}}",
		"cardclass+": "d-flex p-0",
		"title": "",
		"tabs":   tabHeader,
	}
	utils.LoadOptions(b.Options, OptionsTabFrame)
	tabHeader = utils.Transform(cCarTabHeaderFrame, OptionsTabFrame, nil)
	return tabHeader + tabBody
}

// String convert struct to string with builder function
func (b *CardTabsImp) String() string {
	return b.build()
}

// HTML converts
func (s *CardTabsImp) HTML() template.HTML {
	return template.HTML(s.String())
}

// Append is add rendered body string
func (b *CardTabsImp) Append(s string) {
}