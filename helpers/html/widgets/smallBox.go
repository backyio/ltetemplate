package widgets

import (
	"github.com/gobuffalo/helpers/hctx"
	"github.com/gobuffalo/tags/v3"
	"html/template"
	"lte/helpers/html/utils"
)

const (
	// CSmallBoxBody contains html body for small box
	CSmallBoxBody = `<div class="small-box {{bg}} {{shadow}}">
              <div class="inner">
                <h3>{{title}}<sup style="font-size: {{font_size}}">%</sup></h3>
                <p>{{description}}</p>
              </div>
              <div class="icon">
                <i class="ion {{icon}}"></i>
              </div>
              <a href="{{btn_url}}" class="small-box-footer">
                {{btn_title}} <i class="fas {{btn_icon}}"></i>
              </a>
            </div>`
)

/*
 SmallBox implements admin lte small box with link
 Example : SmallBox( { icon: "fa-calendar-alt" } )
 Options :
	bg          - info default, success, warning, danger
	shadow      - none, small default, regular, large
	icon        - fa-arrow-circle-right
	title       -
	description -
	font_size   -
	btn_title   -
	btn_icon    -
	btn_url     -
*/
func SmallBox(opts tags.Options, help hctx.HelperContext) (template.HTML, error) {

	options := map[string]string{
		"bg":          "info",
		"shadow":      "small",
		"icon":        "",
		"description": "",
		"title":       "",
		"font_size":    "20px",
		"btn_title":   "More info",
		"btn_icon":    "fa-arrow-circle-right",
		"btn_url":     "#",
	}

	utils.LoadOptions(opts, options)
	utils.LoadFromBlock("description", help, options)

	fullHtml := utils.Transform(CSmallBoxBody, options, func(name string, value string) string {
		switch name {
		case "shadow":
			return utils.ConvShadow(value)
		case "bg":
			return utils.ConvBG("bg", value)
		}
		return value
	})
	return template.HTML(fullHtml), nil
}