package settings

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/wailsapp/wails"
)

type Settings struct {
	filename string
	runtime  *wails.Runtime
	logger   *wails.CustomLogger
}

func New() (*Settings, error) {
	result := &Settings{}
	return result, nil
}

func (s *Settings) WailsInit(runtime *wails.Runtime) error {
	s.runtime = runtime
	s.logger = s.runtime.Log.New("Settings")

	homedir, err := runtime.FileSystem.HomeDir()
	if err != nil {
		return err
	}
	s.filename = path.Join(homedir, "battery_notification_settings.json")

	s.ensureFileExists()
	return nil
}

func (s *Settings) Save(settings string) error {
	s.logger.Infof("Saving settings: %s", s.filename)
	return ioutil.WriteFile(s.filename, []byte(settings), 0600)
}

func (s *Settings) Load() (string, error) {
	s.logger.Infof("Loading settings from: %s", s.filename)
	bytes, err := ioutil.ReadFile(s.filename)
	if err != nil {
		err = fmt.Errorf("unable to open settings: %s", s.filename)
	}
	return string(bytes), err
}

func (s *Settings) ensureFileExists() {
	_, err := os.Stat(s.filename)
	if os.IsNotExist(err) {
		ioutil.WriteFile(s.filename, []byte(""), 0600)
	}
}
