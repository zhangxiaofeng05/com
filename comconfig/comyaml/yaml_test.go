package comyaml_test

import (
	"testing"

	"github.com/zhangxiaofeng05/com/comconfig/comyaml"
)

func TestParseConfig(t *testing.T) {
	type Config struct {
		A string `yaml:"a"`
		B struct {
			RenamedC int   `yaml:"c"`
			D        []int `yaml:",flow"`
		} `yaml:"b"`
	}

	path := "testdata/test.yaml"
	config := Config{}

	err := comyaml.ParseConfig(path, &config)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("s: %+v", config)
}
