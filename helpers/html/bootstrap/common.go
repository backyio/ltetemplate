package tables

import (
	"github.com/gobuffalo/helpers/hctx"
	"github.com/gobuffalo/tags/v3"
	"html/template"
	"lte/helpers/html/utils"
)

const (
	cAlertBox = `<div style="{{style}}" class="{{class}}">` +
		`<button type="button" class="close" data-dismiss="alert" aria-label="{{close_title}}">` +
		` <span aria-hidden="true" style="{{close_style}}">{{close_sign}}</span></button>{{content}}</div>`
)

/* AlertBox implement Bootstrap alert box
   Options :
	class   - base class options
	style   - base style options
	class+  - additional class options
	style+  - additional style options
	bg      - background color
	content - base message or block
	close_style - base style of close button
	close_title - close button label
    close_sign  - x
*/
func AlertBox(opts tags.Options, help hctx.HelperContext) (template.HTML, error) {
	options := map[string]string{
		"class":       `alert {{bg}} {{fade}} alert-dismissible show {{class+}}`,
		"fade":        "fade in",
		"style":       `margin-top:18px;{{style+}}`,
		"style+":      "",
		"content":     "",
		"class+":      "align-middle",
		"bg":          "warning",
		"close_style": "",
		"close_title": "Close",
		"close_sign":  "x",
	}
	utils.LoadOptions(opts, options)
	utils.LoadFromBlock("content", help, options)

	fullHtml := utils.Transform(cAlertBox, options, func(name string, value string) string {
		switch name {
		case "bg":
			return utils.ConvBG("alert", value)
		case "close_style":
			return utils.ConvStyle(value)

		}
		return value
	})
	return template.HTML(fullHtml), nil
}

/* Button implement Bootstrap button
   Options :
	class   - base class
	class+  - additional class parameters
	style   - style of the button
	bg      - background color
	title   - button title or block
	url     - link url
	size    - size of the button
    disabled - true or false
    block    - block sized, true or falses
*/
const (
	cButton = `<a href="{{url}}" style="{{style}}" class="{{class}}" role="button">{{title}}</a >`
)

func Button(opts tags.Options, help hctx.HelperContext) (template.HTML, error) {
	options := map[string]string{
		"class":    "btn {{bg}} {{size}} {{disabled}} {{block}} {{class+}}",
		"bg":       "info",
		"title":    "",
		"url":      "#",
		"size":     "",
		"disabled": "false",
		"block":    "false",
		"class+":   "",
		"style":    "",
	}
	utils.LoadOptions(opts, options)
	utils.LoadFromBlock("title", help, options)
	html := utils.Transform(cButton, options, func(name string, value string) string {
		switch name {
		case "bg":
			return utils.ConvBG("bg", value)
		case "size":
			return utils.ConvCommon("btn", value)
		case "disabled":
			return utils.ConvBool(value, "disabled", "active")
		case "block":
			return utils.ConvBool(value, "btn-block", "")
		}
		return value
	})
	return template.HTML(html), nil
}

/* ButtonCollapsible implement Bootstrap basic collapsible button
   Options :
	class   - base class
	class+  - additional class parameters
	id      - button custom id
	style   - style of the button
	bg      - background color
	title   - button title or block
	url     - link url
	size    - size of the button
    disabled - true or false
    block    - block sized, true or falses
	expanded - content visible
	content  - button content loaded from block or options
*/
const (
	cButtonCollapsible = `<button  class="{{class}}" style="{{style}}" data-toggle="collapse" data-target="#{{id}}" aria-expanded="{{expanded}}">{{title}}</button>\n<div id="{{id}}" class="collapse">{{content}}</div>\n`
)

func ButtonCollapsible(opts tags.Options, help hctx.HelperContext) (template.HTML, error) {
	options := map[string]string{
		"class":    "btn {{bg}} {{size}} {{disabled}} {{block}} {{class+}}",
		"content":  "",
		"id":       "",
		"bg":       "info",
		"title":    "",
		"url":      "#",
		"size":     "",
		"disabled": "false",
		"block":    "false",
		"class+":   "",
		"style":    "",
	}
	utils.LoadOptions(opts, options)
	utils.LoadFromBlock("content", help, options)
	html := utils.Transform(cButtonCollapsible, options, func(name string, value string) string {
		switch name {
		case "bg":
			return utils.ConvBG("bg", value)
		case "size":
			return utils.ConvCommon("btn", value)
		case "disabled":
			return utils.ConvBool(value, "disabled", "active")
		case "block":
			return utils.ConvBool(value, "btn-block", "")
		}
		return value
	})
	return template.HTML(html), nil
}

/* Jumbotron implement Bootstrap Jumbotron
   Options :
	class   - base class
    class+  - additional class parameters
	title   -
	content -
*/
const (
	cJumbotron = `<div style="{{style}}" class="{{class}}"><h1>{{title}}</h1><p>{{content}}</p></div>`
)

func Jumbotron(opts tags.Options, help hctx.HelperContext) (template.HTML, error) {
	options := map[string]string{
		"title":   "",
		"content": "",
		"class":   "jumbotron {{class+}}",
		"class+":  "",
		"style":   "",
	}
	utils.LoadOptions(opts, options)
	utils.LoadFromBlock("content", help, options)
	return template.HTML(utils.Transform(cJumbotron, options, nil)), nil
}

/* ToolTip implement Bootstrap Jumbotron
   Options :
	class   - additional class options
    style   - additional style options
	title   - tooltip title
	content - tooltip message, coming from options or block
	placement - tooltop message postition ( top, bottom, left, right )
*/
const (
	// cToolTip is represents tooltop html source
	cToolTip = `<a class="{{class}}" style="{{style}}" href="{{url}}" data-toggle="tooltip" data-trigger="{{trigger}}" ` +
		` data-placement="{{placement}}" data-animation="{{animation}}" data-delay="{{delay}}"  data-html="{{html}}" title="{{content}}">{{title}}</a>`
)

func ToolTip(opts tags.Options, help hctx.HelperContext) (template.HTML, error) {
	options := map[string]string{
		"title":     "",
		"content":   "",
		"class":     "",
		"style":     "",
		"url":       "#",
		"placement": "top",
		"animation": "true",
		"delay":     "1",
		"html":      "true",
		"trigger":   "hover",
	}
	utils.LoadOptions(opts, options)
	utils.LoadFromBlock("content", help, options)
	return template.HTML(utils.Transform(cToolTip, options, nil)), nil
}

const (
	// cBadge is represents cBadge html source
	cBadge = `<span style="{{style}}" class="{{class}}">{{title}}</span>`
)

func Badge(opts tags.Options, help hctx.HelperContext) (template.HTML, error) {
	options := map[string]string{
		"title": "",
		"class": "badge {{pill}} badge-{{bg}}",
		"style": "",
		"bg":    "primary",
		"pill":  "false",
	}
	utils.LoadOptions(opts, options)
	utils.LoadFromBlock("title", help, options)
	return template.HTML(utils.Transform(cBadge, options, func(name string, value string) string {
		switch name {
		case "pill":
			if value == "true" {
				value = "badge-pill"
			} else {
				value = ""
			}
		}
		return value
	})), nil
}

const (
	// cBadgeHeading is represents cBadge with heading html source
	cBadgeHeading = `<h{{size}}>{{text}}<span style="{{style}}" class="{{class}}">{{title}}</span></h{{size}}>`
)

// BadgeHeading implements badge with headings
func BadgeHeading(opts tags.Options, help hctx.HelperContext) (template.HTML, error) {
	options := map[string]string{
		"title": "",
		"class": "badge {{pill}} badge-{{bg}}",
		"style": "",
		"bg":    "primary",
		"pill":  "false",
		"size":  "5",
		"text":  "",
	}
	utils.LoadOptions(opts, options)
	utils.LoadFromBlock("title", help, options)
	return template.HTML(utils.Transform(cBadgeHeading, options, func(name string, value string) string {
		switch name {
		case "pill":
			if value == "true" {
				value = "badge-pill"
			} else {
				value = ""
			}
		}
		return value
	})), nil
}
