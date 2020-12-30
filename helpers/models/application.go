package models

type (
	// Application represents main application and settings
	Application struct {
		// Settings contains application settings
		Settings Settings
		// Menu contains application available menus items
		Menu Menu
	}
)

// NewApplication create and initialize new application structure
func NewApplication() Application {
	return Application{
		Settings: NewSettings(),
		Menu: NewMenu(),
	}
}
