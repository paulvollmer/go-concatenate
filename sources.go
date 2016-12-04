package concatenate

import (
	"path/filepath"
)

// Sources store a list of source paths
type Sources []string

// NewSources return a new Sources element
func NewSources() *Sources {
	s := Sources{}
	s = make([]string, 0)
	return &s
}

// Add a source to the Sources array. if src already exist, return false.
func (s *Sources) Add(src string) bool {
	if src == "" {
		return false
	}
	for _, v := range *s {
		if v == src {
			return false
		}
	}
	*s = append(*s, src)
	return true
}

// GetFilepaths return a list of filepaths for the given source.
// if the source is a glob, the function return all matched paths.
func (s Sources) GetFilepaths(i int) ([]string, error) {
	return filepath.Glob(s[i])
}

// GetAllFilepaths return a list of all filepaths
func (s *Sources) GetAllFilepaths() ([]string, error) {
	var paths []string
	for i := range *s {
		// is the path a glob?
		glob, err := s.GetFilepaths(i)
		if err != nil {
			return paths, err
		}
		paths = append(paths, glob...)
	}
	return paths, nil
}

// Total return the total number of sources
func (s *Sources) Total() int {
	return len(*s)
}
