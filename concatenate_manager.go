package concatenate

import (
	"errors"
	"os"
)

// Manager manage a map of sources to concatenate
type Manager map[string][]string

// NewManager return a new Manager element
func NewManager() *Manager {
	m := Manager{}
	m = make(map[string][]string, 0)
	return &m
}

// Set a name and its sources to the Manager
func (m *Manager) Set(name string, src ...string) {
	(*m)[name] = src
}

// TotalFilesInSet return the number of files in a set
func (m *Manager) TotalFilesInSet(name string) int {
	return len((*m)[name])
}

// TotalFiles return the number of files
func (m *Manager) TotalFiles() int {
	counter := 0
	for _, v := range *m {
		counter += len(v)
	}
	return counter
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
