package concatenate

import "errors"

type Manager map[string][]string

func NewManager() *Manager {
	m := Manager{}
	m = make(map[string][]string, 0)
	return &m
}
func (m *Manager) Set(name string, src ...string) {
	(*m)[name] = src
}

func (m *Manager) Process(name string) error {
	d, ok := (*m)[name]
	if !ok {
		return errors.New(name + " not found")
	}
	err := FilesToFile(name, "\n", d...)
	return err
}

func (m *Manager) ProcessAll() error {
	for k, v := range *m {
		err := FilesToFile(k, "\n", v...)
		if err != nil {
			return err
		}
	}
	return nil
}
