package com_yaml_test

import (
	"testing"

	"github.com/zhangxiaofeng05/com/com_config/com_yaml"
)

func TestParseConfig(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		type B struct {
			RenamedC int   `yaml:"c" validate:"required"`
			D        []int `yaml:",flow" validate:"required"`
		}
		type Config struct {
			A string `yaml:"a" validate:"required"`
			B B      `yaml:"b" validate:"required"`
			// no validate https://github.com/go-playground/validator/issues/714
			E  bool `yaml:"e"`
			FC int  `yaml:"f_c"`
			FD struct {
				RenamedC int            `yaml:"c" validate:"required"`
				D        []int          `yaml:"d" validate:"required"`
				Mp       map[string]int `yaml:"mp" validate:"required"`
			} `yaml:"f_d" validate:"required"`
		}

		path := "testdata/test.yaml"
		config := Config{}

		err := com_yaml.Parse(path, &config)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("config: %+v", config)
	})

	//file not exist
	t.Run("file not exist", func(t *testing.T) {
		path := "testdata/test_not_exist.yaml"
		config := struct{}{}
		err := com_yaml.Parse(path, &config)
		if err == nil {
			t.Fatal("want err")
		}
	})

	t.Run("file content not right", func(t *testing.T) {
		path := "testdata/test1.yaml"
		config := struct {
			Name string `yaml:"name"`
		}{}
		err := com_yaml.Parse(path, &config)
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
		err := com_yaml.Parse(path, &config)
		if err == nil {
			t.Fatal("want err")
		}
	})

}
