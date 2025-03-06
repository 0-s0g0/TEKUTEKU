// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package query

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type University string

const (
	UniversityKyutech      University = "kyutech"
	UniversitySciencetokyo University = "science tokyo"
)

func (e *University) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = University(s)
	case string:
		*e = University(s)
	default:
		return fmt.Errorf("unsupported scan type for University: %T", src)
	}
	return nil
}

type NullUniversity struct {
	University University
	Valid      bool // Valid is true if University is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUniversity) Scan(value interface{}) error {
	if value == nil {
		ns.University, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.University.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUniversity) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.University), nil
}

type Message struct {
	MessageID string
	School    int32
	X         int32
	Y         int32
	Message   string
	CreatedAt pgtype.Timestamp
	FloatTime float32
	Likes     int32
}

type User struct {
	UserID         string
	Mail           string
	Name           string
	Belong         University
	HashedPassword string
}
