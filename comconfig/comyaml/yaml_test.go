package comyaml_test

import (
	"testing"

	"github.com/zhangxiaofeng05/com/comconfig/comyaml"
)

func TestParseConfig(t *testing.T) {
	type Config struct {
		A string `yaml:"a" validate:"required"`
		B struct {
			RenamedC int   `yaml:"c" validate:"required"`
			D        []int `yaml:",flow" validate:"required"`
		} `yaml:"b" validate:"required"`
	}

	path := "testdata/test.yaml"
	config := Config{}

	err := comyaml.ParseConfig(path, &config)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("s: %+v", config)
}
