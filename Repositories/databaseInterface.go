package database

import (
	model "github.com/E-Furqan/Markdown-Note-taking-App.git/Model"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

type RepositoryInterface interface {
	CreateNote(notes *model.Notes) error
	UpdateNote(notes *model.Notes) error
	DeleteNote(notes *model.Notes) error
	ListNotes(notes *[]model.Notes) error
	FetchNotes(notesId *model.NotesID) (model.Notes, error)
}
