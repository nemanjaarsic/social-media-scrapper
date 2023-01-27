package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var Conf Config

// Config [Root config structure]
type Config struct {
	API_port string
	Host     string
	Twitch   TwitchConfig
	Twitter  TwitterConfig
	Swagger  SwaggerConfig
}

type TwitchConfig struct {
	Host     string
	ClientID string
	Token    string
}

type TwitterConfig struct {
	Host  string
	Token string
}

type SwaggerConfig struct {
	Host   string
	Scheme string
}

// Override default values with env
func LoadEnv() {
	Conf.API_port = getEnv("API_PORT", Conf.API_port)
	Conf.Host = getEnv("HOST", Conf.Host)
	Conf.Twitch.ClientID = getEnv("TWITCH_CLIENT_ID", Conf.Twitch.ClientID)
	Conf.Twitch.Token = getEnv("TWITCH_TOKEN", Conf.Twitch.Token)
	Conf.Twitter.Token = getEnv("TWITTER_TOKEN", Conf.Twitch.Token)
	Conf.Swagger.Host = getEnv("SWAGGER_HOST", Conf.Swagger.Host)
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func LoadConfJson() error {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&Conf)
	if err != nil {
		fmt.Println("error:", err)
	}
	return err
}
