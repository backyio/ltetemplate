package utils

import "strings"

// ConvShadow is converts regular name to bootstrap shadow
func ConvShadow(name string) string {
	switch strings.ToLower(name) {
	case "small":
		return "shadow-sm"
	case "regular":
		return "shadow"
	case "large":
		return "shadow-lg"
	}
	return ""
}

// ConvBG is converts regular name to bootstrap background
func ConvBG(name string) string {
	switch strings.ToLower(name) {
	case "info":
		return "bg-info"
	case "success":
		return "bg-success"
	case "warning":
		return "bg-warning"
	case "danger":
		return "bg-danger"
	}
	return name
}

// ConvInfoContentType convert info box type name to class
func ConvInfoContentType(name string) string {
	switch strings.ToLower(name) {
	case "number":
		return "info-box-number"
	case "text":
		return "info-box-text"
	}
	return name
}
