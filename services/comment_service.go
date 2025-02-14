package services

import (
	"github.com/moetomato/golang-journal-service-api/apperrors"
	"github.com/moetomato/golang-journal-service-api/models"
	"github.com/moetomato/golang-journal-service-api/repositories"
)

func (s *AppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "failed to record data")
		return models.Comment{}, err
	}

	return newComment, nil
}
