package NotesController

import (
	"net/http"

	model "github.com/E-Furqan/Markdown-Note-taking-App.git/Model"
	"github.com/gin-gonic/gin"
)

func (Ctrl *NotesController) CreateNewNote(c *gin.Context) {
	var notes model.Notes

	if err := c.ShouldBind(&notes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error while binding": err.Error()})
		return
	}

	err := Ctrl.Repo.CreateNote(&notes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error while creating notes": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Notes created successfully": notes})
}

func (Ctrl *NotesController) UpdateNote(c *gin.Context) {
	var notes model.Notes

	if err := c.ShouldBind(&notes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error while binding": err.Error()})
		return
	}

	err := Ctrl.Repo.UpdateNote(&notes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error while updating notes": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Notes updated successfully": notes})
}

func (Ctrl *NotesController) DeleteNote(c *gin.Context) {
	var notes model.Notes

	if err := c.ShouldBind(&notes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error while binding": err.Error()})
		return
	}

	err := Ctrl.Repo.DeleteNote(&notes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error while delete notes": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Notes updated successfully": notes})
}

func (Ctrl *NotesController) ListNotes(c *gin.Context) {
	var notes []model.Notes

	err := Ctrl.Repo.ListNotes(&notes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error while delete notes": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Notes": notes})
}
