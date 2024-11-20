package GrammarController

import (
	"encoding/json"
	"net/http"

	model "github.com/E-Furqan/Markdown-Note-taking-App.git/Model"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func (Ctrl *GrammarController) GrammarCheck(c *gin.Context) {
	var request model.NotesID
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var notes model.Notes
	notes, err := Ctrl.Repo.FetchNotes(&request)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Notes not found", "details": err.Error()})
		return
	}

	// Make API call to LanguageTool for grammar check
	client := resty.New()
	resp, err := client.R().
		SetQueryParam("text", notes.Content).
		SetQueryParam("language", "en-US").
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		Post("https://api.languagetool.org/v2/check")

	if err != nil || resp.StatusCode() != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check grammar"})
		return
	}

	var apiResponse struct {
		Matches []model.GrammarCheckResponse `json:"matches"`
	}

	if err := json.Unmarshal(resp.Body(), &apiResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse grammar check response"})
		return
	}

	var mistakes []gin.H
	for _, match := range apiResponse.Matches {
		mistakes = append(mistakes, gin.H{
			"message":      match.Message,
			"replacements": match.Replacements,
			"sentence":     match.Sentence,
			"offset":       match.Offset,
			"length":       match.Length,
		})
	}

	if len(mistakes) > 0 {
		c.JSON(http.StatusOK, gin.H{"mistakes": mistakes})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "No mistakes found"})
	}
}
