package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

//从.env中读取环境变量
type Config struct {
	DatabaseURL string
	Port string
}


//从.env中载入文件
func Load() *Config {
	if err := godotenv.Load(); err != nil {
		slog.Error("Card Load .env file Error","Load Error",err)
		panic(err)		
	}
	//读取config
	config := &Config{
		DatabaseURL: getEnv("DB_URL"),
		Port: getEnvWithDefault("PORT","8080"),
	}
	return config

}


//得到Env
func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		slog.Error("Environment get error",key,"is empty")
		panic(fmt.Sprintf("%s is empty",key))
	}
	return value
}


//得到Env with fallback
func getEnvWithDefault(key string , fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		slog.Error("Environment get error",key,"is empty")
		return fallback
	}
	return value
}