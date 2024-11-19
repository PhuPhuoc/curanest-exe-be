package nursemodel

type WeeklyWorkSchedule struct {
	WeekFrom string  `json:"week_from" example:"2024-11-17"`
	WeekTo   string  `json:"week_to" example:"2024-11-23"`
	Shifts   []Shift `json:"shifts"`
}

type Shift struct {
	ID        string `db:"id" json:"-"`
	NurseID   string `db:"nurse_id" json:"-"`
	ShiftDate string `db:"shift_date" json:"shift_date" example:"2024-11-20"`
	ShiftFrom string `db:"shift_from" json:"shift_from" example:"00:00:00"`
	ShiftTo   string `db:"shift_to" json:"shift_to" example:"00:00:00"`
	Status    string `db:"status" json:"-"`
}

type ShiftCurrent struct {
	ID             string  `db:"id" json:"id"`
	ShiftDate      string  `db:"shift_date" json:"shift_date"`
	ShiftFrom      string  `db:"shift_from" json:"shift_from"`
	ShiftTo        string  `db:"shift_to" json:"shift_to"`
	Status         string  `db:"status" json:"status"`
	Appointment_id *string `db:"appointment_id" json:"appointment_id"`
}
