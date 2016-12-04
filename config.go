package concatenate

import "fmt"

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
