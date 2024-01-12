package config

import (
	"log"
	"os"

	"github.com/amerikarno/icoApi/models"
	"gopkg.in/yaml.v2"
)

func LoadConfig() (*models.SMTPConfig, error) {
	config := new(models.SMTPConfig)
	configFile, err := os.Open("config/smtpConfig.yaml")
	if err != nil {
		log.Fatalf("error opening config file: %v", err)
	}
	defer configFile.Close()

	d := yaml.NewDecoder(configFile)

	if err := d.Decode(&config); err != nil {
		log.Fatalf("error decoding config file: %v", err)
	}

	return config, nil

}
