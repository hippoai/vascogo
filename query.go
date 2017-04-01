package vascogo

// Filter builds the query, depth first
type Filter struct {
	StartNodeLabel   string            `json:"startNodeLabel"`
	PropertiesFilter []*PropertyFilter `json:"propertiesFilter"`
	Path             []*Step           `json:"path"`
	Filters          []*Filter         `json:"filters"`
	currentFilter    *Filter
	parent           *Filter
}

// NewFilter instanciates an empty filter (= start query) with a start node
func NewFilter(startNodeLabel string) *Filter {
	return &Filter{
		StartNodeLabel:   startNodeLabel,
		PropertiesFilter: []*PropertyFilter{},
		Path:             []*Step{},
		Filters:          []*Filter{},
		currentFilter:    nil,
		parent:           nil,
	}
}

// AddPropertyFilter to current filter
func (filter *Filter) AddPropertyFilter(propertyName, filterType string, value interface{}) *Filter {
	cf := filter.GetCurrent()

	cf.PropertiesFilter = append(cf.PropertiesFilter,
		NewPropertyFilter(propertyName, filterType, value),
	)
	return filter
}

// GetCurrent is either itself or a deeply nested one
func (filter *Filter) GetCurrent() *Filter {
	if filter.currentFilter == nil {
		return filter
	}

	return filter.currentFilter
}

// AddFilter to the current filter we work on
func (filter *Filter) AddFilter(firstStep *Step, steps ...*Step) *Filter {
	cf := filter.GetCurrent()

	newFilter := &Filter{
		StartNodeLabel:   "",
		PropertiesFilter: []*PropertyFilter{},
		Path:             append([]*Step{firstStep}, steps...),
		Filters:          []*Filter{},
		currentFilter:    nil,
		parent:           cf,
	}

	cf.Filters = append(cf.Filters, newFilter)
	filter.currentFilter = newFilter

	return filter
}

// BubbleUp goes back to the filter above
func (filter *Filter) BubbleUp() *Filter {

	cf := filter.GetCurrent()
	filter.currentFilter = cf.parent

	return filter
}
