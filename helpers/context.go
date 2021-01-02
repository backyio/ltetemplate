package helpers

import (
	"github.com/gobuffalo/buffalo"
	"lte/helpers/models"
	"net/http"
	"strings"
)

type (
	// LteContext is the custom context for lte
	LteContext struct {
		buffalo.Context
		Site models.Site
	}
)

// NewLteContext create custom context
func NewLteContext(ctx buffalo.Context) *LteContext {
	c := &LteContext{
		Context: ctx,
		Site: models.NewSite(),
	}
	c.Init()
	return c
}

// LteCustomContext setup custom lte specific context settings
func LteCustomContext(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		return next(NewLteContext(c))
	}
}

// Init initialize context base settings
func (l *LteContext) Init() {
	l.Set("IsAuth", false)
	l.Set("RenderJustFrame", false)

	l.CheckFrame()

	l.Set("Site", &l.Site)
}

// CheckFrame set frame settings by path
func (l *LteContext) CheckFrame() {
	p := l.Request().URL.Path
	if strings.Contains(p, "/auth") {
		l.Set("IsAuth", true)
		l.Set("RenderJustFrame", true)
	}
}

// LteAuthRedirect setup custom lte specific context settings
func LteAuthRedirect(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		p := c.Request().URL.Path
		if !strings.Contains(p, "/auth") {
			v := c.Session().Get("user_logged")
			if b, ok := v.(bool); ok && !b || !ok {
				c.Redirect(http.StatusPermanentRedirect, "authLoginPath()")
				return nil
			}
		}
		return next(NewLteContext(c))
	}
}
