package vascogo

// Step is a named edge
type Step struct {
	EdgeLabel    string `json:"edgeLabel"`
	EndNodeLabel string `json:"endNodeLabel"`
	Ltr          bool   `json:"ltr"`
}

// NewStep instanciates
func NewStep(edgeLabel, endNodeLabel string, ltr bool) *Step {
	return &Step{
		EdgeLabel:    edgeLabel,
		EndNodeLabel: endNodeLabel,
		Ltr:          ltr,
	}
}
