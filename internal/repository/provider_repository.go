package repository

import (
	"gohtmx/internal/database"
	"gohtmx/internal/entity"
	"gohtmx/internal/logging"
	"gohtmx/internal/model"

	"github.com/pocketbase/dbx"
)

type ProviderRepository interface {
	Get(panels []*entity.PanelSession) ([]*model.ProviderModel, error)
}

type providerRepository struct {
	logger *logging.Logger
	db     *database.DB
}

func NewProviderRepository(logger *logging.Logger, db *database.DB) ProviderRepository {
	return &providerRepository{
		logger: logger,
		db:     db,
	}
}

func (*providerRepository) tableName() string {
	return "providers"
}

func (r *providerRepository) Get(panels []*entity.PanelSession) ([]*model.ProviderModel, error) {
	var panelIds []interface{}
	var providers []*model.ProviderModel

	for _, panel := range panels {
		panelIds = append(panelIds, panel.Id)
	}

	err := r.db.Dao.DB().Select("*").
		From(r.tableName()).
		Where(dbx.In("panel", panelIds...)).
		All(&providers)

	if err != nil {
		return providers, err
	}

	return providers, nil
}

func (r *providerRepository) Delete(providers []*model.ProviderModel) error {
	var providerIds []interface{}

	for _, provider := range providers {
		providerIds = append(providerIds, provider.Id)
	}

	_, err := r.db.Dao.DB().Delete(
		r.tableName(),
		dbx.In(
			"id",
			providerIds...,
		),
	).Execute()

	return err
}
