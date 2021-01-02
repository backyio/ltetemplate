package models

type (
	// Site represents main application and settings
	Site struct {
		// Settings contains application settings
		Settings *Settings
		// Menu contains application available menus items
		Menu*Menu
	}
)

// NewSite create and initialize new application structure
func NewSite() Site {
	return Site{
		Settings: NewSettings(),
		Menu: NewMenu(),
	}
}
