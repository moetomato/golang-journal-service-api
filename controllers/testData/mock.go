package testdata

import "github.com/moetomato/golang-journal-service-api/models"

type serviceMock struct{}

func NewServiceMock() *serviceMock {
	return &serviceMock{}
}

func (s *serviceMock) PostJournalService(journal models.Journal) (models.Journal, error) {
	return journalTestData[1], nil
}

func (s *serviceMock) GetJournalListService(page int) ([]models.Journal, error) {
	return journalTestData, nil
}

func (s *serviceMock) GetJournalByIDService(journalID int) (models.Journal, error) {
	return journalTestData[0], nil
}

func (s *serviceMock) PostNiceService(journal models.Journal) (models.Journal, error) {
	return journalTestData[0], nil
}

func (s *serviceMock) PostCommentService(comment models.Comment) (models.Comment, error) {
	return commentTestData[0], nil
}
