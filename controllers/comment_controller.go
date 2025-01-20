package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/moetomato/golang-journal-service-api/apperrors"
	"github.com/moetomato/golang-journal-service-api/controllers/services"
	"github.com/moetomato/golang-journal-service-api/models"
)

type CommentController struct {
	srvc services.CommentServicer
}

func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{srvc: s}
}

// POST /comment
func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
	}

	comment, err := c.srvc.PostCommentService(reqComment)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	json.NewEncoder(w).Encode(comment)
}
