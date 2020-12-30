package models

import "time"

type (
	// Message represents user internal messages
	Message struct {
		// ID contains the message unique identifier
		ID string `json:"id"`
		// SenderID contains the sender user unique identifier
		SenderID string `json:"sender_id"`
		// SenderName contains the sender user name
		SenderName string `json:"sender_name"`
		// Subject contains the message title
		Subject string `json:"subject"`
		// CreatedAt contains the message creation date
		CreatedAt time.Time `json:"created_at"`
		// SenderIcon contains the message sender additional icon or image name
		SenderIcon string `json:"sender_icon"`
		// State contains the message current state
		State int `json:"state"`
		// StateIcon contains the message custom icon name
		StateIcon string `json:"state_icon"`
		// Priority contains the message priority level
		Priority int `json:"priority"`
	}

	// Messages contains the collection of received messages
	Messages struct {
		Items []Message `json:"items"`
	}
)
