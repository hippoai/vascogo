package vascogo

import "fmt"

// PropertyFilter filters on a node property
type PropertyFilter struct {
	PropertyName string      `json:"propertyName"`
	FilterType   string      `json:"filterType"`
	Value        interface{} `json:"value"`
}

// NewPropertyFilter instanciates
func NewPropertyFilter(propertyName, filterType string, value interface{}) *PropertyFilter {
	return &PropertyFilter{
		PropertyName: propertyName,
		FilterType:   filterType,
		Value:        value,
	}
}

// MakeCypher -
func (pf *PropertyFilter) MakeCypher(startNodeName, propertyName string) (string, error) {

	switch pf.FilterType {
	case "equals":
		return fmt.Sprintf("%s.%s = {props}.%s",
			startNodeName,
			pf.PropertyName,
			propertyName,
		), nil

	case "regex":
		return fmt.Sprintf("%s.%s =~ {props}.%s",
			startNodeName,
			pf.PropertyName,
			propertyName,
		), nil

	case "in":
		return fmt.Sprintf("%s.%s IN {props}.%s",
			startNodeName,
			pf.PropertyName,
			propertyName,
		), nil

	default:
		return "", ErrUnsupportedPropertyFilter(pf.FilterType)
	}

}
