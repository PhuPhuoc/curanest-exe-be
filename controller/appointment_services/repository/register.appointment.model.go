package appointmentrepository

import (
	"fmt"

	appointmentmodel "github.com/PhuPhuoc/curanest_exe_be/controller/appointment_services/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func (store *appoinmentStore) RegisterAppoinment(data *appointmentmodel.RegisterAppoinment) (err error) {
	if err != nil {
		return err
	}

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

	appoint_id := uuid.New().String()
	data.ID = appoint_id
	data.Status = "pending"

	query_appointment := `
		insert into appointments 
		(id, patient_id, nurse_id, appointment_date, time_from_to, status, techniques, total_fee)
		values
		(:id, :patient_id, :nurse_id, :appointment_date, :time_from_to, :status, :techniques, :total_fee)
	`
	if _, err = tx.NamedExec(query_appointment, data); err != nil {
		return fmt.Errorf("cannot insert new appointment into db <%w>", err)
	}

	if len(data.ListNurseWorkSchedules) == 0 {
		return fmt.Errorf("must have nurse's schedule information to be able to register for an appointment")
	}

	query_update_nurse_schedule := `
		update work_schedules set status=?, appointment_id=? 
		where id in (?)
	`
	var args []interface{}
	query_update_nurse_schedule, args, err = sqlx.In(query_update_nurse_schedule, "pending", appoint_id, data.ListNurseWorkSchedules)
	if err != nil {
		return fmt.Errorf("failed to build query <%w>", err)
	}

	query_update_nurse_schedule = tx.Rebind(query_update_nurse_schedule)

	_, err = tx.Exec(query_update_nurse_schedule, args...)
	if err != nil {
		return fmt.Errorf("failed to execute query <%w>", err)
	}

	return nil
}
