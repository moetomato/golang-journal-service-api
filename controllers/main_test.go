package controllers_test

import (
	"testing"

	"github.com/moetomato/golang-journal-service-api/controllers"
	testdata "github.com/moetomato/golang-journal-service-api/controllers/testData"

	_ "github.com/go-sql-driver/mysql"
)

var jcon *controllers.JournalController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	jcon = controllers.NewJournalController(ser)

	m.Run()
}
