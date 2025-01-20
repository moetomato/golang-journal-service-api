package testdata

import "github.com/moetomato/golang-journal-service-api/models"

var journalTestData = []models.Journal{
	{
		ID:          1,
		Title:       "initialPost",
		Contents:    "This is my first post",
		UserName:    "moetomato",
		NiceNum:     2,
		CommentList: commentTestData,
	},
	{
		ID:       2,
		Title:    "second try",
		Contents: "Second blog post",
		UserName: "moetomato",
		NiceNum:  4,
	},
}

var commentTestData = []models.Comment{
	{
		CommentID: 1,
		JournalID: 1,
		Message:   "1st comment yeah",
	},
	{
		CommentID: 2,
		JournalID: 1,
		Message:   "welcome",
	},
}
