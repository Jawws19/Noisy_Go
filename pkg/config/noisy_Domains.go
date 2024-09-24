package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// Domain represents the structure of a domain from the YAML file
type Domain struct {
	Domain   string            `yaml:"domain"`
	Protocol string            `yaml:"protocol"`
	Method   string            `yaml:"method"`
	Headers  map[string]string `yaml:"headers"`
}

// LoadDomainsFromYAML reads a YAML file and returns a slice of Domain structs
func LoadDomainsFromYAML(filePath string) ([]Domain, error) {
	// Open the YAML file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file content
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Parse the YAML content
	var domains []Domain
	err = yaml.Unmarshal(data, &domains)
	if err != nil {
		return nil, err
	}

	return domains, nil
}
