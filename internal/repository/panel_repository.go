package repository

import (
	"gohtmx/internal/database"
	"gohtmx/internal/logging"
	"gohtmx/internal/model"
	"log/slog"

	"github.com/pocketbase/dbx"
)

type PanelRepository interface {
	GetPanelsOfUser(user *model.UserModel) ([]*model.PanelModel, error)
	Create(panel *model.PanelModel) error
}

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

func (r *panelRepository) tableName() string {
	return "panels"
}

func (r *panelRepository) GetPanelsOfUser(user *model.UserModel) ([]*model.PanelModel, error) {
	(*r.logger).Info("Getting panels of the user", slog.String("id", user.Id))
	var panels []*model.PanelModel

	err := r.db.App.DB().Select("*").
		From(r.tableName()).
		Where(dbx.NewExp("user = {:user_id}", dbx.Params{
			"user_id": user.Id,
		})).All(&panels)

	if err != nil {
		return panels, err
	}

	return panels, nil
}

func (r *panelRepository) Create(panel *model.PanelModel) error {
	return nil
}
