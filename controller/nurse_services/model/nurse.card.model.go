package nursemodel

type NurseCardModel struct {
	UserID           string   `db:"user_id" json:"user_id"`
	Avatar           string   `db:"avatar" json:"avatar"`
	FullName         string   `db:"full_name" json:"full_name"`
	CurrentWorkPlace string   `db:"current_workplace" json:"current_workplace"`
	Expertise        string   `db:"expertise" json:"expertise"`
	Techniques       []string `json:"techniques"`
}
