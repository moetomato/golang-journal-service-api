package services

import "database/sql"

type AppService struct {
	db *sql.DB
}

func NewAppService(db *sql.DB) *AppService {
	return &AppService{db: db}
}
