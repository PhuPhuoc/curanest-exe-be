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
		// todo: triết khấu 20% vào ví của điều dưỡng
		// todo 1: lấy ra tổng chi phí của appointment
		var appointment_fee int
		query_get_appointment_fee := `select total_fee from appointments where id=?`
		if err := tx.Get(&appointment_fee, query_get_appointment_fee, appointment_id); err != nil {
			return fmt.Errorf("cannot get total_fee of appointment <%w>", err)
		}
		// todo 2: lấy ra 20% tổng chi phí của appointment trừ vào ví của điều dưỡng
		commission_cost := float64(appointment_fee) * 25 / 100
		fmt.Println("comis cost: ", commission_cost)

		// todo 2.1: thêm record vào wallet_transactions
		// todo 2.1.1: lấy ra id của điều dưỡng và lấy ra số tiền hiện tại của điều dưỡng đó trong ví
		var nurse_id string
		query_get_nurse_id := `select nurse_id from appointments where id=?`
		if err := tx.Get(&nurse_id, query_get_nurse_id, appointment_id); err != nil {
			return fmt.Errorf("cannot get total_fee of appointment <%w>", err)
		}

		var current_money float64
		query_get_nurse_money := `select wallet_amount from users where id=?`
		if err := tx.Get(&current_money, query_get_nurse_money, nurse_id); err != nil {
			return fmt.Errorf("cannot get current amount in nurse wallet <%w>", err)
		}

		if current_money <= commission_cost {
			return fmt.Errorf("nurse cannot take this appointment because there is not enough money in the wallet")
		}

		// todo 2.1.2: trừ tiền vào tài khoản của điều dưỡng

		// todo 2.1.3: lưu giao dịch vào wallet_transaction

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
