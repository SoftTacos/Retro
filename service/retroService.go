package service

import (
	"log"

	gopg "github.com/go-pg/pg"

	models "github.com/softtacos/retroBot/models"
)

func NewRetroService(db *gopg.DB) *RetroService {
	return &RetroService{
		db: db,
	}
}

type RetroService struct {
	db *gopg.DB
}

func (rs *RetroService) HandleRetroRequest(req *models.RetroRequest) *models.Response {
	log.Println("HIT IT")
	return nil
}
