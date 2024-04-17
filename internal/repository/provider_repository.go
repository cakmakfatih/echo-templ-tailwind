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
	Delete(ids []string) error
	Update(provider *model.ProviderModel) error
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

func (r *providerRepository) Delete(ids []string) error {
	var providerIds []interface{}

	for _, id := range ids {
		providerIds = append(providerIds, id)
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

func (r *providerRepository) Update(provider *model.ProviderModel) error {
	_, err := r.db.Dao.DB().Update(
		r.tableName(),
		dbx.Params{
			"url":         provider.Url,
			"alias":       provider.Alias,
			"method":      provider.Method,
			"method_data": provider.MethodData,
		},
		dbx.NewExp(
			"id = {:id}",
			dbx.Params{
				"id": provider.Id,
			},
		),
	).Execute()

	return err
}
