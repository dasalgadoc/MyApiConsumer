package configs

import (
	"gopkg.in/yaml.v3"
	"os"
)

const (
	Resource_GetRestApi    = "GetRestApi"
	Resource_PostRestApi   = "PostRestApi"
	Resource_DeleteRestApi = "DeleteRestApi"
	CsvIoType              = "csv"
	JsonIoType             = "json"
)

type Config struct {
	IO      IO       `yaml:"io"`
	Clients []Client `yaml:"clients"`
}

type IO struct {
	FolderLocation string `yaml:"location"`
	InputFileName  string `yaml:"file_name"`
}

type Client struct {
	Name    string            `yaml:"name"`
	Type    string            `yaml:"type"`
	Path    string            `yaml:"path"`
	Headers map[string]string `yaml:"headers"`
}

func LoadConfig(file string) (Config, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func (c *Config) GetRegisteredClientsNames() []string {
	registeredClients := []string{}
	for _, client := range c.Clients {
		registeredClients = append(registeredClients, client.Name)
	}
	return registeredClients
}
