package store

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const YAML_CONFIG string = ".example.script-kitty.config.yaml"

type Alias struct {
	Alias  string `yaml:"alias"`
	Script string `yaml:"script"`
}

// Define a struct matching the YAML structure
type Config struct {
	Projects map[string]struct {
		Aliases []Alias
	} `yaml:"projects"`
}

func (c Config) ContainsAlias(givenAlias string) bool {
	for _, a := range c.Projects["global"].Aliases {
		if a.Alias == givenAlias {
			return true
		}
	}
	return false
}
func (c Config) GetAliases(givenAlias string) []string {
	var aliases []string
	for _, a := range c.Projects["global"].Aliases {
		aliases = append(aliases, a.Alias)
	}
	return aliases
}

func (c *Config) AddGlobalAlias(newAlias string, newScript string) {
	globalProj := c.Projects["global"]
	globalProj.Aliases = append(globalProj.Aliases, Alias{
		Alias:  newAlias,
		Script: newScript,
	})
	c.Projects["global"] = globalProj
	c.persistNewConfig()
}

func GetConfig() (Config, error) {
	var data []byte
	var err error
	data, err = os.ReadFile(YAML_CONFIG)
	if err != nil {
		return Config{}, fmt.Errorf("error reading file: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("Error parsing YAML:", err)
		return Config{}, fmt.Errorf("error reading file: %w", err)
	}
	return config, nil
}

func (c Config) persistNewConfig() error {
	data, err := yaml.Marshal(&c)
	if err != nil {
		return fmt.Errorf("error marshalling YAML: %w", err)
	}

	err = os.WriteFile(YAML_CONFIG, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	fmt.Println("YAML file successfully written!")
	return nil
}
