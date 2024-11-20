package RenderController

import (
	"html/template"
	"net/http"

	model "github.com/E-Furqan/Markdown-Note-taking-App.git/Model"
	"github.com/gin-gonic/gin"
)

func (ctrl *RenderController) RenderNotesToHtml(c *gin.Context) {
	var notesId model.NotesID
	if err := c.ShouldBind(&notesId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error while binding": err.Error()})
		return
	}
	var notes model.Notes
	notes, err := ctrl.Repo.FetchNotes(&notesId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Notes not found": err.Error()})
		return
	}
	tmpl := `
	<html>
		<head><title>{{.Title}}</title></head>
		<body>
			<h1>{{.Title}}</h1>
			<p><strong>Created:</strong> {{.CreatedTime}}</p>
			<div>{{.Content}}</div>
		</body>
	</html>`

	t, err := template.New("notes").Parse(tmpl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing template"})
		return
	}

	c.Header("Content-Type", "text/html")
	t.Execute(c.Writer, notes)
}
