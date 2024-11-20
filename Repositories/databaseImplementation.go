package database

import (
	"fmt"

	model "github.com/E-Furqan/Markdown-Note-taking-App.git/Model"
)

func (repo *Repository) CreateNotes(notes *model.Notes) error {
	tx := repo.DB.Begin()
	err := repo.DB.Find(notes).Error
	if err != nil {
		tx.Rollback()
		return nil
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	return nil
}

func (repo *Repository) UpdateNotes(notes *model.Notes) error {
	tx := repo.DB.Begin()
	err := repo.DB.Save(notes).Error
	if err != nil {
		tx.Rollback()
		return nil
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	return nil
}

func (repo *Repository) DeleteNotes(notes *model.Notes) error {
	tx := repo.DB.Begin()
	err := repo.DB.Delete(notes).Error
	if err != nil {
		tx.Rollback()
		return nil
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("error committing transaction: %v", err)
	}

	return nil
}
