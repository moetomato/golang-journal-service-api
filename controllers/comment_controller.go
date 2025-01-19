package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/moetomato/golang-journal-service-api/models"
)

// POST /comment
func (c *JournalAppController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "failed to decode json\n", http.StatusBadRequest)
	}

	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "failed internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(comment)
}
