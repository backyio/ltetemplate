package tables

import "github.com/gobuffalo/helpers/hctx"

// New returns a map of the widget helpers within this package.
func New() hctx.Map {
	return hctx.Map{
		"BSAlertBox": AlertBox,
		"BSButton": Button,
		"BSJumbotron": Jumbotron,
		"BSTab": NewTabs,
		"BSTabPage": BSTabPage,
		"BSTabPageContent": BSTabPageContent,
		"BSButtonCollapsible": ButtonCollapsible,
		"BSTooltip": ToolTip,
		"BSBadge": Badge,
		"BSBadgeHeading": BadgeHeading,
	}
}
