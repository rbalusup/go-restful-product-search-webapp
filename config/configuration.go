package config

import (
	"fmt"
	"go-restful-product-search-webapp/util"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	FileContextPath struct {
		ProductsFile string
	}
	DynamoDB struct {
		EnableDataSetup   bool
		ProductsTableName string
		Endpoint          string
	}
	AWS struct {
		Region              string
		CredentialsFilePath string
		CredentialsName     string
	}
}

func GetConfigMap() (Config, error) {
	viper.SetConfigName("config")

	viper.AddConfigPath(util.GetFileName(".."))
	viper.AllSettings()
	viper.AutomaticEnv()

	viper.WatchConfig() // Watch for changes to the configuration file and recompile
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return Config{}, err
	}

	var configuration Config
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		log.Println(errors.Wrap(err, "unmarshal config file"))
	}
	log.Printf("productsFile path is %s", configuration.FileContextPath.ProductsFile)
	log.Printf("DynamoDB Table name is %s", configuration.DynamoDB.ProductsTableName)
	log.Printf("DynamoDB endpoint is %s", configuration.DynamoDB.Endpoint)
	log.Printf("AWS Region is %s", configuration.AWS.Region)

	return configuration, err
}
