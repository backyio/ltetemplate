package models

import "time"

type (
	// Notification represents user system notification message in one category
	Notification struct {
		// CategoryID contains the notification category id
		CategoryID string `json:"category_id"`
		// CategoryIcon contains the item icon name
		CategoryIcon string `json:"category_icon"`
		// TranslationID contains i18n id
		TranslationID string `json:"translation_id"`
		// CreatedAt contains the creation date of first message in the category
		CreatedAt time.Time `json:"created_at"`
		// ItemCount contains them notification count in the category
		ItemCount int `json:"item_count"`
		// Priority contains the notification category priority level
		Priority int `json:"priority"`
	}

	// Notifications contains collection of received notifications
	Notifications struct {
		Items []Notification `json:"items"`
	}
)
