package config

import (
	"io/ioutil"
	json "encoding/json"
)

type Config struct {
	Server string	`json:"server"`
}

func InitConfig() (*Config, error) {
	f, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	c := Config{}
	err = json.Unmarshal([]byte(f), &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

