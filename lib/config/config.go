package config

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Parse yaml config file
func ParseConfigFile() (map[string]string, error) {
	c := make(map[string]string)

	raw, err := ioutil.ReadFile("data/config.yaml")
	if err != nil {
		return c, errors.New("unable to read config file")
	}

	c, err = YamlFileUnmarshal(raw)
	if err != nil {
		return c, errors.New("unable to unmarshal config")
	}

	return c, nil
}

// Unmarshall passed yaml file
func YamlFileUnmarshal(yml []byte) (map[string]string, error) {
	config := make(map[string]string)
	if err := yaml.Unmarshal(yml, config); err != nil {
		return config, errors.New("Unable to unmarshal config file")
	}

	return config, nil
}
