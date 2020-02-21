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
	Leagues []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"Leagues"`
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
		//TODO: implement adding leags logic
		//Leagues:   os.Getenv("LEAGUE_URL"),
		BotToken:  os.Getenv("BOT_TOKEN"),
		Port:      os.Getenv("PORT"),
		PublicURL: os.Getenv("PUBLIC_URL"),
	}

	//TODO: ***if cfg.Leagues == ""*** || cfg.BotToken == "" || cfg.Port == "" || cfg.PublicURL == "" {
	if cfg.BotToken == "" || cfg.Port == "" || cfg.PublicURL == "" {
		return cfg, errors.New("config field is empty")
	}

	return cfg, nil
}
