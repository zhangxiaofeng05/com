package comyaml

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func ParseConfig(path string, config any) error {
	configBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(configBytes, config)
	if err != nil {
		return err
	}
	return nil
}
