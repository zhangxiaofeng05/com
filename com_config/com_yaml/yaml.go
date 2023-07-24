package com_yaml

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"gopkg.in/yaml.v3"
)

func Parse(path string, config any) error {
	configBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(configBytes, config)
	if err != nil {
		return err
	}

	// reference: https://github.com/go-playground/validator/blob/master/_examples/translations/main.go
	en := en.New()
	uni := ut.New(en, en)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("en")

	validate := validator.New()
	err = en_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return err
	}

	err = validate.Struct(config)
	if err != nil {

		// translate all error at once
		errs := err.(validator.ValidationErrors)
		// returns a map with key = namespace & value = translated error
		// NOTICE: 2 errors are returned and you'll see something surprising
		// translations are i18n aware!!!!
		// eg. '10 characters' vs '1 character'
		var s strings.Builder
		for ns, val := range errs.Translate(trans) {
			s.WriteString(fmt.Sprintf("%s: %s\n", ns, val))
		}
		return errors.New(s.String())
	}
	return nil
}
