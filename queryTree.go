package vascogo

// This tree just means that
// whichever child there is will
// have a specific relationship to its parents
// and only use the previously filtered parents (that has been cached, so no need to repeat the traversal)

type FilterTree struct {
	Root     *Filter            `json:"root"`
	Children map[string]*Branch `json:"children"`
}

type Branch struct {
	Path
}
