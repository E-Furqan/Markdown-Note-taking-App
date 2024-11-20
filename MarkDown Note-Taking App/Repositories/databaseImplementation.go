package database

import (
	model "github.com/E-Furqan/Markdown-Note-taking-App.git/Model"
)

func (repo *Repository) CreateNote(notes *model.Notes) error {

	err := repo.DB.Create(notes).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) UpdateNote(notes *model.Notes) error {
	err := repo.DB.Model(&model.Notes{}).Where("notes_id = ?", notes.NotesId).Updates(map[string]interface{}{
		"title":        notes.Title,
		"content":      notes.Content,
		"created_time": notes.CreatedTime,
	}).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) DeleteNote(notes *model.Notes) error {
	err := repo.DB.Delete(notes).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) ListNotes(notes *[]model.Notes) error {
	err := repo.DB.Find(notes).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *Repository) FetchNotes(notesId *model.NotesID) (model.Notes, error) {
	var notes model.Notes

	err := repo.DB.Where("notes_id = ?", notesId.NotesId).First(&notes).Error
	if err != nil {
		return model.Notes{}, err
	}

	return notes, nil
}
