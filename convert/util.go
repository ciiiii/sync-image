package convert

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ReplacerGenerator(p string, s []string) func(path string, fi os.FileInfo, err error) error {
	replacer := strings.NewReplacer(s...)
	return func(path string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fi.IsDir() {
			return nil
		}

		matched, err := filepath.Match(`*.yaml`, fi.Name())
		if err != nil {
			return err
		}

		if matched {
			fmt.Println(strings.Replace(path, p, "", 1))
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			newContent := replacer.Replace(string(content))
			if err := ioutil.WriteFile(path, []byte(newContent), 0); err != nil {
				return err
			}
		}
		return nil
	}
}

// Deprecated
func convertMap2StringList(m map[string]string) []string {
	var s []string
	for k, v := range m {
		s = append(s, k, v)
	}
	return s
}

func NewTempFile() (string, error) {
	dir := os.TempDir()
	f, err := ioutil.TempFile(dir, "sync-image")
	if err != nil {
		return "", err
	}
	p := f.Name()
	defer f.Close()
	return p, nil
}
