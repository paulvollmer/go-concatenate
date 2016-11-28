package concatenate

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// BytesToBytes concatenate a list of bytes by the given delimiter
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

// StringsToString concatenate a list of strings by the given delimiter
func StringsToString(del string, src ...string) string {
	return strings.Join(src, del)
}

// FilesToBytes concatenate a list of files by the given delimiter.
// you can set a matching pattern to select the sources you want to process.
func FilesToBytes(del string, src ...string) ([]byte, error) {
	var tmp []byte

	check := len(src) - 1
	for i, srcfile := range src {
		matches, err := filepath.Glob(srcfile)
		if err != nil {
			return tmp, nil
		}

		totalMatches := len(matches)
		//fmt.Println("GLOB", srcfile, matches)
		if totalMatches == 0 {
			return tmp, errors.New("cannot find " + srcfile)
		}
		for j, matchFiles := range matches {
			d, err := ioutil.ReadFile(matchFiles)
			if err != nil {
				return tmp, err
			}
			tmp = append(tmp, d...)
			if j < totalMatches-1 {
				tmp = append(tmp, []byte(del)...)
			}
		}

		if i < check {
			tmp = append(tmp, []byte(del)...)
		}
	}
	return tmp, nil
}

// FilesToFile concatenate a list of files by the given delimiter
func FilesToFile(filename string, perm os.FileMode, del string, src ...string) error {
	con, err := FilesToBytes(del, src...)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, con, perm)
}
