package helpers

import (
	"log"

	"github.com/spf13/viper"
)

func GetEnvVar(key string) string {

	viper.AddConfigPath("./config")

	// Read the config file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	return viper.GetString(key)

}
