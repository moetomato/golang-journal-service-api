package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/moetomato/golang-journal-service-api/models"
	"github.com/moetomato/golang-journal-service-api/services"
)

type JournalAppController struct {
	service *services.AppService
}

func AppController(s *services.AppService) *JournalAppController {
	return &JournalAppController{service: s}
}

// GET /Journal/{id}
func (c *JournalAppController) JournalDetailHandler(w http.ResponseWriter, req *http.Request) {
	JournalID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	Journal, err := c.service.GetJournalByIDService(JournalID)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(Journal)
}

// GET /Journal/list
func (c *JournalAppController) JournalListHandler(w http.ResponseWriter, req *http.Request) {
	const defaultPage = 1

	queryMap := req.URL.Query()

	// get page query param
	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = defaultPage
	}

	JournalList, err := c.service.GetJournalListService(page)
	if err != nil {
		http.Error(w, "failed internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(JournalList)
}

// POST /Journal
func (c *JournalAppController) PostJournalHandler(w http.ResponseWriter, req *http.Request) {
	var reqJournal models.Journal
	if err := json.NewDecoder(req.Body).Decode(&reqJournal); err != nil {
		http.Error(w, "failed to decode json\n", http.StatusBadRequest)
	}

	Journal, err := c.service.PostJournalService(reqJournal)
	if err != nil {
		http.Error(w, "failed internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(Journal)
}

// POST /Journal/nice
func (c *JournalAppController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqJournal models.Journal
	if err := json.NewDecoder(req.Body).Decode(&reqJournal); err != nil {
		http.Error(w, "failed to decode json\n", http.StatusBadRequest)
	}

	Journal, err := c.service.PostNiceService(reqJournal)
	if err != nil {
		http.Error(w, "failed internal exec\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(Journal)
}
