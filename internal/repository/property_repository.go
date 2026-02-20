package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"vp_backend/internal/domain"
)

// PropertyRepository bertanggung jawab untuk
// mengelola seluruh operasi database yang
// berkaitan dengan entitas Property.
type PropertyRepository struct {
	DB *sql.DB
}

// Create menyimpan data properti baru ke database.
//
// Parameter:
// - ctx : context untuk kontrol lifecycle query
// - p   : pointer ke domain.Property
//
// Return:
// - error jika proses insert gagal
func (r *PropertyRepository) Create(
	ctx context.Context,
	p *domain.Property,
) error {

	query := `
	INSERT INTO properties
	(title, description, price, status, province, regency, district, address,
	building_area, land_area, electricity, water_source, bedrooms, bathrooms,
	floors, garage, carport, certificate, year_constructed, sale_type, latitude, longitude, property_type_id, agent_id)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := r.DB.ExecContext(
		ctx,
		query,
		p.Title, p.Description, p.Price, p.Status, p.Province,
		p.Regency, p.District, p.Address, p.BuildingArea,
		p.LandArea, p.Electricity, p.WaterSource, p.Bedrooms,
		p.Bathrooms, p.Floors, p.Garage, p.Carport,
		p.Certificate, p.YearConstructed, p.SaleType, p.Latitude, p.Longitude,
		p.PropertyTypeId, p.AgentId,
	)

	return err
}

// FindAll mengambil seluruh data properti
// tanpa filter.
func (r *PropertyRepository) FindAll(
	ctx context.Context,
) ([]domain.Property, error) {

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
			&p.Certificate, &p.YearConstructed, &p.SaleType, &p.Latitude, &p.Longitude,
			&p.CreatedAt, &p.CoverImageUrl, &p.PropertyTypeId, &p.AgentId,
		); err != nil {
			return nil, err
		}
		properties = append(properties, p)
	}

	return properties, nil
}

// Update memperbarui data properti
// berdasarkan ID.
func (r *PropertyRepository) Update(
	ctx context.Context,
	p *domain.Property,
) error {
	query := `
	UPDATE properties SET
	title=?, description=?, price=?, status=?, province=?, regency=?,
	district=?, address=?, building_area=?, land_area=?, electricity=?,
	water_source=?, bedrooms=?, bathrooms=?, floors=?, garage=?, carport=?,
	certificate=?, year_constructed=?, sale_type=?, latitude=?, longitude=?,
	cover_image_url=?, property_type_id=?
	WHERE id=?
	`

	_, err := r.DB.ExecContext(
		ctx,
		query,
		p.Title, p.Description, p.Price, p.Status, p.Province,
		p.Regency, p.District, p.Address, p.BuildingArea,
		p.LandArea, p.Electricity, p.WaterSource,
		p.Bedrooms, p.Bathrooms, p.Floors,
		p.Garage, p.Carport, p.Certificate, p.YearConstructed,
		p.SaleType, p.Latitude, p.Longitude, p.CoverImageUrl,
		p.PropertyTypeId, p.ID,
	)

	return err
}

// Delete menghapus data properti
// berdasarkan ID.
func (r *PropertyRepository) Delete(
	ctx context.Context,
	id int,
) error {
	res, err := r.DB.ExecContext(ctx, `DELETE FROM properties WHERE id = ?`, id)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return errors.New("properti tidak ditemukan")
	}

	return nil
}

// FindByID mengambil satu data properti
// berdasarkan ID.
//
// Parameter:
// - ctx : context untuk kontrol lifecycle query
// - id  : ID properti
//
// Return:
// - *domain.Property jika ditemukan
// - error jika tidak ditemukan atau query gagal
func (r *PropertyRepository) FindByID(
	ctx context.Context,
	id int,
) (*domain.Property, error) {

	query := `SELECT * FROM properties WHERE id = ?`

	p := domain.Property{}
	err := r.DB.QueryRowContext(ctx, query, id).Scan(
		&p.ID, &p.Title, &p.Description, &p.Price, &p.Status,
		&p.Province, &p.Regency, &p.District, &p.Address,
		&p.BuildingArea, &p.LandArea, &p.Electricity,
		&p.WaterSource, &p.Bedrooms, &p.Bathrooms,
		&p.Floors, &p.Garage, &p.Carport,
		&p.Certificate, &p.YearConstructed, &p.SaleType, &p.Latitude, &p.Longitude,
		&p.CreatedAt, &p.CoverImageUrl, &p.PropertyTypeId, &p.AgentId,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("property not found")
	}

	return &p, err
}

func (r *PropertyRepository) CountData(ctx context.Context, f *domain.PropertyFilters) (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM properties"
	var args []interface{}

	if f != nil {
		whereClause, whereArgs := r.buildWhereClause(f)
		query += whereClause
		args = whereArgs
	}

	err := r.DB.QueryRowContext(ctx, query, args...).Scan(&count)
	return count, err
}

// FindFiltered mengambil daftar properti
// berdasarkan filter, sorting, dan pagination.
//
// Return:
// - slice domain.Property
// - total data sebelum pagination
// - error jika query gagal
func (r *PropertyRepository) FindFiltered(
	ctx context.Context,
	f *domain.PropertyFilters,
) ([]domain.Property, error) {

	whereClause, args := r.buildWhereClause(f)

	query := fmt.Sprintf(
		"SELECT * FROM properties %s %s LIMIT ? OFFSET ?",
		whereClause,
		r.buildOrderClause(f.SortBy),
	)

	queryArgs := append(args, f.Limit, f.Offset)

	rows, err := r.DB.QueryContext(ctx, query, queryArgs...)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
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
			&p.Certificate, &p.YearConstructed, &p.SaleType, &p.Latitude, &p.Longitude,
			&p.CreatedAt, &p.CoverImageUrl, &p.PropertyTypeId, &p.AgentId,
		); err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}
		properties = append(properties, p)
	}

	return properties, nil
}

// buildWhereClause membangun klausa WHERE
// secara dinamis berdasarkan filter.
//
// Return:
// - string WHERE clause
// - slice argument query
func (r *PropertyRepository) buildWhereClause(
	f *domain.PropertyFilters,
) (string, []interface{}) {

	var conditions []string
	var args []interface{}

	addCondition := func(field, operator string, value interface{}) {
		conditions = append(conditions, fmt.Sprintf("%s %s ?", field, operator))
		args = append(args, value)
	}

	// Filter kategori & lokasi
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

	// Filter harga
	if f.MinPrice > 0 {
		addCondition("price", ">=", f.MinPrice)
	}
	if f.MaxPrice > 0 {
		addCondition("price", "<=", f.MaxPrice)
	}

	// Filter luas bangunan
	if f.MinBuildingArea > 0 {
		addCondition("building_area", ">=", f.MinBuildingArea)
	}
	if f.MaxBuildingArea > 0 {
		addCondition("building_area", "<=", f.MaxBuildingArea)
	}

	// Filter luas tanah
	if f.MinLandArea > 0 {
		addCondition("land_area", ">=", f.MinLandArea)
	}
	if f.MaxLandArea > 0 {
		addCondition("land_area", "<=", f.MaxLandArea)
	}

	// Filter keyword
	if f.Keyword != "" {
		searchPattern := "%" + f.Keyword + "%"
		conditions = append(conditions, "(title LIKE ? OR description LIKE ?)")
		args = append(args, searchPattern, searchPattern)
	}

	if len(conditions) == 0 {
		return "", args
	}

	return " WHERE " + strings.Join(conditions, " AND "), args
}

// buildOrderClause menentukan
// urutan sorting data properti.
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
