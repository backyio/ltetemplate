package models

type (
	// MenuItem represents one menu item in the application menu tree
	MenuItem struct {
		// Title contains menu item translated information
		Title string `json:"title"`
		// ID contains menu item unique identifier
		ID string `json:"id"`
		// Url contains menu item special custom url
		Url string `json:"url"`
		// Icon contains menu item icon name
		Icon string `json:"icon"`
		// Roles contains required role names for accessing to menu
		Roles []string `json:"roles"`
		// Children contains all available child menu data
		Children []MenuItem `json:"children"`
		// Folder contains a simple flag, this menu item is just an folder container
		Folder bool `json:"folder"`
	}

	// Menu is the base element of application menu
	Menu struct {
		// ChildrenList contains all attached menu item
		Children []MenuItem `json:"children"`
	}
)

// NewMenu it create new empty application menu
func NewMenu() Menu {
	return Menu{}
}
