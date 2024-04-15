package repository

import (
	"gohtmx/internal/database"
	"gohtmx/internal/logging"
	"gohtmx/internal/model"

	"github.com/pocketbase/dbx"
)

type PanelRepository interface {
	GetPanelsOfUser(user *model.UserModel) ([]*model.PanelModel, error)
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

func (r *panelRepository) GetPanelsOfUser(user *model.UserModel) ([]*model.PanelModel, error) {
	var panels []*model.PanelModel
	var panelModel *model.PanelModel

	err := r.db.App.DB().Select("*").
		From(panelModel.TableName()).
		Where(dbx.NewExp("user = {:user_id}", dbx.Params{
			"user_id": user.Id,
		})).All(&panels)

	if err != nil {
		return panels, err
	}

	return panels, nil
}
