package profiles

type Profile struct {
	UID          string    `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	Birthdate    string       `json:"birthdate"`
	Interests    []string  `json:"interests"`
	Bio          string    `json:"bio"`
	Availability [7]string `json:"availability"`
	Points       int       `json:"points"`
	Edu          []*Education
	Jobs         []*Job
	SocialMedia	*Links
	// Role        []string `json:"roles"`
}

type Education struct {
	School    string `json:"school"`
	StartDate int    `json:"start_date"`
	EndDate  int    `json:"end_date"`
	Major     string `json:"major"`
	Location
}

type Job struct {
	Company   string `json:"school"`
	StartDate int    `json:"start_date"`
	EndDate   int    `json:"end_date"` //0 = currently working
	Position     string `json:"position"`
	Location
}

type Location struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type Links struct {
	LinkedIn string `json:"linkedin"`
	Twitter  string `json:"twitter"`
	Facebook string `json:"facebook"`
	Github   string `json:"github"`
	Other    string `json:"other"`
}

type Pairing struct {
	MentorUID string `json:"mentorUID"`
	MenteeUID string `json:"menteeUID"`
}

type ProfileService interface {
	CreateProfile(*Profile) (*Profile, error)
	ReadProfile(string) (*Profile, error)
	UpdateProfile(*Profile) (*Profile, error)
	DeleteProfile(string) (bool, error)
	ReadAllProfiles() ([]*Profile, error)
}

type ProfileRepo interface {
	createProfile(*Profile) (*Profile, error)
	readProfile(string) (*Profile, error)
	updateProfile(*Profile) (*Profile, error)
	deleteProfile(string) (bool, error)
	readAllProfiles() ([]*Profile, error)
}
