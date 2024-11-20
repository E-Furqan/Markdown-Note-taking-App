package controller

import (
	"net/http"

	model "github.com/E-Furqan/Markdown-Note-taking-App.git/Model"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func (Ctrl *Controller) CreateNewNotes(c *gin.Context) {
	var notes model.Notes

	if err := c.ShouldBind(&notes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error while binding": err.Error()})
		return
	}

	err := Ctrl.Repo.CreateNotes(&notes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error while creating notes": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Notes created successfully": notes})
}

func (Ctrl *Controller) UpdateNotes(c *gin.Context) {
	var notes model.Notes

	if err := c.ShouldBind(&notes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error while binding": err.Error()})
		return
	}

	err := Ctrl.Repo.UpdateNotes(&notes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error while updating notes": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Notes updated successfully": notes})
}

func (Ctrl *Controller) DeleteNotes(c *gin.Context) {
	var notes model.Notes

	if err := c.ShouldBind(&notes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error while binding": err.Error()})
		return
	}

	err := Ctrl.Repo.DeleteNotes(&notes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error while delete notes": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Notes updated successfully": notes})
}

func GrammarCheck(c *gin.Context) {
	var request model.Notes
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := resty.New()
	resp, err := client.R().
		SetQueryParam("text", request.Content).
		SetQueryParam("language", "en-US").
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		Post("https://api.languagetool.org/v2/check")

	if err != nil || resp.StatusCode() != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check grammar"})
		return
	}

	c.Data(http.StatusOK, "application/json", resp.Body())
}
