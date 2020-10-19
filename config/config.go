package config

import (
	json "encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

type Config struct {
	Server    string `json:"server"`
	Port      int    `json:"port"`
	OutputDir string `json:"outputDir"`
}

func InitConfig() (*Config, error) {
	f, err := ioutil.ReadFile("config/config.json")
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

func GetLogger(name ...string) *log.Entry {
	if len(name) > 0 {
		return log.WithField("module", name[0])
	}
	return log.WithField("module", nil)
}
