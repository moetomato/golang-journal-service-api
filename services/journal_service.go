package services

import (
	"database/sql"
	"errors"

	"github.com/moetomato/golang-journal-service-api/apperrors"
	"github.com/moetomato/golang-journal-service-api/models"
	"github.com/moetomato/golang-journal-service-api/repositories"
)

func (s *AppService) GetJournalByIDService(journalID int) (models.Journal, error) {

	journal, err := repositories.SelectJournalByID(s.db, journalID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NAData.Wrap(err, "no matching record was found.")
			return models.Journal{}, err
		}
		err = apperrors.GetDataFailed.Wrap(err, "failed to get data")
		return models.Journal{}, err
	}
	commentList, err := repositories.SelectCommentList(s.db, journalID)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "failed to get data")
		return models.Journal{}, err
	}
	journal.CommentList = append(journal.CommentList, commentList...)
	return journal, nil

}

func (s *AppService) GetJournalListService(page int) ([]models.Journal, error) {
	journals, err := repositories.SelectJournalList(s.db, page)

	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "failed to get data")
		return nil, err
	}
	if len(journals) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}
	return journals, nil
}

func (s *AppService) PostJournalService(journal models.Journal) (models.Journal, error) {
	journal, err := repositories.InsertJournal(s.db, journal)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "failed to record data")
		return models.Journal{}, err
	}
	return journal, nil
}

func (s *AppService) PostNiceService(journal models.Journal) (models.Journal, error) {
	err := repositories.UpdateNiceNum(s.db, journal.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apperrors.NoTargetData.Wrap(err, "target journal does not exist")
			return models.Journal{}, err
		}
		err = apperrors.UpdateDataFailed.Wrap(err, "failed to update nice count")
		return models.Journal{}, err
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
