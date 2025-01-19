package testdata

import "github.com/moetomato/golang-journal-service-api/models"

var JournalTestData = []models.Journal{
	{
		ID:       1,
		Title:    "firstPost",
		Contents: "This is my first post",
		UserName: "moetomato",
		NiceNum:  2,
	},
	{
		ID:       2,
		Title:    "Second try",
		Contents: "Second journal post",
		UserName: "moetomato",
		NiceNum:  4,
	},
}
