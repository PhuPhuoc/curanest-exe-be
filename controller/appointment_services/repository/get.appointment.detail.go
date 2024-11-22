package appointmentrepository

import (
	"fmt"

	appointmentmodel "github.com/PhuPhuoc/curanest_exe_be/controller/appointment_services/model"
	"github.com/jmoiron/sqlx"
)

func (store *appoinmentStore) GetAppointmentDetail(appointment_id string) (data *appointmentmodel.DetailAppointment, err error) {
	var tx *sqlx.Tx
	tx, err = store.db.Beginx()
	if err != nil {
		return nil, fmt.Errorf("cannot begin transaction <%w>", err)
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else if commitErr := tx.Commit(); commitErr != nil {
			err = fmt.Errorf("cannot commit transaction <%w>", commitErr)
		}
	}()

	data = &appointmentmodel.DetailAppointment{}

	query_get_appointment := `
		select 
		id, patient_id, nurse_id, appointment_date, time_from_to, status, techniques, total_fee
		from 
		appointments 
		where 
		id=?
	`
	if err = tx.Get(&data.AppointmentModel, query_get_appointment, appointment_id); err != nil {
		return nil, fmt.Errorf("cannot get appointment info <%w>", err)
	}

	query_get_nurse := `
		select n.full_name, n.phone_number, u.avatar
		from 
			nurses n
		join 
			users u on u.id=n.user_id
		where user_id=?
	`
	if err = tx.Get(&data.NurseInformation, query_get_nurse, data.NurseID); err != nil {
		return nil, fmt.Errorf("cannot get nurse info <%w>", err)
	}
	query_get_patient := `
		select
			p.avatar, p.full_name, p.phone_number, p.old, p.dob, p.ward, p.district, p.city, p.address, p.medical_description, p.note_for_nurses
		from
			patients p
		where p.id=?
	`
	if err = tx.Get(&data.PatientInformation, query_get_patient, data.PatientID); err != nil {
		return nil, fmt.Errorf("cannot get patient info <%w>", err)
	}
	return data, nil
}
