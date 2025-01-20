package controllers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

const host = "localhost"

func TestJournalListHandler(t *testing.T) {
	var tests = []struct {
		name       string
		query      string
		resultCode int
	}{
		{name: "number query", query: "1", resultCode: http.StatusOK},
		{name: "alphabet query", query: "hogehoge", resultCode: http.StatusBadRequest},
	}

	for _, test := range tests {

		url := fmt.Sprintf("http://%s:8080/journal/list?page=%s", host, test.query)
		req := httptest.NewRequest(http.MethodGet, url, nil)

		res := httptest.NewRecorder()

		jcon.JournalListHandler(res, req)

		if res.Code != test.resultCode {
			t.Errorf("unexpected StatusCode: want %d but %d\n", test.resultCode, res.Code)
		}
	}
}

func TestJournalDetailHandler(t *testing.T) {
	var tests = []struct {
		name       string
		journalID  string
		resultCode int
	}{
		{name: "number pathparam", journalID: "1", resultCode: http.StatusOK},
		{name: "alphabet pathparam", journalID: "aaa", resultCode: http.StatusNotFound},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := fmt.Sprintf("http://%s:8080/journal/%s", host, tt.journalID)
			req := httptest.NewRequest(http.MethodGet, url, nil)

			res := httptest.NewRecorder()

			r := mux.NewRouter()
			r.HandleFunc("/journal/{id:[0-9]+}", jcon.JournalDetailHandler).Methods(http.MethodGet)
			r.ServeHTTP(res, req)

			if res.Code != tt.resultCode {
				t.Errorf("unexpected StatusCode: want %d but %d\n", tt.resultCode, res.Code)
			}
		})
	}
}
