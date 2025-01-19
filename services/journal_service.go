package services

import (
	"github.com/moetomato/golang-journal-service-api/models"
	"github.com/moetomato/golang-journal-service-api/repositories"
)

func (s *AppService) GetJournalByIDService(journalID int) (models.Journal, error) {

	journal, err := repositories.SelectJournalByID(s.db, journalID)

	if err != nil {
		return models.Journal{}, err
	}
	commentList, err := repositories.SelectCommentList(s.db, journalID)
	if err != nil {
		return models.Journal{}, err
	}
	journal.CommentList = append(journal.CommentList, commentList...)
	return journal, nil

}

func (s *AppService) GetJournalListService(page int) ([]models.Journal, error) {
	journals, err := repositories.SelectJournalList(s.db, page)

	if err != nil {
		return nil, err
	}
	return journals, nil
}

func (s *AppService) PostJournalService(journal models.Journal) (models.Journal, error) {
	journal, err := repositories.InsertJournal(s.db, journal)
	if err != nil {
		return models.Journal{}, err
	}
	return journal, nil
}

func (s *AppService) PostNiceService(journal models.Journal) (models.Journal, error) {
	err := repositories.UpdateNiceNum(s.db, journal.ID)
	if err != nil {
		return models.Journal{}, nil
	}
	return models.Journal{
		ID:        journal.ID,
		Title:     journal.Title,
		Contents:  journal.Contents,
		UserName:  journal.UserName,
		NiceNum:   journal.NiceNum + 1,
		CreatedAt: journal.CreatedAt,
	}, nil
}
