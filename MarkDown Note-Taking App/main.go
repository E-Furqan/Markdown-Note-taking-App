package main

import (
	GrammarController "github.com/E-Furqan/Markdown-Note-taking-App.git/Controller/GrammerController"
	"github.com/E-Furqan/Markdown-Note-taking-App.git/Controller/NotesController"
	"github.com/E-Furqan/Markdown-Note-taking-App.git/Controller/RenderController"
	databaseConfig "github.com/E-Furqan/Markdown-Note-taking-App.git/DatabaseConfig"
	environmentVariable "github.com/E-Furqan/Markdown-Note-taking-App.git/EnviormentVariable"
	database "github.com/E-Furqan/Markdown-Note-taking-App.git/Repositories"
	routes "github.com/E-Furqan/Markdown-Note-taking-App.git/Routes"
	"github.com/gin-gonic/gin"
)

func main() {
	DatabaseEnv := environmentVariable.ReadEnv()

	databaseConfig := databaseConfig.NewDatabase(DatabaseEnv)
	db := databaseConfig.Connection()

	var repo database.RepositoryInterface = database.NewRepository(db)

	var nCtrl NotesController.NotesControllerInterface
	var gCtrl GrammarController.GrammarControllerInterface
	var rCtrl RenderController.RenderControllerInterface

	nCtrl = NotesController.NewController(repo)
	gCtrl = GrammarController.NewController(repo)
	rCtrl = RenderController.NewController(repo)

	server := gin.Default()

	routes.User_routes(nCtrl, gCtrl, rCtrl, server)

	server.Run(":8084")
}
