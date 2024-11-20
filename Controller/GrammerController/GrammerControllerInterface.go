package GrammarController

import (
	database "github.com/E-Furqan/Markdown-Note-taking-App.git/Repositories"
	"github.com/gin-gonic/gin"
)

type GrammarController struct {
	Repo database.RepositoryInterface
}

func NewController(repo database.RepositoryInterface) *GrammarController {
	return &GrammarController{
		Repo: repo,
	}
}

type GrammarControllerInterface interface {
	GrammarCheck(c *gin.Context)
}
