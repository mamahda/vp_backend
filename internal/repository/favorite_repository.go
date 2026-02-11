package repository

import (
	"context"
	"database/sql"

	"vp_backend/internal/domain"
)

// FavoriteRepository bertanggung jawab untuk
// mengelola interaksi database yang berkaitan
// dengan fitur properti favorit user.
type FavoriteRepository struct {
	DB *sql.DB
}

// Add menyimpan relasi antara user dan properti
// ke dalam tabel favorites.
//
// Parameter:
// - ctx        : context untuk kontrol lifecycle query
// - userID     : ID user
// - propertyID : ID properti
//
// Return:
// - error jika proses insert gagal
func (r *FavoriteRepository) Add(
	ctx context.Context,
	userID int,
	propertyID int,
) error {

	query := `
		INSERT INTO favorites (user_id, property_id) 
		VALUES (?, ?);
	`

	_, err := r.DB.ExecContext(
		ctx,
		query,
		userID,
		propertyID,
	)

	return err
}

// Remove menghapus relasi antara user dan properti
// dari tabel favorites.
//
// Parameter:
// - ctx        : context untuk kontrol lifecycle query
// - userID     : ID user
// - propertyID : ID properti
//
// Return:
// - error jika proses delete gagal
func (r *FavoriteRepository) Remove(
	ctx context.Context,
	userID int,
	propertyID int,
) error {

	query := `
		DELETE FROM favorites 
		WHERE user_id = ? AND property_id = ?;
	`

	_, err := r.DB.ExecContext(
		ctx,
		query,
		userID,
		propertyID,
	)

	return err
}

// FindAll mengambil seluruh properti
// yang difavoritkan oleh user.
//
// Parameter:
// - ctx    : context untuk kontrol lifecycle query
// - userID : ID user
//
// Return:
// - slice domain.Property
// - error jika proses query gagal
func (r *FavoriteRepository) FindAll(
	ctx context.Context,
	userID int,
) ([]domain.Property, error) {

	query := `
        SELECT 
            p.id, p.title, p.description, p.price, p.status,
            p.province, p.regency, p.district, p.address,
            p.building_area, p.land_area, p.electricity,
            p.water_source, p.bedrooms, p.bathrooms,
            p.floors, p.garage, p.carport,
            p.certificate, p.year_constructed,
            p.created_at, p.property_type_id, p.agent_id
        FROM properties p
        JOIN favorites f ON p.id = f.property_id
        WHERE f.user_id = ?
    `

	rows, err := r.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var properties []domain.Property

	for rows.Next() {
		var p domain.Property
		if err := rows.Scan(
			&p.ID, &p.Title, &p.Description, &p.Price, &p.Status,
			&p.Province, &p.Regency, &p.District, &p.Address,
			&p.BuildingArea, &p.LandArea, &p.Electricity,
			&p.WaterSource, &p.Bedrooms, &p.Bathrooms,
			&p.Floors, &p.Garage, &p.Carport,
			&p.Certificate, &p.YearConstructed,
			&p.CreatedAt, &p.PropertyTypeId, &p.AgentId,
		); err != nil {
			return nil, err
		}

		properties = append(properties, p)
	}

	return properties, nil
}
