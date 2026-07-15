package config

import (
	"flag"
	"log"
	"os"

	"github.com/GourabDas18/g-rest/utility"
	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address"`
}

type Config struct {
	Env         string     `yaml:"env" env:"ENV" env-required:"true" env-default:"Production"`
	StoragePath string     `yaml:"storage_path"`
	HTTPServer  HTTPServer `yaml:"http_server"`
}



func MustLoad(validatorG *validator.Validate) *Config {

	var trans ut.Translator

	validatorG,trans: utility.ValidatorG()

	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {

		_flag := flag.String("config", "", "Configuration file path")
		flag.Parse()
		configPath = *_flag
		if configPath == "" {
			log.Fatalf("Configuration file not found!")
		}
	}

	_, err := os.Stat(configPath)
	if err != nil {
		log.Fatalf(`Error config file %s`, err.Error())
	}
	var config Config

	error := cleanenv.ReadConfig(configPath, &config)

	if error != nil {
		log.Fatalf(`Error in read config %s`, error.Error())
	}

	return &config

}
