package db

import (
	"context"
	"database/sql"
	"log"
	"strconv"

	"github.com/ilovejs/profile/schema"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgres(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{
		db,
	}, nil
}

func (r *PostgresRepository) Close() {
	r.db.Close()
}

func (r *PostgresRepository) CreateProfile(ctx context.Context, p schema.Profile) (int, error) {
	rows := r.db.QueryRow(
		"INSERT INTO profiles (name, gender, dob, postcode, phone_no) VALUES($1, $2, $3, $4, $5) "+
			"RETURNING id;",
		p.Name, p.Gender, p.DoB, p.PostCode, p.PhoneNo)

	err := rows.Scan(&p.ID)
	if err != nil {
		log.Print("CreateProfile: ", err)
	}
	pid, _ := strconv.Atoi(p.ID)
	return pid, err
}

func (r *PostgresRepository) UpdateProfile(ctx context.Context, p schema.Profile, pid int) error {
	_, err := r.db.Exec(
		"UPDATE profiles SET name=$1, gender=$2, dob=$3, postcode=$4, phone_no=$5 WHERE id=$6",
		p.Name, p.Gender, p.DoB, p.PostCode, p.PhoneNo, pid)

	if err != nil {
		log.Print(err)
	}
	return err
}

func (r *PostgresRepository) ListProfiles(
	ctx context.Context,
	skip uint64, take uint64) ([]schema.Profile, error) {
	rows, err := r.db.Query(
		"SELECT * FROM profiles ORDER BY id DESC OFFSET $1 LIMIT $2", skip, take)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var profiles []schema.Profile
	// parse rows
	for rows.Next() {
		p := schema.Profile{}
		if err = rows.Scan(&p.ID, &p.Name, &p.Gender, &p.DoB, &p.PostCode, &p.PhoneNo); err == nil {
			profiles = append(profiles, p)
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return profiles, nil
}
