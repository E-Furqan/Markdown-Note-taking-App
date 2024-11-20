package NotesController

import (
	database "github.com/E-Furqan/Markdown-Note-taking-App.git/Repositories"
	"github.com/gin-gonic/gin"
)

type NotesController struct {
	Repo database.RepositoryInterface
}

func NewController(repo database.RepositoryInterface) *NotesController {
	return &NotesController{
		Repo: repo,
	}
}

type NotesControllerInterface interface {
	CreateNewNote(c *gin.Context)
	UpdateNote(c *gin.Context)
	DeleteNote(c *gin.Context)
	ListNotes(c *gin.Context)
}
