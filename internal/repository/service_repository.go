package repository

import (
	"gohtmx/internal/database"
	"gohtmx/internal/entity"
	"gohtmx/internal/logging"
	"gohtmx/internal/model"

	"github.com/pocketbase/dbx"
)

type ServiceRepository interface {
	Get(panels []*entity.PanelSession) ([]*model.ServiceModel, error)
	GetFromProvider(providerId string) ([]*model.ServiceModel, error)
}

type serviceRepository struct {
	logger *logging.Logger
	db     *database.DB
}

func NewServiceRepository(logger *logging.Logger, db *database.DB) ServiceRepository {
	return &serviceRepository{
		logger: logger,
		db:     db,
	}
}

type ServiceCreateForm struct {
	Provider       string `form:"provider"`
	ServiceId      int    `form:"serviceId"`
	Name           string `form:"name"`
	HasRefill      bool   `form:"hasRefill"`
	RefillDuration int    `form:"refillDuration"`
}

func (r *serviceRepository) tableName() string {
	return "services"
}

func (r *serviceRepository) Create(form *ServiceCreateForm) (*model.ServiceModel, error) {
	(*r.logger).Info("Creating a service")
	serviceModel := &model.ServiceModel{
		Provider:       form.Provider,
		ServiceId:      form.ServiceId,
		Name:           form.Name,
		HasRefill:      form.HasRefill,
		RefillDuration: form.RefillDuration,
	}

	err := r.db.Dao.Save(serviceModel)

	if err != nil {
		return serviceModel, err
	}

	return serviceModel, nil
}

func (r *serviceRepository) Get(panels []*entity.PanelSession) ([]*model.ServiceModel, error) {
	var panelIds []interface{}
	var services []*model.ServiceModel

	for _, panel := range panels {
		panelIds = append(panelIds, panel.Id)
	}

	err := r.db.Dao.DB().Select("`services`.*").
		From(r.tableName()).
		Join("INNER JOIN", "providers", dbx.In("providers.panel", panelIds...)).
		All(&services)

	if err != nil {
		return services, err
	}

	return services, nil
}

func (r *serviceRepository) GetFromProvider(providerId string) ([]*model.ServiceModel, error) {
	var services []*model.ServiceModel

	err := r.db.Dao.DB().Select("*").
		From(r.tableName()).
		Where(dbx.NewExp("provider = {:provider_id}", dbx.Params{"provider_id": providerId})).
		All(&services)

	if err != nil {
		return services, err
	}

	return services, nil
}
