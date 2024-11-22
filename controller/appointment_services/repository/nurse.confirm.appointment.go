package appointmentrepository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func (store *appoinmentStore) ConfirmAppointment(appointment_id, status string) (err error) {
	var tx *sqlx.Tx
	tx, err = store.db.Beginx()
	if err != nil {
		return fmt.Errorf("cannot begin transaction <%w>", err)
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

	// todo: check current status of appointment -> if status is not pending -> return error
	flag_check_status_pending := false
	query_checkStatusPending := `select exists(select 1 from appointments where status=? and id=?)`

	if err := store.db.Get(&flag_check_status_pending, query_checkStatusPending, "pending", appointment_id); err != nil {
		return fmt.Errorf("unable to check current status of appointment <%w>", err)
	}
	if !flag_check_status_pending {
		return fmt.Errorf("this appointment has already confirmed/cancel/completed")
	}

	if status == "confirmed" {
		query_confirm_appoinment := `
			update appointments set status=? where id=?
		`
		if _, err = tx.Exec(query_confirm_appoinment, "confirmed", appointment_id); err != nil {
			return fmt.Errorf("cannot update status of appointment to confirmed <%w>", err)
		}

		query_update_nurse_schedule_status := `
			update work_schedules set status=? where appointment_id=?
		`
		if _, err = tx.Exec(query_update_nurse_schedule_status, "not-available", appointment_id); err != nil {
			return fmt.Errorf("cannot update status of work_schedules to not-available <%w>", err)
		}
	} else if status == "cancel" {
		query_cancel_appoinment := `
			update appointments set status=? where id=?
		`
		if _, err = tx.Exec(query_cancel_appoinment, "cancel", appointment_id); err != nil {
			return fmt.Errorf("cannot update status of appointment to cancel <%w>", err)
		}

		query_update_nurse_schedule_status := `
			update work_schedules set status=?, appointment_id=null where appointment_id=?
		`
		if _, err = tx.Exec(query_update_nurse_schedule_status, "available", appointment_id); err != nil {
			return fmt.Errorf("cannot update status of work_schedules to available <%w>", err)
		}
	}
	return nil
}
