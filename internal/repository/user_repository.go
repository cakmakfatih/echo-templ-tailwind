package repository

import (
	"database/sql"
	"errors"
	"gohtmx/internal/database"
	"gohtmx/internal/logging"
	"gohtmx/internal/model"
)

type UserRepository interface {
	AuthenticateWithEmailAndPassword(creds *LoginForm) (*model.UserModel, error)
}

type userRepository struct {
	logger *logging.Logger
	db     *database.DB
}

type LoginForm struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func NewUserRepository(logger *logging.Logger, db *database.DB) UserRepository {
	return &userRepository{
		logger: logger,
		db:     db,
	}
}

func (r *userRepository) AuthenticateWithEmailAndPassword(creds *LoginForm) (*model.UserModel, error) {
	authRecord, err := r.db.Dao.FindAuthRecordByEmail("users", creds.Email)

	if err == sql.ErrNoRows {
		return nil, errors.New("incorrect username or password")
	} else if err != nil {
		return nil, err
	}

	isPasswordCorrect := authRecord.ValidatePassword(creds.Password)

	if !isPasswordCorrect {
		return nil, errors.New("incorrect username or password")
	}

	userModel := &model.UserModel{
		Email:    authRecord.Email(),
		Username: authRecord.GetString("username"),
		Roles:    authRecord.GetStringSlice("roles"),
	}

	userModel.SetId(authRecord.Id)

	return userModel, nil
}
