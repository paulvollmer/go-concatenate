package concatenate

import (
	"fmt"
	"path/filepath"
)

// Config store a map of sources to concatenate.
type Config map[string]Sources

// NewConfig return a new Config element
func NewConfig() *Config {
	c := Config{}
	c = make(map[string]Sources, 0)
	return &c
}

// TotalSets return the number of total sets
func (c Config) TotalSets() int {
	return len(c)
}

// TotalFiles return the number of files
func (c Config) TotalFiles() (int, error) {
	counter := 0
	for k := range c {
		t, err := c.TotalFilesInSet(k)
		if err != nil {
			return counter, err
		}
		counter += t
	}
	return counter, nil
}

// TotalFilesInSet return the number of files of a sepcific set
func (c Config) TotalFilesInSet(name string) (int, error) {
	set, ok := c[name]
	if !ok {
		return 0, fmt.Errorf("set %q not found", name)
	}
	paths, err := set.GetAllFilepaths()
	return len(paths), err
}

// AddSet a new set (the name) and its sources (the filepaths) to the config.
func (c Config) AddSet(name string, src ...string) error {
	_, ok := c[name]
	if ok {
		return fmt.Errorf("set with name %q does not exist", name)
	}
	newSources := NewSources()
	err := newSources.AddSources(src...)
	if err != nil {
		return err
	}
	c[name] = *newSources
	return nil
}

// ExistSource return true i the given source was found at the sets
func (c Config) ExistSource(src string) bool {
	for _, v := range c {
		if v.ExistSource(src) {
			return true
		}
	}
	return false
}

// GetDirs return a list of all target directories
func (m *Manager) GetDirs() ([]string, error) {
	dirs := make([]string, 0)
	// tmpIndex := make(map[string]int, 0)
	for _, v := range m.Config {
		tmpDirs, err := v.GetAllDirs()
		if err != nil {
			return dirs, err
		}
		dirs = append(dirs, tmpDirs...)
	}
	return dirs, nil
}

// GetDirsOfSources return a list of all source directories
func (m *Manager) GetDirsOfSources() []string {
	dirs := make([]string, 0)
	tmpIndex := make(map[string]int, 0)
	for _, target := range m.Config {
		// fmt.Println(target)
		for _, fname := range target {
			kdir := filepath.Dir(fname)
			_, ok := tmpIndex[kdir]
			// fmt.Println(kdir, ok)
			if !ok {
				dirs = append(dirs, kdir)
				tmpIndex[kdir] = 1
			} else {
				tmpIndex[kdir]++
			}
		}

	}
	return dirs
}
