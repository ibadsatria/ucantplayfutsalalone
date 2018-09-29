package repository

import (
	"database/sql"
	"fmt"

	model "github.com/ibadsatria/ucantplayalone/futsalfield"
	"github.com/sirupsen/logrus"
)

type postgreFutsalfieldRepository struct {
	Conn *sql.DB
}

func NewPostgreFutsalfieldRepository(conn *sql.DB) FutsalfieldRepository {
	return &postgreFutsalfieldRepository{conn}
}

func (p *postgreFutsalfieldRepository) fetch(query string, args ...interface{}) ([]*model.Futsalfield, error) {
	rows, err := p.Conn.Query(query, args...)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer rows.Close()

	result := make([]*model.Futsalfield, 0)
	for rows.Next() {
		t := new(model.Futsalfield)

		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.Address,
			&t.HourlyPrice,
			&t.Longitude,
			&t.Latitude,
			&t.CreatedAt,
			&t.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}

func (p *postgreFutsalfieldRepository) GetByHourlyPriceRange(low, high float64) ([]*model.Futsalfield, error) {
	query := `SELECT id, name, address, hourly_price, long, lat, updated_at, created_at 
				FROM futsalfield WHERE hourly_price >= $1 AND hourly_price <= $2`
	return p.fetch(query, low, high)
}

func (p *postgreFutsalfieldRepository) GetByName(key string) (*model.Futsalfield, error) {
	query := `SELECT id, name, address, hourly_price, long, lat, updated_at, created_at
				FROM futsalfield WHERE name LIKE '%' || $1 || '%'`
	list, err := p.fetch(query, key)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	f := &model.Futsalfield{}
	if len(list) > 0 {
		f = list[0]
	} else {
		return nil, model.NOT_FOUND_ERROR
	}

	return f, nil
}

func (p *postgreFutsalfieldRepository) GetByID(id int64) (*model.Futsalfield, error) {
	query := `SELECT id, name, address, hourly_price, long, lat, updated_at, created_at
				FROM futsalfield WHERE id = $1`
	list, err := p.fetch(query, id)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	f := &model.Futsalfield{}
	if len(list) > 0 {
		f = list[0]
	} else {
		return nil, model.NOT_FOUND_ERROR
	}

	return f, nil
}

func (p *postgreFutsalfieldRepository) Update(m *model.Futsalfield) (*model.Futsalfield, error) {
	query := `UPDATE futsalfield SET name = $1, address = $2, hourly_price = $3, long = $4, lat = $5, updated_at = $6, created_at = $7 
		WHERE id = $5`
	stmt, err := p.Conn.Prepare(query)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	res, err := stmt.Exec(m.Name, m.Address, m.HourlyPrice, m.Longitude, m.Latitude, m.UpdatedAt, m.CreatedAt, m.ID)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	if affected != 1 {
		logrus.Error("Weird behaviour. Total affected: ", affected)
		return nil, err
	}

	return m, nil
}

func (p *postgreFutsalfieldRepository) Store(m *model.Futsalfield) (int64, error) {
	query := `INSERT futsalfield SET name = $1, address = $2, hourly_price = $3, long = $4, lat = $5, updated_at = $6, created_at = $7`
	stmt, err := p.Conn.Prepare(query)

	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	logrus.Debug("Created at: ", m.CreatedAt)
	res, err := stmt.Exec(m.Name, m.Address, m.HourlyPrice, m.Longitude, m.Latitude, m.UpdatedAt, m.CreatedAt)
	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	lastInsertedID, _ := res.LastInsertId()
	return lastInsertedID, nil
}

func (p *postgreFutsalfieldRepository) Delete(id int64) (bool, error) {
	query := "DELETE FROM futsal_field WHERE id = $1"

	stmt, err := p.Conn.Prepare(query)
	if err != nil {
		return false, err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	if rowsAffected != 1 {
		err = fmt.Errorf("Weird behaviour. Total affected: %d", rowsAffected)
		logrus.Error(err)
		return false, err
	}

	return true, nil
}
