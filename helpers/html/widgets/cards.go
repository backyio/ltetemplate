package widgets

import (
	"fmt"
	"github.com/gobuffalo/helpers/hctx"
	"github.com/gobuffalo/tags/v3"
	"html/template"
	"lte/helpers/html/utils"
	"strings"
)

/*
 Card implements admin lte card box
 Example : Card( { bg: "info", shadow: "small", title: "test" } )
 Options :
	bg          - info default, success, warning, danger
	shadow      - none, small default, regular, large
	title       -
	description -
	refresh_url - 'refresh maximize collapse remove'
    btn         -
*/
func Card(opts tags.Options, help hctx.HelperContext) (template.HTML, error) {

	options := map[string]string{
		"bg":          "info",
		"shadow":      "small",
		"title":       "",
		"description": "",
		"btn":         "refresh maximize collapse remove",
		"refresh_url":     "#",
	}

	var strData string
	btnData := map[string]string {
		"refresh": `<button type="button" class="btn btn-tool" data-card-widget="card-refresh" data-source="{{refresh_url}}" data-source-selector="#card-refresh-content" data-load-on-init="false"><i class="fas fa-sync-alt"></i></button>`,
		"maximize": `<button type="button" class="btn btn-tool" data-card-widget="maximize"><i class="fas fa-expand"></i></button>`,
		"collapse": `<button type="button" class="btn btn-tool" data-card-widget="collapse"><i class="fas fa-minus"></i></button>`,
		"remove": `<button type="button" class="btn btn-tool" data-card-widget="remove"><i class="fas fa-times"></i></button>`,
	}
	_ = utils.LoadOptions(opts, options)
	utils.LoadFromBlock("description", help, options)

	btnOpts := options["btn"]
	for k, v := range btnData {
		if strings.Contains(btnOpts, k) {
			strData = fmt.Sprintf("%s%s", strData, v)
		}
	}

body := `<div class="card card-success">
<div class="card-header">
 <h3 class="card-title">{{title}}</h3>
 <div class="card-tools">%s</div>
</div>
<div class="card-body">{{description}}</div>
<div class="overlay"><i class="fas fa-2x fa-sync-alt fa-spin"></i></div></div>`

	body = fmt.Sprintf(body, strData)

	fullHtml := utils.FormatMap(body, options, func(name string, value string) string {
		switch name {
		case "shadow":
			return utils.ConvShadow(value)
		case "bg":
			return utils.ConvBG("card", value)
		}
		return value
	})
	return template.HTML(fullHtml), nil
}
