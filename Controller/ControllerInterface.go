package controller

import database "github.com/E-Furqan/Markdown-Note-taking-App.git/Repositories"

type Controller struct {
	Repo database.Repository
}

func NewController(repo database.Repository) *Controller {
	return &Controller{
		Repo: repo,
	}
}
