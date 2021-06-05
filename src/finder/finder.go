package finder

type Finder interface {
	Find() []string
}

type Files struct {
	Path    string
	Ignore  []string
	results []string
}
