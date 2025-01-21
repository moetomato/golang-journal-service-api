package services

import (
	"database/sql"
	"errors"
	"sync"

	"github.com/moetomato/golang-journal-service-api/apperrors"
	"github.com/moetomato/golang-journal-service-api/models"
	"github.com/moetomato/golang-journal-service-api/repositories"
)

func (s *AppService) GetJournalByIDService(journalID int) (models.Journal, error) {
	var journal models.Journal
	var comments []models.Comment
	var journalGetErr, commentGetErr error

	var jmu sync.Mutex
	var cmu sync.Mutex

	var wg sync.WaitGroup
	wg.Add(2)

	go func(db *sql.DB, journalID int) {
		defer wg.Done()
		j, e := repositories.SelectJournalByID(s.db, journalID)
		jmu.Lock()
		journal, journalGetErr = j, e
		jmu.Unlock()
	}(s.db, journalID)

	go func(db *sql.DB, journalID int) {
		defer wg.Done()
		c, e := repositories.SelectCommentList(s.db, journalID)
		cmu.Lock()
		comments, commentGetErr = c, e
		cmu.Unlock()

	}(s.db, journalID)

	wg.Wait()

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
