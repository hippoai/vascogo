package vascogo

import (
	"fmt"
	"log"
	"strings"

	"github.com/hippoai/goutil"
	"github.com/hippoai/neo4jclient"
)

// Cypher makes cypher query from the filter
// Depth first
func (filter *Filter) Cypher() *neo4jclient.Statement {

	rows := []string{}
	props := map[string]interface{}{}
	nodeIndex, edgeIndex, propertyIndex := 0, 0, 0

	// Find the first node
	rows = append(rows, fmt.Sprintf(
		"MATCH (%s:%s)", getNodeName(nodeIndex), filter.StartNodeLabel,
	))

	// shallowCypher now
	props, rows, nodeIndex, edgeIndex, propertyIndex = filter.shallowCypher(
		props, rows, nodeIndex, edgeIndex, propertyIndex,
	)

	// Add the return statement
	rows = append(rows, fmt.Sprintf("RETURN %s",
		getNodeName(0),
	))

	return neo4jclient.NewStatement(
		strings.Join(rows, "\n"),
		"ok",
		props,
	)

}

// shallowCypher does property filter at this level
// and goes to deeper level
func (filter *Filter) shallowCypher(
	props map[string]interface{},
	rows []string,
	nodeIndex, edgeIndex, propertyIndex int,
) (map[string]interface{}, []string, int, int, int) {

	startNodeName := getNodeName(nodeIndex)

	// Add the property filters for this guy
	pfs := []string{}
	for _, propertyFilter := range filter.PropertiesFilter {
		propertyIndex++
		propertyName := getPropertyName(propertyIndex)

		c, err := propertyFilter.MakeCypher(startNodeName, propertyName)
		if err != nil {
			log.Fatalf("Err - %s", goutil.Pretty(err))
		}

		pfs = append(pfs, c)

		props[propertyName] = propertyFilter.Value
	}

	if len(pfs) > 0 {
		rows = append(rows, fmt.Sprintf("WHERE %s", strings.Join(pfs, ", ")))
	}

	// Go one level down
	for _, child := range filter.Filters {

		// Find it
		findMeSlice := []string{
			fmt.Sprintf("MATCH (%s)", startNodeName),
		}
		for _, step := range child.Path {

			edgeIndex++
			nodeIndex++
			edgeName := getEdgeName(edgeIndex)
			nodeName := getNodeName(nodeIndex)

			leftConnector, rightConnector := getConnectors(step.Ltr)
			findMeSlice = append(findMeSlice, fmt.Sprintf(
				"%s[%s:%s]%s(%s:%s)",
				leftConnector, edgeName, step.EdgeLabel, rightConnector, nodeName, step.EndNodeLabel,
			))
		}

		// shallowCypher it
		rows = append(rows, strings.Join(findMeSlice, ""))

		props, rows, nodeIndex, edgeIndex, propertyIndex = child.shallowCypher(
			props, rows, nodeIndex, edgeIndex, propertyIndex,
		)

	}

	return props, rows, nodeIndex, edgeIndex, propertyIndex

}

func getConnectors(ltr bool) (string, string) {
	if ltr {
		return "-", "->"
	} else {
		return "<-", "-"
	}
}

func getNodeName(i int) string {
	return fmt.Sprintf("n%d", i)
}

func getEdgeName(i int) string {
	return fmt.Sprintf("e%d", i)
}

func getPropertyName(i int) string {
	return fmt.Sprintf("p%d", i)
}
