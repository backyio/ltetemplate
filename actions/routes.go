package actions

import "github.com/gobuffalo/buffalo"

// registerRoutes is an helper to paths
func registerRoutes(*buffalo.App) {

	app.GET("/", HomeHandler)
	app.POST("/", HomeHandler)

	// user authentication
	g := app.Group("/auth")
	g.GET("/login", AuthShowLogin)
	g.POST("/login", AuthLogin)
	g.GET("/logout", AuthLogout)
	g.GET("/register", AuthShowRegister)
	g.POST("/register", AuthRegister)
	g.GET("/forgot", AuthShowForgot)
	g.POST("/forgot", AuthForgot)

	app.GET("/examples/bootstrap", ExamplesBootstrap)
	app.GET("/examples/widgets", ExamplesWidgets)
	app.GET("/examples/tables", ExamplesTables)
	/*
	app.ErrorHandlers[404] = func(status int, err error, c buffalo.Context) error {
		c.Render(404, r.HTML("404.html"))
		return nil
	}

	app.ErrorHandlers[500] = func(status int, err error, c buffalo.Context) error {
		c.Render(500, r.HTML("500.html"))
		return nil
	}
	*/
}
