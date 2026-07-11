package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

var en = en.New()
var uni = ut.New(en, en)

var Trans, _ = uni.GetTranslator("en")

type HTTPServer struct {
	Addr string `yaml:"address"`
}

type Config struct {
	Env         string     `yaml:"env" env:"ENV" env-required:"true" env-default:"Production"`
	StoragePath string     `yaml:"storage_path"`
	HTTPServer  HTTPServer `yaml:"http_server"`
}

func MustLoad() *Config {

	enTranslations.RegisterDefaultTranslations(validate, Trans)
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
