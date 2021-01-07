package widgets

import (
	"fmt"
	"github.com/gobuffalo/helpers/hctx"
	"github.com/gobuffalo/tags/v3"
	"html/template"
	"lte/helpers/html/utils"
)

var (
	CSocial = `<div class="card card-widget widget-user-{{size}}">
              <div class="widget-user-header {{bg}} {{shadow}}">
                <div class="widget-user-image">
                  <img class="img-{{mode}} elevation-2" src="{{image}}" alt="{{alt}}">
                </div>
                <h3 class="widget-user-username">{{name}}</h3>
                <h5 class="widget-user-desc">{{description}}</h5>
              </div>
              <div class="card-footer p-0">
                <ul class="nav flex-column">
                  {{rows}}
                 </ul>
              </div>
            </div>`

	CSocialRow             = `<li class="nav-item">{{url}}{{title}}<span {{class}}>{{content}}</span></a></li>`
	CSocialRowUrl          = `<a href="{{url}}" class="nav-link">`
	CSocialRowDefaultClass = `class="float-right badge {{bg}}"`
)

type (
	// SocialBoxImp represents social grid
	SocialBoxImp struct {
		// BuffaloHelper is interface to helper function
		utils.BuffaloHelper
		// Options is represents initialization paramteres
		tags.Options
		// Rows is represents table rows
		Rows string
	}

	// SocialBoxHelper contains pointer to social box implementation
	SocialBoxHelper struct {
		// SocialBoxImp is implement social box
		*SocialBoxImp
	}
)

// Row is add new row to social box
func (s *SocialBoxImp) Row(opts tags.Options, help hctx.HelperContext) {
	options := map[string]string{
		"url":     "#",
		"title":   "",
		"content": "",
		"bg":      "info",
		"class":   "",
	}
	utils.LoadOptions(opts, options)
	class := options["class"]
	if len(class) == 0 {
		class = CSocialRowDefaultClass
	}
	// custom class
	row := utils.Replace(CSocialRow, "class", class)
	if v, ok := options["content"]; !ok || len(v) == 0 {
		utils.LoadFromBlock("content", help, options)
	}
	// url link
	url := options["url"]
	if url != "#" {
		row = utils.Replace(row, "url", CSocialRowUrl)
	}
	rowHtml := utils.Transform(row, options, func(name string, value string) string {
		switch name {
		case "bg":
			return utils.ConvBG("bg", value)
		}
		return value
	})
	s.Append(rowHtml)
}

// Build is render social box
func (s *SocialBoxImp) Build() string {
	options := map[string]string{
		"alt":         "Image",
		"bg":          "info",
		"image":       "",
		"size":        "2",
		"mode":        "circle",
		"name":        "",
		"description": "",
		"rows":        s.Rows,
		"shadow":      "",
	}
	utils.LoadOptions(s.Options, options)
	return utils.Transform(CSocial, options, nil)
}

// Append add new values into the row body
func (s *SocialBoxImp) Append(body string) {
	s.Rows = fmt.Sprintf("%s%s", s.Rows, body)
}

// HTML converts
func (s *SocialBoxImp) HTML() template.HTML {
	return template.HTML(s.Build())
}

// newSocialBox create simple table
func newSocialBox(opts tags.Options) utils.BuffaloHelper {
	return SocialBoxHelper{
		SocialBoxImp: &SocialBoxImp{
			Options: opts,
		},
	}
}

// NewSocial implements bootstrap social window
func NewSocial(opts tags.Options, help hctx.HelperContext) (template.HTML, error) {
	if _, ok := opts["var"]; !ok {
		opts["var"] = "s"
	}
	return utils.BuffaloHelperCallback(nil, nil, opts, help, func(model interface{}, data interface{}, opts tags.Options) utils.BuffaloHelper {
		return newSocialBox(opts)
	})
}
