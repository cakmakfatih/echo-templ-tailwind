package repository

import (
	"gohtmx/internal/database"
	"gohtmx/internal/logging"
)

type PanelRepository interface{}

type panelRepository struct {
	logger *logging.Logger
	db     *database.DB
}

func NewPanelRepository(logger *logging.Logger, db *database.DB) PanelRepository {
	return &panelRepository{
		logger: logger,
		db:     db,
	}
}
