package routes

import (
	GrammarController "github.com/E-Furqan/Markdown-Note-taking-App.git/Controller/GrammerController"
	"github.com/E-Furqan/Markdown-Note-taking-App.git/Controller/NotesController"
	"github.com/E-Furqan/Markdown-Note-taking-App.git/Controller/RenderController"
	"github.com/gin-gonic/gin"
)

func User_routes(nCtrl NotesController.NotesControllerInterface, gCtrl GrammarController.GrammarControllerInterface,
	rCtrl RenderController.RenderControllerInterface, server *gin.Engine) {

	notes := server.Group("/notes")
	notes.POST("/create", nCtrl.CreateNewNote)
	notes.PUT("/update", nCtrl.UpdateNote)
	notes.DELETE("/delete", nCtrl.DeleteNote)
	notes.GET("/list", nCtrl.ListNotes)

	grammar := server.Group("/grammar")
	grammar.GET("/check", gCtrl.GrammarCheck)

	render := server.Group("/render")
	render.POST("/note", rCtrl.RenderNotesToHtml)

}
