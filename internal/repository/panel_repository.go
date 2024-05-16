package repository

import (
	"database/sql"
	"errors"
	"gohtmx/internal/database"
	"gohtmx/internal/entity"
	"gohtmx/internal/logging"
	"gohtmx/internal/model"
	"log/slog"

	"github.com/pocketbase/dbx"
)

type PanelRepository interface {
	GetPanelsOfUser(user *model.UserModel) ([]*model.PanelModel, error)
	Create(user *entity.UserSession, panelForm *PanelForm) (*model.PanelModel, error)
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

type PanelForm struct {
	LoginURL        string `form:"loginUrl"`
	SupportUsername string `form:"supportUsername"`
	SupportPassword string `form:"supportPassword"`
	TelegramToken   string `form:"telegramToken"`
	WhatsappToken   string `form:"whatsappToken"`
}

func (r *panelRepository) Create(user *entity.UserSession, panelForm *PanelForm) (*model.PanelModel, error) {
	(*r.logger).Info("Creating a panel")
	panelModel := &model.PanelModel{
		User:            user.Id,
		LoginURL:        panelForm.LoginURL,
		SupportUsername: panelForm.SupportUsername,
		SupportPassword: panelForm.SupportPassword,
		WhatsappToken:   panelForm.WhatsappToken,
		TelegramToken:   panelForm.TelegramToken,
	}

	err := r.db.Dao.DB().
		Select("*").
		From(r.tableName()).
		Where(dbx.NewExp("user = {:user_id}", dbx.Params{"user_id": user.Id})).
		One(&panelModel)

	if !errors.Is(err, sql.ErrNoRows) {
		return panelModel, errors.New("user already has a panel")
	}

	err = r.db.Dao.Save(panelModel)

	if err != nil {
		(*r.logger).Warn("Could not create a panel")

		return panelModel, err
	}

	return panelModel, nil
}
