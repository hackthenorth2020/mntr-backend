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

// func (srv *profileService) DeleteProfile(uid string) (bool, error) {
// 	log.Printf("[PROFILE SRV] [DELETE PROFILE] %v", uid)
// 	result, err := srv.repo.deleteProfile(uid)
// 	if err != nil {
// 		return false, err
// 	}
// 	return result, nil
// }

func (srv *profileService) ReadAllProfiles() ([]*Profile, error) {
	log.Printf("[PROFILE SRV] [READ ALL PROFILES]")
	result, err := srv.repo.readAllProfiles()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (srv *profileService) DeleteJob(req *DeleteJobRequest) (bool, error) {
	log.Printf("[PROFILE SRV] [DELETE JOB] %v", req)
	result, err := srv.repo.deleteJob(req)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (srv *profileService) DeleteEdu(req *DeleteEduRequest) (bool, error) {
	log.Printf("[PROFILE SRV] [DELETE JOB] %v", req)
	result, err := srv.repo.deleteEdu(req)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (srv *profileService) FindMentor(UID string) ([]*Profile, error) {
	log.Printf("[PROFILE SRV] [FIND MENTOR] %v", UID)
	//if already bending requset return nil?
	result, err := srv.repo.findMentor(UID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (srv *profileService) RequestMentor(req *MentorRequest) (bool, error) {
	log.Printf("[PROFILE SRV] [REQ MENTOR] %v", req)
	//check if already requests?
	result, err := srv.repo.requestMentor(req)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (srv *profileService) DeleteMentor(req *MentorRequest) (bool, error) {
	log.Printf("[PROFILE SRV] [DELETE MENTOR] %v", req)
	result, err := srv.repo.deleteMentor(req)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (srv *profileService) ViewMentorRequests(UID string) ([]*Profile, error) {
	log.Printf("[PROFILE SRV] [VIEW MENTOR REQS] %v", UID)
	result, err := srv.repo.viewMentorRequests(UID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (srv *profileService) HandleMentorRequest(req *MentorResponse) (bool, error) {
	log.Printf("[PROFILE SRV] [HANDLE MENTOR REQUEST] %v", req)
	result, err := srv.repo.handleMentorRequest(req)
	if err != nil {
		return false, err
	}
	return result, nil
}

// func (srv *profileService) DeleteMentee(req *MentorRequest) (bool, error) {
// 	log.Printf("[PROFILE SRV] [DELETE MENTEE] %v", req)
// 	result, err := srv.repo.deleteMentor(req)
// 	if err != nil {
// 		return false, err
// 	}
// 	return result, nil
// }

func (srv *profileService) GetMessages(req *GetMessageRequest) ([]*Message, error) {
	log.Printf("[PROFILE SRV] [GET MESSAGES] %v", req)
	result, err := srv.repo.getMessages(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (srv *profileService) SendMessage(req *Message) (bool, error) {
	log.Printf("[PROFILE SRV] [GET MESSAGES] %v", req)
	result, err := srv.repo.sendMessage(req)
	if err != nil {
		return false, err
	}
	return result, nil
}
