package concatenate

import (
	"errors"
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

// Add a filepath to the Sources array. if src is an empty string or already exist, return an error.
func (s *Sources) Add(src string) error {
	if src == "" {
		return errors.New("src is an empty string")
	}
	for _, v := range *s {
		if v == src {
			return errors.New("src already exist")
		}
	}
	*s = append(*s, src)
	return nil
}

// AddSources add a list of sources to the array
func (s *Sources) AddSources(src ...string) error {
	for _, v := range src {
		err := s.Add(v)
		if err != nil {
			return err
		}
	}
	return nil
}

// Get the source of the given index
func (s Sources) Get(i int) (string, error) {
	if i < s.Total() {
		return s[i], nil
	}
	return "", errors.New("index out of range")
}

// GetFilepaths return a list of filepaths for the given source.
// if the source is a glob, the function return all matched paths.
func (s Sources) GetFilepaths(i int) ([]string, error) {
	if i < s.Total() {
		return filepath.Glob(s[i])
	}
	return []string{}, errors.New("index out of range")
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

// GetAllDirs get dir and get all dirs
func (s *Sources) GetAllDirs() ([]string, error) {
	filepaths, err := s.GetAllFilepaths()
	if err != nil {
		return []string{}, err
	}

	cache := make([]string, 0)
	unique := make(map[string]int, 0)
	for _, v := range filepaths {
		tmpDir := filepath.Dir(v)
		_, ok := unique[tmpDir]
		if !ok {
			unique[tmpDir] = 1
			cache = append(cache, tmpDir)
		}
	}
	return cache, nil
}

// Total return the total number of sources
func (s *Sources) Total() int {
	return len(*s)
}

// ExistSource return true i the given source was found at the sets
func (s *Sources) ExistSource(src string) bool {
	for _, path := range *s {
		matches, _ := filepath.Glob(path)
		for _, item := range matches {
			if item == src {
				return true
			}
		}
	}
	return false
}
