package concatenate

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func BytesToBytes(del []byte, src ...[]byte) []byte {
	var tmp []byte
	check := len(src) - 1
	for i, v := range src {
		tmp = append(tmp, v...)
		if i < check {
			tmp = append(tmp, del...)
		}
	}
	return tmp
}

func StringsToString(del string, src ...string) string {
	return strings.Join(src, del)
}

func FilesToBytes(del string, src ...string) ([]byte, error) {
	var tmp []byte
	check := len(src) - 1
	for i, srcfile := range src {
		d, err := ioutil.ReadFile(srcfile)
		if err != nil {
			return []byte{}, err
		}
		tmp = append(tmp, d...)
		if i < check {
			tmp = append(tmp, []byte(del)...)
		}
	}
	return tmp, nil
}

func FilesToFile(file, del string, src ...string) error {
	con, err := FilesToBytes(del, src...)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(file, con, 0666)
}

type Manager map[string][]string

func NewManager() *Manager {
	m := Manager{}
	m = make(map[string][]string, 0)
	return &m
}
func (m *Manager) Set(name string, src ...string) {
	fmt.Println("add")
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
		fmt.Println("###", k, v)
		err := FilesToFile(k, "\n", v...)
		if err != nil {
			return err
		}
	}
	return nil
}
