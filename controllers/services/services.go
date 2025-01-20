package services

import "github.com/moetomato/golang-journal-service-api/models"

type JournalServicer interface {
	GetJournalListService(page int) ([]models.Journal, error)
	GetJournalByIDService(journalID int) (models.Journal, error)
	PostJournalService(journal models.Journal) (models.Journal, error)
	PostNiceService(journal models.Journal) (models.Journal, error)
}

type CommentServicer interface {
	PostCommentService(comment models.Comment) (models.Comment, error)
}
