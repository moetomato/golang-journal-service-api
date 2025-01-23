package services

import (
	"database/sql"
	"errors"

	"github.com/moetomato/golang-journal-service-api/apperrors"
	"github.com/moetomato/golang-journal-service-api/models"
	"github.com/moetomato/golang-journal-service-api/repositories"
)

func (s *AppService) GetJournalByIDService(journalID int) (models.Journal, error) {
	var journal models.Journal
	var comments []models.Comment
	var journalGetErr, commentGetErr error

	type journalResult struct {
		journal models.Journal
		err     error
	}
	jch := make(chan journalResult)

	go func(ch chan<- journalResult, db *sql.DB, journalID int) {
		j, e := repositories.SelectJournalByID(s.db, journalID)
		ch <- journalResult{journal: j, err: e}
	}(jch, s.db, journalID)

	type commentResult struct {
		comments []models.Comment
		err      error
	}
	cch := make(chan commentResult)

	go func(ch chan<- commentResult, db *sql.DB, journalID int) {
		c, e := repositories.SelectCommentList(s.db, journalID)
		ch <- commentResult{comments: c, err: e}

	}(cch, s.db, journalID)

	for i := 0; i < 2; i++ {
		select {
		case j := <-jch:
			journal, journalGetErr = j.journal, j.err
		case c := <-cch:
			comments, commentGetErr = c.comments, c.err
		}
	}

	if journalGetErr != nil {
		if errors.Is(journalGetErr, sql.ErrNoRows) {
			journalGetErr = apperrors.NAData.Wrap(journalGetErr, "no matching record was found.")
			return models.Journal{}, journalGetErr
		}
		journalGetErr = apperrors.GetDataFailed.Wrap(journalGetErr, "failed to get data")
		return models.Journal{}, journalGetErr
	}
	if commentGetErr != nil {
		commentGetErr = apperrors.GetDataFailed.Wrap(commentGetErr, "failed to get data")
		return models.Journal{}, commentGetErr
	}

	journal.CommentList = append(journal.CommentList, comments...)
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
