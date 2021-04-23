package main

import (
	"context"
	"init/models"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func newPGconn(ctx context.Context) (*pgxpool.Pool, error) {
	dbURL := "postgresql://postgres:mypassword@db/simpleTable"
	dbpool, err := pgxpool.Connect(ctx, dbURL)
	if err != nil {
		return nil, err
	}
	dbpool.Config().MaxConns = 20
	dbpool.Config().MinConns = 1
	dbpool.Config().HealthCheckPeriod = 1 * time.Minute
	dbpool.Config().MaxConnLifetime = 1 * time.Minute
	dbpool.Config().MaxConnIdleTime = 1 * time.Minute
	dbpool.Config().ConnConfig.ConnectTimeout = 1 * time.Second
	return dbpool, err
}

// Select rows
func getAllUsers(ctx context.Context, db *pgxpool.Pool, uu models.Users) (models.Users, error) {
	sql := "SELECT * FROM users"
	rows, err := db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// rows.Next() итерируется по всем строкам, полученным из базы.
	for rows.Next() {
		var u models.User
		// Scan записывает значения столбцов в свойства структуры
		err = rows.Scan(&u.ID, &u.Name, &u.PcName, &u.Group, &u.PhoneNumber, &u.Cabinet, &u.Discription, &u.BirthDate, &u.CreatedAt, &u.ModifiedAt)
		if err != nil {
			return nil, err
		}
		uu = append(uu, u)
	}

	return uu, err
}

// Insert
func insertUser(ctx context.Context, db *pgxpool.Pool, u *models.User) (*models.User, error) {
	sql := `INSERT INTO users(username, pc_name, user_group, phone_number, cabinet, discription, birthdate, created_at, modified_at)
	  VALUES($1, $2, $3, $4, $5, $6, $7, NOW(), NOW()) returning user_id`
	err := db.QueryRow(ctx, sql, u.Name, u.PcName, u.Group, u.PhoneNumber, u.Cabinet, u.Discription, u.BirthDate).Scan(&u.ID)
	if err != nil {
		return u, err
	}

	return u, err
}

// Update
func updateUser(ctx context.Context, db *pgxpool.Pool, u *models.User) error {
	sql := `UPDATE users SET (username, pc_name, user_group, phone_number, cabinet, discription, birthdate, modified_at)
	  = ($1, $2, $3, $4, $5, $6, $7, NOW()) WHERE user_id= $8`
	_, err := db.Exec(ctx, sql, u.Name, u.PcName, u.Group, u.PhoneNumber, u.Cabinet, u.Discription, u.BirthDate, u.ID)
	if err != nil {
		return err
	}
	return err
}

// delete row
func deleteUserByID(ctx context.Context, db *pgxpool.Pool, u *models.User) error {
	const sql = "Delete FROM users where user_id = $1"
	_, err := db.Exec(ctx, sql, u.ID)
	if err != nil {
		return err
	}
	return err
}
