package nursemodel

type NurseModel struct {
	UserID           string `db:"user_id" json:"user_id"`
	CitizenID        string `db:"citizen_id" json:"citizen_id"`
	FullName         string `db:"full_name" json:"full_name"`
	PhoneNumber      string `db:"phone_number" json:"phone_number"`
	CurrentWorkPlace string `db:"current_workplace" json:"current_workplace"`
	Expertise        string `db:"expertise" json:"expertise"`
	Certificate      string `db:"certificate" json:"certificate"`
	EducationLevel   string `db:"education_level" json:"education_level"`
	WorkExperience   string `db:"work_experience" json:"work_experience"`
	Slogan           string `db:"slogan" json:"slogan"`
}

type DetailNurseUserView struct {
	AccountNurseUserView
	NurseDetailUser
	Techniques []TechniqueModel
}

type DetailNurseAdminView struct {
	AccountNurseAdminView
	NurseDetailAdmin
	Techniques []TechniqueModel
}

type TechniqueModel struct {
	ID   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	// EstimatedTime string `db:"estimated_time" json:"estimated_time"`
	// Fee           int64  `db:"fee" json:"fee"`
}

type AccountNurseUserView struct {
	Avatar string `db:"avatar" json:"avatar"`
}

type AccountNurseAdminView struct {
	Email  string `db:"email" json:"email"`
	Name   string `db:"name" json:"user_name"`
	Avatar string `db:"avatar" json:"avatar"`
}

type NurseDetailUser struct {
	UserID           string `db:"user_id" json:"user_id"`
	FullName         string `db:"full_name" json:"full_name"`
	CurrentWorkPlace string `db:"current_workplace" json:"current_workplace"`
	Expertise        string `db:"expertise" json:"expertise"`
	Certificate      string `db:"certificate" json:"certificate"`
	EducationLevel   string `db:"education_level" json:"education_level"`
	WorkExperience   string `db:"work_experience" json:"work_experience"`
	Slogan           string `db:"slogan" json:"slogan"`
}

type NurseDetailAdmin struct {
	UserID           string `db:"user_id" json:"user_id"`
	CitizenID        string `db:"citizen_id" json:"citizen_id"`
	FullName         string `db:"full_name" json:"full_name"`
	PhoneNumber      string `db:"phone_number" json:"phone_number"`
	CurrentWorkPlace string `db:"current_workplace" json:"current_workplace"`
	Expertise        string `db:"expertise" json:"expertise"`
	Certificate      string `db:"certificate" json:"certificate"`
	EducationLevel   string `db:"education_level" json:"education_level"`
	WorkExperience   string `db:"work_experience" json:"work_experience"`
	Slogan           string `db:"slogan" json:"slogan"`
}
