package appointmentrepository

import "fmt"

func (store *appoinmentStore) CompletedAppointment(appointment_id string) error {
	// todo: check current status of appointment -> if status is not confirmed -> return error
	flag_check_status_confirmed := false
	query_checkStatusPending := `select exists(select 1 from appointments where status=? and id=?)`

	if err := store.db.Get(&flag_check_status_confirmed, query_checkStatusPending, "confirmed", appointment_id); err != nil {
		return fmt.Errorf("unable to check current status of appointment <%w>", err)
	}
	if !flag_check_status_confirmed {
		return fmt.Errorf("this appointment is not confirmed or has already been completed")
	}

	query_confirm_appoinment := `
			update appointments set status=? where id=?
		`
	if _, err := store.db.Exec(query_confirm_appoinment, "completed", appointment_id); err != nil {
		return fmt.Errorf("cannot update status of appointment to completed <%w>", err)
	}
	return nil
}
