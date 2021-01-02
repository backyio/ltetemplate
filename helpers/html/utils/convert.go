package utils

import (
	"fmt"
	"strings"
)

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
func ConvBG(prefix string, name string) string {
	if len(prefix) == 0 {
		prefix = "bg"
	}
	switch strings.ToLower(name) {
	case "info":
		return fmt.Sprintf("%s-info", prefix)
	case "success":
		return fmt.Sprintf("%s-success", prefix)
	case "warning":
		return fmt.Sprintf("%s-warning", prefix)
	case "danger":
		return fmt.Sprintf("%s-danger", prefix)
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
