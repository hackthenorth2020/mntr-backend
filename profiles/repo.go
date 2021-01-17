package profiles

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

type profileRepo struct {
	// db *pgx.Conn
	conn *pgx.Conn
}

func NewProfileRepo(conn string) ProfileRepo {
	return &profileRepo{
		// db: initDB(conn),
		conn: initDB(conn),
	}
}

func initDB(connStr string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	// defer conn.Close(context.Background())
	// defer log.Printf("Conn closing")
	err = conn.Ping(context.Background())
	if err != nil {
		log.Printf("Unable to ping database: %v\n", err)
		os.Exit(1)
	}
	log.Printf("Connected to database\n")
	return conn
}

//Create item
func (repo *profileRepo) createProfile(profile *Profile) (*Profile, error) {
	_, err := repo.conn.Exec(context.Background(), "INSERT INTO users (uid, firstname, lastname) VALUES ($1, $2, $3)", &profile.UID, &profile.FirstName, &profile.LastName)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

//Read item
func (repo *profileRepo) readProfile(uid string) (*Profile, error) {
	profile := &Profile{}
	row := repo.conn.QueryRow(context.Background(), "SELECT * FROM profiles WHERE uid = $1", uid)
	err := row.Scan(&profile.UID, &profile.FirstName, &profile.LastName)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

//Update item
func (repo *profileRepo) updateProfile(profile *Profile) (*Profile, error) {
	_, err := repo.conn.Exec(context.Background(), "UPDATE profiles SET firstname = $2, lastname = $3 role = $4 WHERE uid = $1", &profile.UID, &profile.FirstName, &profile.LastName)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

//Delete item
func (repo *profileRepo) deleteProfile(uid string) (bool, error) {
	_, err := repo.conn.Exec(context.Background(), "DELETE FROM profiles WHERE uid = $1", &uid)
	if err != nil {
		return false, err
	}

	return true, nil
}

//Read All profiles
func (repo *profileRepo) readAllProfiles() ([]*Profile, error) {
	rows, err := repo.conn.Query(context.Background(), "SELECT * FROM profiles")
	defer rows.Close()

	profiles := make([]*Profile, 0)

	for rows.Next() {
		profile := &Profile{}
		err = rows.Scan(&profile.UID, &profile.FirstName, &profile.LastName)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, profile)
	}

	return profiles, nil
}
