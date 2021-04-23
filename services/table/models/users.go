package models

import "time"

// User -
type User struct {
	ID          int64      `db:"user_id" json:"user_id"`
	Name        string     `db:"username" json:"username,omitempty"`
	PcName      string     `db:"pc_name" json:"pc_name,omitempty"`
	Group       string     `db:"user_group" json:"user_group,omitempty"`
	PhoneNumber string     `db:"phone_number" json:"phone_number"`
	Cabinet     string     `db:"cabinet" json:"cabinet,omitempty"`
	Discription *string    `db:"discription" json:"discription"`
	BirthDate   *time.Time `db:"birthdate" json:"birthdate"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	ModifiedAt  time.Time  `db:"modified_at" json:"modified_at"`
}

//Users -
//easyjson:json
type Users []User
