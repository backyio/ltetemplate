package utils

import (
	"fmt"
	"github.com/gobuffalo/helpers/hctx"
	"github.com/gobuffalo/tags/v3"
	"strings"
)

const (
	formatLeft = "{{"
	formatRight = "}}"
)

// FormatMap is an helper for replace html options in code
func FormatMap(s string, m map[string]string, f func(name string, value string) string) (n string) {
	n = s
	for k, v := range m {
		rep := fmt.Sprintf("%s%s%s", formatLeft, formatRight, k)
		if (f != nil) {
			v = f(k, v)
		}
		n = strings.Replace(n, rep, v, -1)
	}
	return
}

// LoadOptions is loading opt tags into local option storage
func LoadOptions(opts tags.Options, options map[string]string) error {
	for k, _ := range options {
		if ov, ok := opts[k]; ok {
			options[k] = fmt.Sprintf("%v", ov)
		}
	}
	return nil
}

// LoadFromBlock is check option variable empty state and load value from block is exists
func LoadFromBlock(name string, help hctx.HelperContext, options map[string]string) bool {
	v, ok := options[name];
	if  !ok || len(v) == 0 {
		s, err := help.Block()
		if err != nil {
			options[name] = s
			return true
		}
	}
	return false
}