package nursemodel

type TimeFrameServices struct {
	AppointmentDate  string   `json:"appoinment_date"`
	From             string   `json:"from"`
	To               string   `json:"to"`
	NurseScheduleIDs []string `json:"nurse_schedule_ids"`
}
