package Subscriber

import (
	"context"
	"database/sql"

	models "github.com/ishvaram/betal-kamailio/models"
	sRepo "github.com/ishvaram/betal-kamailio/repository"
)

// NewSQLSubscriberRepo retunrs implement of subscriber repository interface
func NewSQLSubscriberRepo(Conn *sql.DB) sRepo.SubscriberRepo {
	return &mysqlSubscriberRepo{
		Conn: Conn,
	}
}

type mysqlSubscriberRepo struct {
	Conn *sql.DB
}

func (m *mysqlSubscriberRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Subscriber, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Subscriber, 0)
	for rows.Next() {
		data := new(models.Subscriber)

		err := rows.Scan(
			&data.ID,
			&data.Username,
			&data.Domain,
			&data.Password
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *mysqlSubscriberRepo) Fetch(ctx context.Context, num int64) ([]*models.Subscriber, error) {
	query := "Select id, username, domain From subscriber limit ?"

	return m.fetch(ctx, query, num)
}

func (m *mysqlSubscriberRepo) GetByID(ctx context.Context, id int64) (*models.Subscriber, error) {
	query := "SELECT id, username, domain FROM `subscriber` WHERE id=?"

	rows, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	payload := &models.Subscriber{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, models.ErrNotFound
	}

	return payload, nil
}

func (m *mysqlSubscriberRepo) Create(ctx context.Context, p *models.Subscriber) (int64, error) {
	query := "INSERT subscriber SET username=?, domain=?, password=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(ctx, p.Username, p.Domain, p.Password)
	defer stmt.Close()

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (m *mysqlSubscriberRepo) Update(ctx context.Context, p *models.Subscriber) (*models.Subscriber, error) {
	query := "UPDATE suscriber SET title=?, content=? WHERE id=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.Title,
		p.Content,
		p.ID,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

func (m *mysqlSubscriberRepo) Delete(ctx context.Context, id int64) (bool, error) {
	query := "DELETE FROM subscriber WHERE id=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
