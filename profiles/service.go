package profiles

import "log"

type profileService struct {
	repo ProfileRepo
}

func NewProfileService(conn string) ProfileService {
	return &profileService{
		repo: NewProfileRepo(conn),
	}
}

func (srv *profileService) CreateProfile(profile *Profile) (*Profile, error) {
	log.Printf("[PROFILE SRV] [CREATE PROFILE] %v", profile)
	result, err := srv.repo.createProfile(profile)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (srv *profileService) ReadProfile(uid string) (*Profile, error) {
	log.Printf("[PROFILE SRV] [READ PROFILE] %v", uid)
	result, err := srv.repo.readProfile(uid)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (srv *profileService) UpdateProfile(profile *Profile) (*Profile, error) {
	log.Printf("[PROFILE SRV] [UPDATE PROFILE] %v", profile)
	result, err := srv.repo.updateProfile(profile)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (srv *profileService) DeleteProfile(uid string) (bool, error) {
	log.Printf("[PROFILE SRV] [DELETE PROFILE] %v", uid)
	result, err := srv.repo.deleteProfile(uid)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (srv *profileService) ReadAllProfiles() ([]*Profile, error) {
	log.Printf("[PROFILE SRV] [READ ALL PROFILES]")
	result, err := srv.repo.readAllProfiles()
	if err != nil {
		return nil, err
	}
	return result, nil
}
