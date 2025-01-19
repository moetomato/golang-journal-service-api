package repositories

import (
	"database/sql"

	"github.com/moetomato/golang-journal-service-api/models"
)

func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const q = `
		insert into comments (journal_id, message, created_at) 
		values
		(?, ?, now());
	`
	res, err := db.Exec(q, comment.JournalID, comment.Message)
	if err != nil {
		return models.Comment{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return models.Comment{}, err
	}

	return models.Comment{
		CommentID: int(id),
		JournalID: comment.JournalID,
		Message:   comment.Message,
		CreatedAt: comment.CreatedAt,
	}, nil
}

func SelectCommentList(db *sql.DB, journalID int) ([]models.Comment, error) {
	const q = `
		select * 
		from comments
		where journal_id = ?;
	`
	rows, err := db.Query(q, journalID)
	if err != nil {
		return []models.Comment{}, err
	}

	var res []models.Comment
	for rows.Next() {
		var comment models.Comment
		var createdTime sql.NullTime
		rows.Scan(&comment.CommentID, &comment.JournalID, &comment.Message, &createdTime)

		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time
		}
		res = append(res, comment)
	}
	return res, nil
}
