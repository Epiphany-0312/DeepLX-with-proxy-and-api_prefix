package main

import (
	"flag"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Port         int    `yaml:"port"`
	Token        string `yaml:"token"`
	AuthKey      string `yaml:"auth_key"`
	DlSession    string `yaml:"dl_session"`
	ApiPrefix    string `yaml:"api_prefix"`
	UseProxy     bool   `yaml:"use_proxy"`
	ProxyAddress string `yaml:"proxy_address"`
}

func initConfig() *Config {
	cfg := &Config{}
	configPath := "config.yaml"
	if envConfigPath, exists := os.LookupEnv("CONFIG_PATH"); exists {
		configPath = envConfigPath
	}

	// Load configuration from config.yaml
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("Failed to open %s: %v", configPath, err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatalf("Failed to parse %s: %v", configPath, err)
	}

	// Override with command line flags if provided
	flag.IntVar(&cfg.Port, "port", cfg.Port, "set up the port to listen on")
	flag.IntVar(&cfg.Port, "p", cfg.Port, "set up the port to listen on")
	flag.StringVar(&cfg.DlSession, "s", cfg.DlSession, "set the dl-session for /v1/translate endpoint")
	flag.StringVar(&cfg.Token, "token", cfg.Token, "set the access token for /translate endpoint")
	flag.StringVar(&cfg.AuthKey, "authkey", cfg.AuthKey, "The authentication key for DeepL API")
	flag.StringVar(&cfg.ApiPrefix, "api_prefix", cfg.ApiPrefix, "set the API prefix")
	flag.BoolVar(&cfg.UseProxy, "use_proxy", cfg.UseProxy, "whether to use proxy")
	flag.StringVar(&cfg.ProxyAddress, "proxy_address", cfg.ProxyAddress, "set the proxy address")

	flag.Parse()

	return cfg
}
