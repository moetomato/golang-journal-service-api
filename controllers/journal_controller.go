package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/moetomato/golang-journal-service-api/apperrors"
	"github.com/moetomato/golang-journal-service-api/common"
	"github.com/moetomato/golang-journal-service-api/controllers/services"
	"github.com/moetomato/golang-journal-service-api/models"
)

type JournalController struct {
	srvc services.JournalServicer
}

func NewJournalController(s services.JournalServicer) *JournalController {
	return &JournalController{srvc: s}
}

// GET : /journal/{id}
func (c *JournalController) JournalDetailHandler(w http.ResponseWriter, req *http.Request) {
	journalID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "path param {id} must be a number")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	journal, err := c.srvc.GetJournalByIDService(journalID)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	json.NewEncoder(w).Encode(journal)
}

// GET /journal/list?page=[0-9]+
func (c *JournalController) JournalListHandler(w http.ResponseWriter, req *http.Request) {
	const defaultPage = 1

	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			err = apperrors.BadParam.Wrap(err, "query param must be a number")
			apperrors.ErrorHandler(w, req, err)
			return
		}
	} else {
		page = defaultPage
	}
	journalList, err := c.srvc.GetJournalListService(page)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	json.NewEncoder(w).Encode(journalList)
}

// POST /journal
func (c *JournalController) PostJournalHandler(w http.ResponseWriter, req *http.Request) {
	var reqJournal models.Journal

	if err := json.NewDecoder(req.Body).Decode(&reqJournal); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	authedUserName := common.GetUserName(req.Context())
	if reqJournal.UserName != authedUserName {
		err := apperrors.UserUnmatched.Wrap(errors.New("does not match user names in req body and idtoken"), "invalid parameter")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	journal, err := c.srvc.PostJournalService(reqJournal)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	json.NewEncoder(w).Encode(journal)
}

// POST /journal/nice
func (c *JournalController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqJournal models.Journal
	if err := json.NewDecoder(req.Body).Decode(&reqJournal); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
	}

	journal, err := c.srvc.PostNiceService(reqJournal)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(journal)
}
