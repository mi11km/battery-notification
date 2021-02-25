package settings

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type Settings struct {
	filename string
}

func New() (*Settings, error) {
	result := &Settings{}
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	filename := path.Join(cwd, "settings.json")
	result.filename = filename
	return result, nil
}

func (s *Settings) Save(settings string) error {
	return ioutil.WriteFile(s.filename, []byte(settings), 0600)
}

func (s *Settings) Load() (string, error) {
	bytes, err := ioutil.ReadFile(s.filename)
	if err != nil {
		err = fmt.Errorf("unable to open settings: %s", s.filename)
	}
	return string(bytes), err
}
