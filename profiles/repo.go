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

//Create profile
func (repo *profileRepo) createProfile(profile *Profile) (*Profile, error) {
	_, err := repo.conn.Exec(context.Background(), CREATE_PROFILE, &profile.UID, &profile.FirstName, &profile.LastName, &profile.Email, &profile.Birthdate, &profile.Interests, &profile.Bio)
	if err != nil {
		return nil, err
	}
	_, err = repo.conn.Exec(context.Background(), CREATE_LINKS, &profile.UID, &profile.SocialMedia.LinkedIn, &profile.SocialMedia.Twitter, &profile.SocialMedia.Facebook, &profile.SocialMedia.Github, &profile.SocialMedia.Other)
	if err != nil {
		return nil, err
	}
	_, err = repo.conn.Exec(context.Background(), CREATE_SCHEDULES, &profile.UID, &profile.Availability[0], &profile.Availability[1], &profile.Availability[2], &profile.Availability[3], &profile.Availability[4], &profile.Availability[5], &profile.Availability[6])
	if err != nil {
		return nil, err
	}

	for _, education := range profile.Edu {
		_, err = repo.conn.Exec(context.Background(), CREATE_EDUCATION, &education.School, &education.StartDate, &education.EndDate, &education.Major, &education.City, &education.Country)
		if err != nil {
			return nil, err
		}
	}
	for _, jobs := range profile.Jobs {
		_, err = repo.conn.Exec(context.Background(), CREATE_JOBS, &profile.UID, &jobs.Company, &jobs.StartDate, &jobs.EndDate, &jobs.Position, &jobs.City, &jobs.Country)
		if err != nil {
			return nil, err
		}
	}
	return profile, nil
}

//Read profile
func (repo *profileRepo) readProfile(uid string) (*Profile, error) {
	profile := &Profile{}
	err := repo.conn.QueryRow(context.Background(), SELECT_PROFILE, uid).Scan(&profile.UID, &profile.FirstName, &profile.LastName, &profile.Email, &profile.Birthdate, &profile.Interests, &profile.Bio)
	if err != nil {
		return nil, err
	}
	rows, err := repo.conn.Query(context.Background(), SELECT_EDUCATION, uid)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		education := &Education{}
		err = rows.Scan(&education.School, &education.StartDate, &education.EndDate, &education.Major, &education.City, &education.Country)
		if err != nil {
			return nil, err
		}
		profile.Edu = append(profile.Edu, education)
	}

	for rows.Next() {
		jobs := &Job{}
		err = rows.Scan(&profile.UID, &jobs.Company, &jobs.StartDate, &jobs.EndDate, &jobs.Position, &jobs.City, &jobs.Country)
		if err != nil {
			return nil, err
		}
		profile.Jobs = append(profile.Jobs, jobs)
	}

	err = repo.conn.QueryRow(context.Background(), SELECT_LINKS, uid).Scan(&profile.UID, &profile.SocialMedia.LinkedIn, &profile.SocialMedia.Twitter, &profile.SocialMedia.Facebook, &profile.SocialMedia.Github, &profile.SocialMedia.Other)
	if err != nil {
		return nil, err
	}
	err = repo.conn.QueryRow(context.Background(), SELECT_SCHEDULE, uid).Scan(&profile.UID, &profile.Availability[0], &profile.Availability[1], &profile.Availability[2], &profile.Availability[3], &profile.Availability[4], &profile.Availability[5], &profile.Availability[6])
	if err != nil {
		return nil, err
	}
	return profile, nil
}

//Update profile
func (repo *profileRepo) updateProfile(profile *Profile) (*Profile, error) {
	oldProfile, err := repo.readProfile(profile.UID)
	if err != nil {
		return nil, err
	}

	if profile.UID == "" {
		profile.UID = oldProfile.UID
	}
	if profile.FirstName == "" {
		profile.FirstName = oldProfile.FirstName
	}
	if profile.LastName == "" {
		profile.LastName = oldProfile.LastName
	}
	if profile.Email == "" {
		profile.Email = oldProfile.Email
	}
	if profile.Birthdate == "" {
		profile.Birthdate = oldProfile.Birthdate
	}
	if profile.Interests == nil {
		profile.Interests = oldProfile.Interests
	}
	if profile.Bio == "" {
		profile.Bio = oldProfile.Bio
	}
	if len(profile.Availability) <= 0 {
		profile.Availability = oldProfile.Availability
	}
	if profile.Points == 0 {
		profile.Points = oldProfile.Points
	}
	if profile.Edu == nil {
		profile.Edu = oldProfile.Edu
	}
	if profile.SocialMedia == nil {
		profile.SocialMedia = oldProfile.SocialMedia
	}

	// TODO EMILY!!!!
	_, err = repo.conn.Exec(context.Background(), UPDATE_PROFILE, &profile.FirstName, &profile.LastName, &profile.Email, &profile.Birthdate, &profile.Interests, &profile.Bio)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

//Delete profile
// func (repo *profileRepo) deleteProfile() (bool, error) {
// 	_, err := repo.conn.Exec(context.Background(), Del, &profile.UID)
// 	if err != nil {
// 		return false, err
// 	}

// 	return true, nil
// }

//Read All profiles
func (repo *profileRepo) readAllProfiles() ([]*Profile, error) {
	rows, err := repo.conn.Query(context.Background(), SELECT_ALL_PROFILES)
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

func (repo *profileRepo) deleteJob(profile *Profile) (bool, error) {
	_, err := repo.conn.Exec(context.Background(),DELETE_JOBS, &profile.UID, &profile.Jobs, &profile.)
	if err != nil {
		return false, err
	}

	return true, nil
}
