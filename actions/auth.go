package actions

import (
    "net/http"
    
	"github.com/gobuffalo/buffalo"
)

// AuthLogin default implementation.
func AuthLogin(c buffalo.Context) error {
	c.Session().Set("user_logged", true)
	return c.Redirect(http.StatusPermanentRedirect, "rootPath()")
}

// AuthLogout default implementation.
func AuthLogout(c buffalo.Context) error {
	c.Session().Set("user_logged", false)
	return c.Render(http.StatusOK, r.HTML("auth/logout.plush.html"))
}

// AuthRegister default implementation.
func AuthRegister(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.String("User registered"))
}

// AuthShowRegister default implementation.
func AuthShowRegister(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("auth/register.plush.html"))
}

// AuthForgot default implementation.
func AuthForgot(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.String("Email sent"))
}

// AuthShowForgot default implementation.
func AuthShowForgot(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("auth/forgot.plush.html"))
}

// AuthShowLogin default implementation.
func AuthShowLogin(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("auth/login.plush.html"))
}

