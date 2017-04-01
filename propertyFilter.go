package vascogo

import "fmt"

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

func (pf *PropertyFilter) MakeCypher(startNodeName, propertyName string) string {

	switch pf.FilterType {
	case "equals":
		return fmt.Sprintf("%s.%s = {props}.%s",
			startNodeName,
			pf.PropertyName,
			propertyName,
		)

	case "regex":
		return fmt.Sprintf("%s.%s =~ {props}.%s",
			startNodeName,
			pf.PropertyName,
			propertyName,
		)

	default:
		return ""
	}

}
