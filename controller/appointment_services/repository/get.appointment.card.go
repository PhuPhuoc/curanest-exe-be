package appointmentrepository

import (
	"fmt"

	appointmentmodel "github.com/PhuPhuoc/curanest_exe_be/controller/appointment_services/model"
)

func (store *appoinmentStore) GetAppointmentCard(account_id, role, from, to string) (any, error) {
	if role == "patient" {
		data := []appointmentmodel.AppointmentCardForCustomer{}
		query := `
		select 
			a.id, a.appointment_date, a.time_from_to, a.techniques, a.status, a.total_fee, u.avatar, n.full_name, n.phone_number
		from appointments a
		left join 
			nurses n on n.user_id=a.nurse_id
		left join 
			users u on u.id=a.nurse_id
		where a.patient_id=? and a.appointment_date between ? and ? 
		order by a.appointment_date desc, a.status, a.time_from_to
		`
		if err := store.db.Select(&data, query, account_id, from, to); err != nil {
			return nil, fmt.Errorf("cannot get appointment of the patient <%w>", err)
		}
		return data, nil

	} else if role == "nurse" {
		data := []appointmentmodel.AppointmentCardForNurse{}
		query := `
		select 
			a.id, a.appointment_date, a.time_from_to, a.techniques, a.status, a.total_fee, p.full_name, p.phone_number, p.avatar
		from appointments a
		left join 
			patients p on p.id=a.patient_id
		where a.nurse_id=? and a.appointment_date between ? and ? 
		order by a.appointment_date desc, a.status, a.time_from_to
		`
		if err := store.db.Select(&data, query, account_id, from, to); err != nil {
			return nil, fmt.Errorf("cannot get appointment of the nurse <%w>", err)
		}
		return data, nil
	} else {
		return nil, nil
	}
}
