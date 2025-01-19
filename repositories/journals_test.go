package repositories_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/moetomato/golang-journal-service-api/models"
	"github.com/moetomato/golang-journal-service-api/repositories"
	testdata "github.com/moetomato/golang-journal-service-api/repositories/testData"
)

func TestSelectJournalList(t *testing.T) {
	expectedNum := len(testdata.JournalTestData)

	got, err := repositories.SelectJournalList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d journals\n", expectedNum, num)
	}
}

func TestSelectJournalByID(t *testing.T) {
	tests := []struct {
		name     string
		expected models.Journal
	}{
		{
			name:     "jounal detail subtest 1",
			expected: testdata.JournalTestData[0],
		}, {
			name:     "journal detail subtest 2",
			expected: testdata.JournalTestData[1],
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e := test.expected
			got, err := repositories.SelectJournalByID(testDB, e.ID)
			if err != nil {
				t.Fatal(err)
			}
			if got.ID != e.ID {
				t.Errorf("ID: got %d but want %d\n", got.ID, e.ID)
			}
			if got.Title != e.Title {
				t.Errorf("Title: got %s but want %s\n", got.Title, e.Title)
			}
			if got.Contents != e.Contents {
				t.Errorf("Content: got %s but want %s\n", got.Contents, e.Contents)
			}
			if got.UserName != e.UserName {
				t.Errorf("UserName: got %s but want %s\n", got.UserName, e.UserName)
			}
			if got.NiceNum != e.NiceNum {
				t.Errorf("NiceNum: got %d but want %d\n", got.NiceNum, e.NiceNum)
			}
		})
	}
}

func TestInsertJournal(t *testing.T) {
	journal := models.Journal{
		Title:    "insertTest",
		Contents: "testest",
		UserName: "moetomato",
	}

	expectedJournalNum := 9
	newJournal, err := repositories.InsertJournal(testDB, journal)
	if err != nil {
		t.Error(err)
	}
	if newJournal.ID != expectedJournalNum {
		t.Errorf("new journal id is expected %d but got %d\n", expectedJournalNum, newJournal.ID)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from journals
			where title = ? and contents = ? and username = ?
		`
		testDB.Exec(sqlStr, journal.Title, journal.Contents, journal.UserName)
	})
}
