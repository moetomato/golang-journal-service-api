package api

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moetomato/golang-journal-service-api/api/loggings"
	"github.com/moetomato/golang-journal-service-api/controllers"
	"github.com/moetomato/golang-journal-service-api/services"
)

func NewRouter(db *sql.DB) *mux.Router {
	ser := services.NewAppService(db)
	jcon := controllers.NewJournalController(ser)
	ccon := controllers.NewCommentController(ser)

	r := mux.NewRouter()

	r.HandleFunc("/journal", jcon.PostJournalHandler).Methods(http.MethodPost)
	r.HandleFunc("/journal/list", jcon.JournalListHandler).Methods(http.MethodGet)
	r.HandleFunc("/journal/{id:[0-9]+}", jcon.JournalDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/journal/nice", jcon.PostNiceHandler).Methods(http.MethodPost)

	r.HandleFunc("/comment", ccon.PostCommentHandler).Methods(http.MethodPost)
	r.Use(loggings.LoggingMiddleware)

	return r
}
