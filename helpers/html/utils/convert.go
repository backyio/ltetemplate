package utils

import (
	"fmt"
	"github.com/gobuffalo/tags/v3"
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

// ConvStyle is convert style to html format
func ConvStyle(value string) string {
	return fmt.Sprintf(`style: "%s"`, value)
}

// ConvBG is converts regular name to bootstrap background
func ConvBG(prefix string, name string) string {
	if len(prefix) == 0 {
		prefix = "bg"
	}
	if len(name) == 0 {
		return name
	}
	return fmt.Sprintf("%s-%s", prefix,strings.ToLower(name))
}

// ConvColor is converts regular name to bootstrap background
func ConvColor(prefix string, name string) string {
	if len(prefix) == 0 {
		prefix = "text"
	}
	if len(name) == 0 {
		return name
	}
	return fmt.Sprintf("%s-%s", prefix,strings.ToLower(name))
}

// ConvCommon is implement common conversion
func ConvCommon(prefix string, name string) string {
	if len(prefix) == 0 {
		return ""
	}
	return fmt.Sprintf("%s-%s", prefix,strings.ToLower(name))
}

// ConvBool
func ConvBool(value string, tValue string, fValue string) string {
	if value == "true" {
		return tValue
	}
	return fValue
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

// GetValueAsString returns option value as string
func GetValueAsString(name string, def string, opts tags.Options) (string, error) {
	if v, ok := opts[name]; ok {
		if vT, okT := v.(string); okT {
			return vT, nil
		}
	} else if !ok {
		return def, nil
	}
	return "", CErrorNoVariableParameter
}

// GetValueAsBool returns option value as bool
func GetValueAsBool(name string, def bool, opts tags.Options) (bool, error) {
	if v, ok := opts[name]; ok {
		if vT, okT := v.(bool); okT {
			return vT, nil
		}
	} else if !ok {
		return def, nil
	}
	return false, CErrorNoVariableParameter
}

// GetValueAsBoolStr returns option value as bool string representation
func GetValueAsBoolStr(name string, def string, opts tags.Options) (string, error) {
	if v, ok := opts[name]; ok {
		if vT, okT := v.(bool); okT {
			return fmt.Sprintf("%t", vT), nil
		}
	} else if !ok {
		return def, nil
	}
	return "", CErrorNoVariableParameter
}