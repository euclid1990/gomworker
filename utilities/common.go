package utilities

import (
	"github.com/euclid1990/gomworker/configs"
	"github.com/joho/godotenv"
)

func LoadEnv(file string) {
	if file == "" {
		file = ".env"
	}
	err := godotenv.Load(file)
	if err != nil {
		Logf(configs.LOG_CRITICAL, "Error loading %v file", file)
	}
}
