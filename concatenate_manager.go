package concatenate

import (
	"errors"
	"os"
	"path/filepath"
)

// Manager manage a map of sources to concatenate
// TODO: chacke file body to faster concatenate files
type Manager map[string]Sources

// NewManager return a new Manager element
func NewManager() *Manager {
	m := Manager{}
	m = make(map[string]Sources, 0)
	return &m
}

// // AddSource add a list of sources to a specific set
// func (s *Sources) Add(src ...string) bool {
// 	_, ok := (*m)[name]
// 	if ok {
// 		return false
// 	}
// 	(*m)[name] = append((*m)[name], src...)
// 	return true
// }

// Add a name and its sources to the Manager
func (m *Manager) Add(name string, src ...string) bool {
	_, ok := (*m)[name]
	if ok {
		return false
	}
	(*m)[name] = src
	return true
}

// TotalSets return the number of total sets
func (m *Manager) TotalSets() int {
	return len((*m))
}

// TotalFiles return the number of files
func (m *Manager) TotalFiles() int {
	counter := 0
	for _, v := range *m {
		counter += len(v)
	}
	return counter
}

// TotalFilesInSet return the number of files of a sepcific set
func (m *Manager) TotalFilesInSet(name string) int {
	d, ok := (*m)[name]
	if !ok {
		return 0
	}
	return len(d)
}

// Process a given set
func (m *Manager) Process(filename string, perm os.FileMode) error {
	d, ok := (*m)[filename]
	if !ok {
		return errors.New(filename + " not found")
	}
	err := FilesToFile(filename, perm, "\n", d...)
	return err
}

// ProcessAll run the Process func at all sets
func (m *Manager) ProcessAll(perm os.FileMode) error {
	for k, v := range *m {
		err := FilesToFile(k, perm, "\n", v...)
		if err != nil {
			return err
		}
	}
	return nil
}

// ExistSource return true i the given source was found at the sets
func (m *Manager) ExistSource(src string) bool {
	for _, v := range *m {
		for _, v2 := range v {
			matches, _ := filepath.Glob(v2)
			for _, item := range matches {
				if item == src {
					return true
				}
			}
		}
	}
	return false
}

// GetDirs return a list of all target directories
func (m *Manager) GetDirs() []string {
	dirs := make([]string, 0)
	tmpIndex := make(map[string]int, 0)
	for k := range *m {
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
	for _, target := range *m {
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
