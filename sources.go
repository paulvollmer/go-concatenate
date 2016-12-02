package concatenate

// Sources store a list of source paths
type Sources []string

// NewSources return a new Sources element
func NewSources() *Sources {
	s := Sources{}
	s = make([]string, 0)
	return &s
}

// Add a source to the list of sources, if not exist
func (s *Sources) Add(src string) bool {
	for _, v := range *s {
		if v == src {
			return false
		}
	}
	*s = append(*s, src)
	return true
}

// Total return the total number of sources
func (s *Sources) Total() int {
	return len(*s)
}
