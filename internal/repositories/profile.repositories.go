package repositories

import (
	"context"
	"main/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProfileRepositories struct {
	db *pgxpool.Pool
}

func NewProfileRepository(pg *pgxpool.Pool) *ProfileRepositories {
	return &ProfileRepositories{
		db: pg,
	}
}

func (p *ProfileRepositories) GetProfileByUserID(ctx context.Context, userID int) (models.Profile, error) {
	query := `
	SELECT u.id, u.email, p.profile_image, p.first_name, p.last_name, p.phone
	FROM profile p
	JOIN users u ON p.users_id = u.id
	WHERE u.id = $1`
	var profil models.Profile
	err := p.db.QueryRow(ctx, query, userID).Scan(&profil.User, &profil.Email, &profil.Image, &profil.First_name, &profil.Last_name, &profil.Phone)
	return profil, err
}

func (p *ProfileRepositories) UpdateProfile(ctx context.Context, profile models.Profile) error {
	query := `
	UPDATE profile
	SET profile_image = $1, first_name = $2, last_name = $3, phone = $4, email = $5
	WHERE users_id = $6`
	_, err := p.db.Exec(ctx, query, profile.Image, profile.First_name, profile.Last_name, profile.Phone, profile.Email, profile.User)
	return err
}

func (p *ProfileRepositories) UpdatePassword(ctx context.Context, userID int, newHashed string) error {
	query := `UPDATE users SET password = $1, updated_at = now() WHERE id = $2`
	_, err := p.db.Exec(ctx, query, newHashed, userID)
	return err
}
