package configuration

import (
	"encoding/json"
	"io/ioutil"
)

var Database struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func LoadConfig(configFile string) error {
	jsonData, err := ioutil.ReadFile(configFile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonData, &Database)
	return err
}
