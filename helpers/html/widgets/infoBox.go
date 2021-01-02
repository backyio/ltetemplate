package widgets

import (
	"github.com/gobuffalo/helpers/hctx"
	"github.com/gobuffalo/tags/v3"
	"html/template"
	"lte/helpers/html/utils"
)

/*
 InfoBox implements admin lte info box
 Example : widget.InfoBox( { icon: "fa-calendar-alt" } )
 Options examples :
	bg      - info default, success, warning, danger
	icon    - fa-calendar-alt
	shadow  - none, small default, regular, large
	ct      - text, number
	title   -
	block   - additional content
	description -
*/
func InfoBox(opts tags.Options, help hctx.HelperContext) (template.HTML, error) {

	options := map[string]string{
		"bg":          "info",
		"shadow":      "small",
		"icon":        "",
		"ct":          "number",
		"description": "",
		"title":       "",
	}

	_ = utils.LoadOptions(opts, options)
	utils.LoadFromBlock("description", help, options)
	body := `<div class="info-box {{shadow}}">
				 <span class="info-box-icon {{bg}}">
					<i class="far {{icon}}"></i>
                 </span>
				<div class="info-box-content">
                 <span class="info-box-text">{{title}}</span>
                 <span class="{{ct}}">{{description}}</span>
               </div>
            </div>`

	fullHtml := utils.FormatMap(body, options, func(name string, value string) string {
		switch name {
		case "shadow":
			return utils.ConvShadow(value)
		case "bg":
			return utils.ConvBG(value)
		case "ct":
			return utils.ConvInfoContentType(value)
		}
		return value
	})
	return template.HTML(fullHtml), nil
}
