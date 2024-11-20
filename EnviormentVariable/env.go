package environmentVariable

import (
	"log"
	"strconv"

	model "github.com/E-Furqan/Markdown-Note-taking-App.git/Model"
	utils "github.com/E-Furqan/Markdown-Note-taking-App.git/Utils"
	"github.com/joho/godotenv"
)

func ReadEnv() model.DatabaseEnvVar {
	var envVar model.DatabaseEnvVar

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	envVar.HOST = utils.GetEnv("HOST", "0.0.0.0")
	envVar.USER = utils.GetEnv("USER1", "furqan")
	envVar.PASSWORD = utils.GetEnv("PASSWORD", "furqan")
	envVar.DB_NAME = utils.GetEnv("DB_NAME", "User")
	portStr := utils.GetEnv("PORT", "5430")
	envVar.PORT, err = strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Error converting PORT to integer: %v", err)
		envVar.PORT = 5430
	}
	return envVar
}
