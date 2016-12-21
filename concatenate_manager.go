package concatenate

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

// Manager create a manager to organize the source files and process the data.
type Manager struct {
	Config Config
	// TODO: cache file content to faster concatenate a set if only one source has changed.
	// Cache string
}

// NewManager return a new Manager element
func NewManager() *Manager {
	m := Manager{}
	m.Config = *NewConfig()
	return &m
}

// ReadConfig reads a json configuration file and set the data to the Config variable
func (m *Manager) ReadConfig(filename string) error {
	d, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(d, &m.Config)
	return err
}

// Process a given set
func (m *Manager) Process(filename string, perm os.FileMode) error {
	d, ok := m.Config[filename]
	if !ok {
		return errors.New(filename + " not found")
	}
	err := FilesToFile(filename, perm, "\n", d...)
	return err
}

// ProcessAll run the Process func at all sets
func (m *Manager) ProcessAll(perm os.FileMode) error {
	for k, v := range m.Config {
		err := FilesToFile(k, perm, "\n", v...)
		if err != nil {
			return err
		}
	}
	return nil
}
