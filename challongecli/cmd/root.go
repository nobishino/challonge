package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const challongeApikeyEnv = "CHALLONGE_APIKEY"

func init() {
	if configEnv := os.Getenv("CHALLONGE_CONFIG"); configEnv != "" {
		viper.AddConfigPath(configEnv)
		log.Println("searching challonge.yaml from", configEnv)
	}
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.challonge")
	viper.SetConfigFile("challonge.yaml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Config file not found. Using environment variables or command-line arguments.")
		} else {
			log.Printf("failed to read config: %v\n", err)
		}
	}
	rootCmd.PersistentFlags().StringP("apikey", "k", "", "API Key for access to challonge API.")
	viper.BindPFlag("apikey", rootCmd.PersistentFlags().Lookup("apikey"))
	viper.BindEnv("CHALLONGE_APIKEY")
	log.Println("env", viper.GetString("CHALLONGE_APIKEY"))

	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(destroyCmd)
}

var rootCmd = &cobra.Command{
	Use:   "challonge",
	Short: "access challonge API",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig(configPath string) {
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func getApiKey() string {
	return viper.GetString("apikey")
}
