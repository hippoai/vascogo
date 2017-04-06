package main

import (
	"fmt"

	"github.com/hippoai/goutil"
	"github.com/hippoai/vascogo"
)

func main() {

	q := vascogo.NewFilter("Person").
		AddPropertyFilter(
			"name", "regex", "hello",
		).
		AddFilter(
			vascogo.NewStep("worksFor", "Location", true),
		).
		AddFilter(
			vascogo.NewStep("inCompany", "Company", true),
		).
		AddPropertyFilter("name", "regex", "Hello").
		BubbleUp().
		AddFilter(
			vascogo.NewStep("worksFor", "People", false),
		)

	fmt.Println(q.Cypher().Cypher)
	fmt.Println(goutil.Pretty(q.Cypher().Parameters.Props))

}
