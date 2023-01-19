package comyaml_test

import (
	"testing"

	"github.com/zhangxiaofeng05/com/comconfig/comyaml"
)

func TestParseConfig(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		type Config struct {
			A string `yaml:"a" validate:"required"`
			B struct {
				RenamedC int   `yaml:"c" validate:"required"`
				D        []int `yaml:",flow" validate:"required"`
			} `yaml:"b" validate:"required"`
		}

		path := "testdata/test.yaml"
		config := Config{}

		err := comyaml.Parse(path, &config)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("s: %+v", config)
	})

	//file not exist
	t.Run("file not exist", func(t *testing.T) {
		path := "testdata/test_not_exist.yaml"
		config := struct{}{}
		err := comyaml.Parse(path, &config)
		if err == nil {
			t.Fatal("want err")
		}
	})

	t.Run("file content not right", func(t *testing.T) {
		path := "testdata/test1.yaml"
		config := struct {
			Name string `yaml:"name"`
		}{}
		err := comyaml.Parse(path, &config)
		if err == nil {
			t.Fatal("want err")
		}
	})

	t.Run("file not pass validate", func(t *testing.T) {
		path := "testdata/test2.yaml"
		config := struct {
			Name string `yaml:"name" validate:"required"`
			Age  string `yaml:"age" validate:"required"`
		}{}
		err := comyaml.Parse(path, &config)
		if err == nil {
			t.Fatal("want err")
		}
	})

}
