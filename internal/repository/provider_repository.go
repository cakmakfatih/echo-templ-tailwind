package repository

import (
	"encoding/json"
	"gohtmx/internal/database"
	"gohtmx/internal/entity"
	"gohtmx/internal/logging"
	"gohtmx/internal/model"

	"github.com/pocketbase/dbx"
)

type ProviderRepository interface {
	Create(providerForm *ProviderForm) (*model.ProviderModel, error)
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

type ProviderForm struct {
	PanelID         string `form:"panel"`
	URL             string `form:"url"`
	Method          string `form:"method"`
	SupportUsername string `form:"supportUsername"`
	SupportPassword string `form:"supportPassword"`
	TelegramChatID  string `form:"telegramChatId"`
	Alias           string `form:"alias"`
}

func (r *providerRepository) Create(providerForm *ProviderForm) (*model.ProviderModel, error) {
	methodData := map[string]string{}

	if providerForm.Method == "telegram" {
		methodData["telegram_chat_id"] = providerForm.TelegramChatID
	} else if providerForm.Method == "web" {
		methodData["support_username"] = providerForm.SupportUsername
		methodData["support_password"] = providerForm.SupportPassword
	}

	methodDataJsonStr, err := json.Marshal(methodData)

	if err != nil {
		(*r.logger).Warn(err.Error())

		return nil, err
	}

	providerModel := &model.ProviderModel{
		Panel:      providerForm.PanelID,
		Url:        providerForm.URL,
		Alias:      providerForm.Alias,
		Method:     model.Method(providerForm.Method),
		MethodData: string(methodDataJsonStr),
	}

	err = r.db.Dao.Save(providerModel)

	if err != nil {
		(*r.logger).Warn(err.Error())
	}

	return providerModel, nil
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
		OrderBy("created desc").
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
