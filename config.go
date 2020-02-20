package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

var (
	//CfgBuilders stores references to the config builder functions, for possible environments
	CfgBuilders = map[string]ConfigBuilder{
		"DEV":  BuildConfigJSON,
		"PROD": BuildConfigEnvVars,
	}
)

//Config stores global app configuration
type Config struct {
	LeagueURL string `json:"LeagueURL"`
	BotToken  string `json:"BotToken"`
	Port      string `json:"Port"`
	PublicURL string `json:"PublicURL"`
}

//ConfigBuilder type to represent config builder functions
type ConfigBuilder func() (Config, error)

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
		LeagueURL: os.Getenv("LEAGUE_URL"),
		BotToken:  os.Getenv("BOT_TOKEN"),
		Port:      os.Getenv("PORT"),
		PublicURL: os.Getenv("PUBLIC_URL"),
	}

	if cfg.LeagueURL == "" || cfg.BotToken == "" || cfg.Port == "" || cfg.PublicURL == "" {
		return cfg, errors.New("config field is empty")
	}

	return cfg, nil
}
