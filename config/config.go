package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	// TODO: use irma.IrmaConfig here?
	IrmaURL            string `yaml:"irmaUrl,omitempty"`
	IrmaProductionMode bool   `yaml:"irmaProductionMode,omitempty"`
	Port               int    `yaml:"port,omitempty"`
	CookieSecret       string `yaml:"cookieSecret,omitempty"`
}

func getDefaultConfig() *Config {
	return &Config{
		IrmaURL:            "http://localhost:3047/api/irma",
		IrmaProductionMode: false,
		Port:               3047,
		CookieSecret:       "VeryS3cret",
	}
}

func readConfig() (*Config, error) {
	configFile, err := ioutil.ReadFile("config.yaml")
	parsedConfig := &Config{}
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(configFile, parsedConfig)

	if err != nil {
		return nil, err
	}

	return parsedConfig, nil
}

func GetConfig() *Config {
	parsedConfig, err := readConfig()
	defaultConfig := getDefaultConfig()

	if err != nil {
		fmt.Printf("Warning, cannot parse config.yaml: %v\n", err.Error())
		fmt.Println("Falling back to default config")
		return defaultConfig
	}

	// Fill empty fields with defaults
	if parsedConfig.Port == 0 {
		parsedConfig.Port = defaultConfig.Port
	}
	if parsedConfig.IrmaURL == "" {
		parsedConfig.IrmaURL = defaultConfig.IrmaURL
	}
	if parsedConfig.CookieSecret == "" {
		parsedConfig.CookieSecret = defaultConfig.CookieSecret
	}
	return parsedConfig
}
