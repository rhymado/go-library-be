package repositories

import (
	"lib/internal/models"

	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	*sqlx.DB
}

func InitAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db}
}

func (a *AuthRepository) Register(body *models.UserModel, hashedPassword string) error {
	query := "INSERT INTO users (username, email, userpass) VALUES ($1, $2, $3)"
	values := []any{body.Username, body.Email, hashedPassword}
	if _, err := a.Exec(query, values...); err != nil {
		return err
	}
	return nil
}

func (a *AuthRepository) GetPassword(body *models.GetUserInfoModel) ([]models.GetUserInfoModel, error) {
	query := "SELECT id, userpass, username, user_role FROM users WHERE email = $1"
	values := []any{body.Email}
	result := []models.GetUserInfoModel{}
	if err := a.Select(&result, query, values...); err != nil {
		return nil, err
	}
	return result, nil
}
