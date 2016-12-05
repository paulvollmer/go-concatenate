package concatenate

import (
	"fmt"
	"path/filepath"
)

// Config store a map of sources to concatenate.
type Config map[string]Sources

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

// Add a new set (the name) and its sources (the filepaths) to the config.
func (c Config) Add(name string, src ...string) bool {
	_, ok := c[name]
	if ok {
		return false
	}
	c[name] = src
	return true
}

// // ExistSource return true i the given source was found at the sets
// func (m *Manager) ExistSource(src string) bool {
// 	for _, v := range m.Config {
// 		for _, v2 := range v {
// 			matches, _ := filepath.Glob(v2)
// 			for _, item := range matches {
// 				if item == src {
// 					return true
// 				}
// 			}
// 		}
// 	}
// 	return false
// }

// GetDirs return a list of all target directories
func (m *Manager) GetDirs() []string {
	dirs := make([]string, 0)
	tmpIndex := make(map[string]int, 0)
	for k := range m.Config {
		kdir := filepath.Dir(k)
		_, ok := tmpIndex[kdir]
		// fmt.Println(k, kdir, ok)
		if !ok {
			dirs = append(dirs, kdir)
			tmpIndex[kdir] = 1
		} else {
			tmpIndex[kdir]++
		}
	}
	return dirs
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
