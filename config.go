// Package config helps to load a configuration file into the structure.
// This library supports yaml, json formats of files or slice of bytes.
package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"

	"gopkg.in/yaml.v3"
)

var (
	ErrDataNull       = errors.New("bytes must not be a null")
	ErrDataEmpty      = errors.New("bytes must not be an empty slice")
	ErrNotImplemented = errors.New("file's format is not implemented")
	ErrNotPointer     = errors.New("param must be a pointer to a struct or a map")
	ErrReadFile       = errors.New("read config file")
	ErrDecodeData     = errors.New("decode config")
)

// FromFile unmarshals data from a file into a structure.
func FromFile(cfg interface{}, pathToFile string) error {
	if !isPointer(cfg) {
		return ErrNotPointer
	}
	fileData, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		return fmt.Errorf("%w: %s", ErrReadFile, err)
	}
	ext := filepath.Ext(pathToFile)
	fileType := detectType(ext)

	switch fileType {
	case TypeYaml:
		err = fromYaml(fileData, cfg)
	case TypeJson:
		err = fromJson(fileData, cfg)
	default:
		return ErrNotImplemented
	}
	if err != nil {
		return fmt.Errorf("%w (%s): %s", ErrDecodeData, fileType, err)
	}
	return nil
}

// FromBytes unmarshals data from a slice of data into a structure.
func FromBytes(cfg interface{}, data []byte, t typeFile) error {
	if !isPointer(cfg) {
		return ErrNotPointer
	}
	if data == nil {
		return ErrDataNull
	}
	if len(data) == 0 {
		return ErrDataEmpty
	}
	var err error
	switch t {
	case TypeYaml:
		err = fromYaml(data, cfg)
	case TypeJson:
		err = fromJson(data, cfg)
	default:
		return ErrNotImplemented
	}
	if err != nil {
		return fmt.Errorf("%w (%s): %s", ErrDecodeData, t, err)
	}
	return nil
}

func fromJson(data []byte, cfg interface{}) error {
	return json.Unmarshal(data, cfg)
}

func fromYaml(data []byte, cfg interface{}) error {
	return yaml.Unmarshal(data, cfg)
}

func isPointer(cfg interface{}) bool {
	if cfg == nil {
		return false
	}
	t := reflect.TypeOf(cfg)
	return t.Kind() == reflect.Ptr
}
