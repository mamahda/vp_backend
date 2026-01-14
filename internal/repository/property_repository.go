package repository

import (
	"context"
	"database/sql"
	"errors"

	"vp_backend/internal/domain"
)

type PropertyRepository struct {
	DB *sql.DB
}

func (r *PropertyRepository) Create(ctx context.Context, p *domain.Property) error {
	query := `
		INSERT INTO properties
		(title, description, price, status, province, regency, district, address,
		building_area, land_area, electricity, water_source, bedrooms, bathrooms,
		floors, garage, carport, certificate, year_constructed, property_type_id, user_id)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := r.DB.ExecContext(ctx, query,
		p.Title, p.Description, p.Price, p.Status, p.Province,
		p.Regency, p.District, p.Address, p.BuildingArea,
		p.LandArea, p.Electricity, p.WaterSource, p.Bedrooms,
		p.Bathrooms, p.Floors, p.Garage, p.Carport,
		p.Certificate, p.YearConstructed, p.PropertyTypeId, p.UserId,
	)

	return err
}

func (r *PropertyRepository) FindByID(ctx context.Context, id int) (*domain.Property, error) {
	query := `SELECT * FROM properties WHERE id = ?`

	p := domain.Property{}
	err := r.DB.QueryRowContext(ctx, query, id).Scan(
		&p.ID, &p.Title, &p.Description, &p.Price, &p.Status,
		&p.Province, &p.Regency, &p.District, &p.Address,
		&p.BuildingArea, &p.LandArea, &p.Electricity,
		&p.WaterSource, &p.Bedrooms, &p.Bathrooms,
		&p.Floors, &p.Garage, &p.Carport,
		&p.Certificate, &p.YearConstructed,
		&p.CreatedAt, &p.PropertyTypeId, &p.UserId,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("property not found")
	}

	return &p, err
}

func (r *PropertyRepository) FindAll(ctx context.Context) ([]domain.Property, error) {
	rows, err := r.DB.QueryContext(ctx, `SELECT * FROM properties`)
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
			&p.CreatedAt, &p.PropertyTypeId, &p.UserId,
		); err != nil {
			return nil, err
		}
		properties = append(properties, p)
	}

	return properties, nil
}

func (r *PropertyRepository) Update(ctx context.Context, p *domain.Property) error {
	query := `
		UPDATE properties SET
		title=?, description=?, price=?, status=?, province=?, regency=?,
		district=?, address=?, building_area=?, land_area=?, electricity=?,
		water_source=?, bedrooms=?, bathrooms=?, floors=?, garage=?, carport=?,
		certificate=?, year_constructed=?, property_type_id=?
		WHERE id=?
	`

	_, err := r.DB.ExecContext(ctx, query,
		p.Title, p.Description, p.Price, p.Status, p.Province,
		p.Regency, p.District, p.Address, p.BuildingArea,
		p.LandArea, p.Electricity, p.WaterSource,
		p.Bedrooms, p.Bathrooms, p.Floors,
		p.Garage, p.Carport, p.Certificate,
		p.YearConstructed, p.PropertyTypeId, p.ID,
	)

	return err
}

func (r *PropertyRepository) Delete(ctx context.Context, id int) error {
	_, err := r.DB.ExecContext(ctx, `DELETE FROM properties WHERE id = ?`, id)
	return err
}

