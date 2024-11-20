package RenderController

import (
	database "github.com/E-Furqan/Markdown-Note-taking-App.git/Repositories"
	"github.com/gin-gonic/gin"
)

type RenderController struct {
	Repo database.RepositoryInterface
}

func NewController(repo database.RepositoryInterface) *RenderController {
	return &RenderController{
		Repo: repo,
	}
}

type RenderControllerInterface interface {
	RenderNotesToHtml(c *gin.Context)
}
