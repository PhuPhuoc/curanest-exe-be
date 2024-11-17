package nursemodel

type WeeklyWorkSchedule struct {
	WeekFrom string  `json:"week-from"`
	WeekTo   string  `json:"week-to"`
	Shifts   []Shift `json:"shifts"`
}

type Shift struct {
	ShiftDate string `json:"shift-date"`
	ShiftFrom string `json:"shift-from"`
	ShiftTo   string `json:"shift-to"`
}
