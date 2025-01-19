package models

import "time"

type Comment struct {
	CommentID int       `json:"comment_id"`
	JournalID int       `json:"journal_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}
type Journal struct {
	ID          int       `json:"journal_id"`
	Title       string    `json:"title"`
	Contents    string    `json:"contents"`
	UserName    string    `json:"user_name"`
	NiceNum     int       `json:"nice"`
	CommentList []Comment `json:"comments"`
	CreatedAt   time.Time `json:"created_at"`
}
