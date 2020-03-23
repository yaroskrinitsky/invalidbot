package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

var (
	//Builders stores references to the config builder functions, for possible environments
	Builders = map[string]Builder{
		"dev":  BuildConfigJSON,
		"prod": BuildConfigEnvVars,
	}
)

//Config stores global app configuration
type Config struct {
	Leagues []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"Leagues"`
	BotToken  string `json:"BotToken"`
	Port      string `json:"Port"`
	PublicURL string `json:"PublicURL"`
}

//Builder type to represent config builder functions
type Builder func() (Config, error)

//BuildConfigJSON builds config based on the config.json file
func BuildConfigJSON() (Config, error) {
	var cfg Config
	f, err := ioutil.ReadFile("config.json")
	if err != nil {
		return cfg, err
	}

	err = json.Unmarshal(f, &cfg)

	return cfg, err
}

//BuildConfigEnvVars builds config retrieving variables from the environment, suitable for hosting platforms
func BuildConfigEnvVars() (Config, error) {
	cfg := Config{
		BotToken:  os.Getenv("BOT_TOKEN"),
		Port:      os.Getenv("PORT"),
		PublicURL: os.Getenv("PUBLIC_URL"),
	}
	leagues := os.Getenv("LEAGUES")
	err := json.Unmarshal([]byte(leagues), &cfg.Leagues)
	if err != nil {
		return cfg, err
	}

	//TODO: ***if cfg.Leagues == ""*** || cfg.BotToken == "" || cfg.Port == "" || cfg.PublicURL == "" {
	if len(cfg.Leagues) == 0 || cfg.BotToken == "" || cfg.Port == "" || cfg.PublicURL == "" {
		return cfg, errors.New("config field is empty")
	}

	return cfg, nil
}
