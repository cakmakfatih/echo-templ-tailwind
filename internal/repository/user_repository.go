package repository

import (
	"database/sql"
	"errors"
	"gohtmx/internal/database"
	"gohtmx/internal/logging"
)

type UserRepository interface {
	AuthenticateWithEmailAndPassword(creds *LoginForm) error
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

func (r *userRepository) AuthenticateWithEmailAndPassword(creds *LoginForm) error {
	authRecord, err := r.db.Dao.FindAuthRecordByEmail("users", creds.Email)

	if err == sql.ErrNoRows {
		return errors.New("incorrect username or password")
	} else if err != nil {
		return err
	}

	isPasswordCorrect := authRecord.ValidatePassword(creds.Password)

	if !isPasswordCorrect {
		return errors.New("incorrect username or password")
	}

	return nil
}
