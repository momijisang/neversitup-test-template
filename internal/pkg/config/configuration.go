package config

import (
	"log"

	"github.com/spf13/viper"
)

var Config *Configuration

type Configuration struct {
	Server   ServerConfiguration
	Mongo    DatabaseConfiguration
	API      APIConfiguration
	App      AppConfiguration
	S3Server S3Configuration
}

type DatabaseConfiguration struct {
	Driver       string
	Dbname       string
	Username     string
	Password     string
	Host         string
	Port         string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}

type ServerConfiguration struct {
	Port       string
	Secret     string
	Mode       string
	IncTimeBKK int
}

type APIConfiguration struct {
	AnotherApi1 string
	AnotherApi2 string
	AnotherApi3 string
}

type S3Configuration struct {
	Region string
	Secret string
	Bucket string
	Key    string
}

type AppConfiguration struct {
	Env string
}

// SetupDB initialize configuration
func Setup(configPath string) {
	var configuration *Configuration

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	Config = configuration
}

// GetConfig helps you to get configuration data
func GetConfig() *Configuration {
	return Config
}
