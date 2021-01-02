package models

import (
	"fmt"
	"github.com/gobuffalo/helpers/hctx"
	"html/template"
)

type (
	// MenuItem represents one menu item in the application menu tree
	MenuItem struct {
		// ID contains menu item unique identifier
		ID string `json:"id"`
		// ParentID contains menu item parent unique identifier
		ParentID string `json:"id"`
		// Title contains menu item translated information
		Title string `json:"title"`
		// Url contains menu item special custom url
		Url string `json:"url"`
		// Icon contains menu item icon name
		Icon string `json:"icon"`
		// Roles contains required role names for accessing to menu
		Roles []string `json:"roles"`
		// Folder contains a simple flag, this menu item is just an folder container
		Folder bool `json:"folder"`
		// Order represents menu order
		Order int `json:"order"`
	}

	// Menu is the base element of application menu
	Menu struct {
		// ChildrenList contains all attached menu item
		Children []*MenuItem `json:"children"`
	}
)

// NewMenu it create new empty application menu
func NewMenu() *Menu {
	return &Menu{}
}

// Render is convert menu to html
func (m *Menu) Render(help hctx.HelperContext) (template.HTML, error) {
	return template.HTML(""), nil
}

// Render is convert menu to html with user roles
func (m *Menu) RenderForUser(user *User, help hctx.HelperContext) (template.HTML, error) {
	return template.HTML(""), nil
}

// Add adding new menu item to the list
func (m *Menu) Add(parent string, item *MenuItem) error {
	if item == nil {
		return fmt.Errorf("Empty item")
	}
	if v := item.IsValid(); !v {
		return fmt.Errorf("Invalid item")
	}
	return m.internalAdd(parent, item)
}

// internalAdd adding new menu item into the tree
func (m *Menu) internalAdd(parent string, item *MenuItem) error {
	bP := false
	bId := false
	for _, v := range m.Children {
		if v.ID == parent {
			bP = true
		}
		if v.ID == item.ID {
			bId = true
		}
	}
	if !bP {
		return fmt.Errorf("Parent id not exists!")
	}
	if bId {
		return fmt.Errorf("Menu item id already exists!")
	}
	m.Children = append(m.Children, item)
	return nil
}

// IsValid returns validated state
func (m *MenuItem) IsValid() bool {
	return len(m.ID) > 0 && len(m.Title) > 0
}