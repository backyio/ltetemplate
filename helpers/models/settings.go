package models

const (
	// LayoutTopNav specify site template use top navigation
	LayoutTopNav = "topnav"
	// LayoutSideBar specify site template use sidebar navigation
	LayoutSideBar = "sidebar"
	// DefaultLayout contains default value of application layout
	DefaultLayout = LayoutSideBar
	// DefaultTabbedIFrame contains default value of tabbed mode
	DefaultTabbedIFrame = true
	// DefaultFixedFooter contains application footer mode
	DefaultFixedFooter = true
	// DefaultFixedSideBarCustomArea contains default value of custom area visibility
	DefaultFixedSideBarCustomArea = true
	// DefaultFooterVersion contains default value of version on fixed footer
	DefaultFooterVersion = "1.0"
	// DefaultFooterCopyRight contains default value of copyright on fixed footer
	DefaultFooterCopyRight = "AdminLTE 2019-2021"
	// DefaultChatMessages contains default value of chat messages
	DefaultChatMessages = true
	// DefaultNotifications contains default value of notification messages
	DefaultNotifications = true
	// DefaultFullScreen contains default value of fullscreen view
	DefaultFullScreen = true
	// DefaultCollapsedSideBar contains default value of collapsed side bar
	DefaultCollapsedSideBar = false
	// DefaultBreadcrumbs contains default value of breadcrumbs
	DefaultBreadcrumbs = true
	// DefaultFixedSideBar contains default value of FixedSideBar
	DefaultFixedSideBar = false
	// DefaultBoxed contains default value of Boxed layout
	DefaultBoxed = false
	// DefaultNoNavBarBorder contains default value of NoNavBarBorder
	DefaultNoNavBarBorder = false
	// DefaultBrandSmallText contains default value of DefaultBrandSmallText
	DefaultBrandSmallText = false
	// DefaultFooterSmallText contains default value of DefaultFooterSmallText
	DefaultFooterSmallText = false
	// DefaultSideBarSmallText contains default value of DefaultSideBarSmallText
	DefaultSideBarSmallText = false
	// DefaultSideBarFlat contains default value of DefaultSideBarFlat
	DefaultSideBarFlat = false
	// DefaultSideBarCompact contains default value of DefaultSideBarCompact
	DefaultSideBarCompact = false
)

type (
	// Settings contains current application default settings
	Settings struct {
		// Layout define application layout style
		Layout string `json:"layout"`
		// Boxed define application layout boxed style
		Boxed bool `json:"boxed"`
		// TabbedIFrame define the application pages will open in different tab
		TabbedIFrame bool `json:"tabbed_i_frame"`
		// FixedSideBarCustomArea define custom area enabled on bottom side
		FixedSideBarCustomArea bool `json:"fixed_side_bar_custom_area"`
		// FixedFooter define the application footer showed as fixed line
		FixedFooter bool `json:"fixed_footer"`
		// FixedSideBar define the navigation bar is fixed on site
		FixedSideBar bool `json:"fixed_side_bar"`
		// FooterVersion is contains current application version
		FooterVersion string `json:"footer_version"`
		// FooterCopyRight is contains current application copyright message
		FooterCopyRight string `json:"footer_copy_right"`
		// ChatMessages is contains current application chat message history list
		ChatMessages bool `json:"chat_messages"`
		// Notifications is contains current application notifications history list
		Notifications bool `json:"notifications"`
		// Fullscreen is contains fullscreen button enabled state
		Fullscreen bool `json:"fullscreen"`
		// CollapsedSideBar is contains collapsed sidebar enabled state
		CollapsedSideBar bool `json:"collapsed_side_bar"`
		// Breadcrumbs is contains breadcrumbs enabled state
		Breadcrumbs bool `json:"breadcrumbs"`
		// NoNavBarBorder control top navbar border size
		NoNavBarBorder bool `json:"no_nav_bar_border"`
		// BrandSmallText is control bran font size on top bar
		BrandSmallText bool `json:"brand_small_text"`
		// FooterSmallText is control font size on footer
		FooterSmallText bool `json:"footer_small_text"`
		// SideBarSmallText is control sidebar font size on side bar
		SideBarSmallText bool `json:"side_bar_small_text"`
		// SideBarFlat is control sidebar flat style
		SideBarFlat bool `json:"side_bar_flat"`
		// SideBarCompact is control sidebar compact style
		SideBarCompact bool `json:"side_bar_compact"`
	}
)

// NewApplicationSettings returns new application settings with default values
func NewSettings() *Settings {
	return &Settings{
		Layout:                 DefaultLayout,
		TabbedIFrame:           DefaultTabbedIFrame,
		FixedFooter:            DefaultFixedFooter,
		FixedSideBar:           DefaultFixedSideBar,
		FixedSideBarCustomArea: DefaultFixedSideBarCustomArea,
		FooterVersion:          DefaultFooterVersion,
		FooterCopyRight:        DefaultFooterCopyRight,
		ChatMessages:           DefaultChatMessages,
		Notifications:          DefaultNotifications,
		Fullscreen:             DefaultFullScreen,
		CollapsedSideBar:       DefaultCollapsedSideBar,
		Breadcrumbs:            DefaultBreadcrumbs,
		Boxed:                  DefaultBoxed,
		NoNavBarBorder:         DefaultNoNavBarBorder,
		BrandSmallText:         DefaultBrandSmallText,
		FooterSmallText:        DefaultFooterSmallText,
		SideBarSmallText:       DefaultSideBarSmallText,
		SideBarFlat:            DefaultSideBarFlat,
		SideBarCompact:         DefaultSideBarCompact,
	}
}
