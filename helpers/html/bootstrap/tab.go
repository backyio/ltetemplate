package tables

import (
	"fmt"
	"github.com/gobuffalo/helpers/hctx"
	"github.com/gobuffalo/tags/v3"
	"html/template"
	"lte/helpers/html/utils"
)

const (
	cTabFrame = `<ul class="{{class}}">{{tabs}}</ul>`
	cTabItem  = `<li class="{{class}}"><a class="{{title_class}}" data-toggle="{{toggle}}" href="#{{id}}">{{title}}</a></li>`
	cTabContentFrame = `<div class="{{class}}">{{tabs}}</div>`
	cTabContentItem = `<div class="{{class}}" id="{{id}}">{{content}}</div>`
)

type (
	//TabImp represents one page on tab
	TabImp struct {
		tags.Options
		Active string
		Disabled string
		ID string
		Title string
		Body string
		Header string
	}

	// TabsImp represents BS tabs
	TabsImp struct {
		// BuffaloHelper is interface to helper function
		utils.BuffaloHelper
		// Options is represents initialization paramteres
		tags.Options
		tabs []*TabImp
	}

	// TabHelper contains pointer to tab implementation
	TabHelper struct {
		// TabsImp is implement core logic
		*TabsImp
	}
)

// newTab it returns TabHelper
func newTabs(opts tags.Options) TabHelper {
	return TabHelper {
		TabsImp: &TabsImp{
			Options: opts,
		},
	}
}

// NewTab implements bootstrap tabs
func NewTabs(opts tags.Options, help hctx.HelperContext) (template.HTML, error) {
	if _, ok := opts["var"]; !ok {
		opts["var"] = "t"
	}
	return utils.BuffaloHelperCallback(nil, nil, opts, help, func(model interface{}, data interface{}, opts tags.Options) utils.BuffaloHelper {
		return newTabs(opts)
	})
}

// BSTabPageContent
func BSTabPage(v string, opts tags.Options, help hctx.HelperContext) (template.HTML, error) {
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
	if t, ok := o.(TabHelper); ok {
		if t.TabsImp == nil {
			return "", utils.CErrorEmptyInstanceVariable
		}
		page, pErr := t.page(opts, id, title)
		if pErr != nil {
			return template.HTML(""), pErr
		}
		// cTabItem  = `<li class="{{class}}"><a class="{{title_class}}" href="#{{id}}">{{title}}</a></li>`
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
		h := utils.Transform(cTabItem, OptionsTabItem, func(name string, value string) string {
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

// BSTabPageContent
func BSTabPageContent(v string, opts tags.Options, help hctx.HelperContext) (h template.HTML, e error) {
	if len(v) == 0 {
		return template.HTML(""), utils.CErrorNoVariableParameter
	}
	id, err := utils.GetValueAsString("id", "", opts)
	if err != nil {
		e = err
		return
	}
	o := help.Value(v)
	if t, ok := o.(TabHelper); ok {
		if t.TabsImp == nil {
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


// pageContent is render tab page body
func (b *TabsImp) pageContent(id string , opts tags.Options, c string) (e error) {
	for _, v := range b.tabs {
		if v.ID == id {
			optionsContentItem := map[string]string {
				"class": `tab-pane container {{class+}} {{fade}} {{active}} {{disabled}}`,
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

			v.Body = utils.Transform(cTabContentItem, optionsContentItem, func(name string, value string) string {
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
func (b *TabsImp) page(opts tags.Options, id string, title string) (*TabImp, error) {
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

	t := &TabImp{
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

func (b *TabsImp) build() string {
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
		"class": "nav nav-tabs {{class+}}",
		"class+": "",
		"tabs":   tabHeader,
	}
	utils.LoadOptions(b.Options, OptionsTabFrame)
	tabHeader = utils.Transform(cTabFrame, OptionsTabFrame, nil)
	return tabHeader + tabBody
}

// String convert struct to string with builder function
func (b *TabsImp) String() string {
	return b.build()
}

// HTML converts
func (s *TabsImp) HTML() template.HTML {
	return template.HTML(s.String())
}

// Append is add rendered body string
func (b *TabsImp) Append(s string) {
}