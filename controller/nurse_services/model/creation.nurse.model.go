package nursemodel

type NurseCreationModel struct {
	NurseAccount
	NurseData
	Techniques []string
}

type NurseAccount struct {
	ID        string `db:"id" json:"-"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"password"`
	Name      string `db:"name" json:"name"`
	Avatar    string `db:"avatar" json:"avatar"`
	Role      string `db:"role" json:"-"`
	CreatedAt string `db:"created_at" json:"-"`
}

type NurseData struct {
	UserID           string `db:"user_id" json:"-"`
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
