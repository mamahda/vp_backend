package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

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
		floors, garage, carport, certificate, year_constructed, sale_type, property_type_id, user_id)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := r.DB.ExecContext(ctx, query,
		p.Title, p.Description, p.Price, p.Status, p.Province,
		p.Regency, p.District, p.Address, p.BuildingArea,
		p.LandArea, p.Electricity, p.WaterSource, p.Bedrooms,
		p.Bathrooms, p.Floors, p.Garage, p.Carport,
		p.Certificate, p.YearConstructed, p.SaleType, p.PropertyTypeId, p.UserId,
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
		&p.Certificate, &p.YearConstructed, &p.SaleType,
		&p.CreatedAt, &p.PropertyTypeId, &p.UserId,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("property not found")
	}

	return &p, err
}

func (r *PropertyRepository) FindFiltered(ctx context.Context, f *domain.PropertyFilters) ([]domain.Property, int, error) {
	whereClause, args := r.buildWhereClause(f)

	totalCount, err := r.getCount(ctx, whereClause, args)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count properties: %w", err)
	}

	if totalCount == 0 {
		return []domain.Property{}, 0, nil
	}

	query := fmt.Sprintf("SELECT * FROM properties %s %s LIMIT ? OFFSET ?",
		whereClause, r.buildOrderClause(f.SortBy))

	queryArgs := append(args, f.Limit, f.Offset)

	rows, err := r.DB.QueryContext(ctx, query, queryArgs...)
	if err != nil {
		return nil, 0, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var properties []domain.Property
	for rows.Next() {
		p := domain.Property{}
		err := rows.Scan(
			&p.ID, &p.Title, &p.Description, &p.Price, &p.Status,
			&p.Province, &p.Regency, &p.District, &p.Address,
			&p.BuildingArea, &p.LandArea, &p.Electricity,
			&p.WaterSource, &p.Bedrooms, &p.Bathrooms,
			&p.Floors, &p.Garage, &p.Carport,
			&p.Certificate, &p.YearConstructed, &p.SaleType,
			&p.CreatedAt, &p.PropertyTypeId, &p.UserId,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("scan error: %w", err)
		}
		properties = append(properties, p)
	}

	return properties, totalCount, nil
}

func (r *PropertyRepository) buildWhereClause(f *domain.PropertyFilters) (string, []interface{}) {
	var conditions []string
	var args []interface{}

	// Helper function untuk mempermudah append
	addCondition := func(field string, operator string, value interface{}) {
		conditions = append(conditions, fmt.Sprintf("%s %s ?", field, operator))
		args = append(args, value)
	}

	// --- 1. Filter Kategori & Lokasi ---
	if f.SaleType != "" {
		addCondition("sale_type", "=", f.SaleType)
	}
	if f.PropertyTypeID != 0 {
		addCondition("property_type_id", "=", f.PropertyTypeID)
	}
	if f.Province != "" {
		addCondition("province", "=", f.Province)
	}
	if f.Regency != "" {
		addCondition("regency", "=", f.Regency)
	}

	// --- 2. Filter Harga ---
	if f.MinPrice > 0 {
		addCondition("price", ">=", f.MinPrice)
	}
	if f.MaxPrice > 0 {
		addCondition("price", "<=", f.MaxPrice)
	}

	// --- 3. Filter Luas Bangunan ---
	if f.MinBuildingArea > 0 {
		addCondition("building_area", ">=", f.MinBuildingArea)
	}
	if f.MaxBuildingArea > 0 {
		addCondition("building_area", "<=", f.MaxBuildingArea)
	}

	// --- 4. Filter Luas Tanah ---
	if f.MinLandArea > 0 {
		addCondition("land_area", ">=", f.MinLandArea)
	}
	if f.MaxLandArea > 0 {
		addCondition("land_area", "<=", f.MaxLandArea)
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = " WHERE " + strings.Join(conditions, " AND ")
	}
	return whereClause, args
}

func (r *PropertyRepository) getCount(ctx context.Context, where string, args []interface{}) (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM properties" + where
	err := r.DB.QueryRowContext(ctx, query, args...).Scan(&count)
	return count, err
}

func (r *PropertyRepository) buildOrderClause(sortBy string) string {
	mapping := map[string]string{
		"price_asc":  "ORDER BY price ASC, id ASC",
		"price_desc": "ORDER BY price DESC, id DESC",
		"oldest":     "ORDER BY created_at ASC, id ASC",
	}

	if val, ok := mapping[strings.ToLower(sortBy)]; ok {
		return val
	}
	return "ORDER BY created_at DESC, id DESC"
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
			&p.Certificate, &p.YearConstructed, &p.SaleType,
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
		certificate=?, year_constructed=?, sale_type=?, property_type_id=?
		WHERE id=?
	`

	_, err := r.DB.ExecContext(ctx, query,
		p.Title, p.Description, p.Price, p.Status, p.Province,
		p.Regency, p.District, p.Address, p.BuildingArea,
		p.LandArea, p.Electricity, p.WaterSource,
		p.Bedrooms, p.Bathrooms, p.Floors,
		p.Garage, p.Carport, p.Certificate,
		p.YearConstructed, p.SaleType, p.PropertyTypeId, p.ID,
	)

	return err
}

func (r *PropertyRepository) Delete(ctx context.Context, id int) error {
	_, err := r.DB.ExecContext(ctx, `DELETE FROM properties WHERE id = ?`, id)
	return err
}
