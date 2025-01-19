package repositories

import (
	"database/sql"

	"github.com/moetomato/golang-journal-service-api/models"
)

const journalsPerPage = 5

func SelectJournalByID(db *sql.DB, journalID int) (models.Journal, error) {
	const q = `
	select * 
	from journals 
	where journal_id=?;
	`
	row := db.QueryRow(q, journalID)
	if err := row.Err(); err != nil {
		return models.Journal{}, err
	}

	var journal models.Journal
	var createdTime sql.NullTime
	err := row.Scan(&journal.ID, &journal.Title, &journal.Contents, &journal.UserName, &journal.NiceNum, &createdTime)
	if err != nil {
		return models.Journal{}, err
	}

	if createdTime.Valid {
		journal.CreatedAt = createdTime.Time
	}

	return journal, nil
}

func SelectJournalList(db *sql.DB, page int) ([]models.Journal, error) {
	const q = `
		select journal_id, title, contents, username, nice
		from journals
		limit ?
		offset ?
		;	
	`
	rows, err := db.Query(q, journalsPerPage, (page-1)*5)

	if err != nil {
		return []models.Journal{}, err
	}

	var res []models.Journal

	for rows.Next() {
		var journal models.Journal
		rows.Scan(&journal.ID, &journal.Title, &journal.Contents, &journal.UserName, &journal.NiceNum)
		res = append(res, journal)
	}
	return res, nil
}

func InsertJournal(db *sql.DB, journal models.Journal) (models.Journal, error) {

	const q = `
	insert into
    journals (title, contents, username, nice, created_at)
	values
		(
			?,
			?,
			?,
			0,
			now()
		);
	`
	res, err := db.Exec(q, journal.Title, journal.Contents, journal.UserName)
	if err != nil {
		return models.Journal{}, err
	}

	id, _ := res.LastInsertId()
	return models.Journal{
		ID:       int(id),
		Title:    journal.Title,
		Contents: journal.Contents,
		UserName: journal.UserName,
	}, nil
}

func UpdateNiceNum(db *sql.DB, journalID int) error {
	const getNiceNum = `
		select nice
		from journals
		where journal_id = ?;
	`

	const updateNiceNum = `
		update journals 
		set nice = ? 
		where journal_id = ?
	`

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	row := tx.QueryRow(getNiceNum, journalID)
	if err := row.Err(); err != nil {
		tx.Rollback()
		return err
	}
	var niceNum int
	row.Scan(&niceNum)
	_, err = tx.Exec(updateNiceNum, niceNum+1, journalID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
