package handler

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) indexPage(c *gin.Context) {
	// get mainpage data from DB

	t, err := template.ParseFiles("/home/cora/Documents/projects/golang-projects/online-diler/templates/base.html",
		"/home/cora/Documents/projects/golang-projects/online-diler/templates/index.html")
	if err != nil {
		log.Fatal("Error parsing templates: ", err)
	}
	err = t.ExecuteTemplate(c.Writer, "base", nil)
	if err != nil {
		log.Fatal("Error executing template: ", err)
	}
}
