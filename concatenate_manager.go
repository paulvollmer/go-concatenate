package concatenate

import "errors"

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

// Process a given set
func (m *Manager) Process(name string) error {
	d, ok := (*m)[name]
	if !ok {
		return errors.New(name + " not found")
	}
	err := FilesToFile(name, "\n", d...)
	return err
}

// ProcessAll run the Process func at all sets
func (m *Manager) ProcessAll() error {
	for k, v := range *m {
		err := FilesToFile(k, "\n", v...)
		if err != nil {
			return err
		}
	}
	return nil
}
