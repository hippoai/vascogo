package vascogo

import "github.com/hippoai/goerr"

func ErrUnsupportedPropertyFilter(filterType string) error {
	return goerr.New(ERR_UNSUPPORTED_PROPERTY_FILTER, map[string]interface{}{
		"filterType": filterType,
	})
}
