package nurserepository

import (
	"fmt"

	nursemodel "github.com/PhuPhuoc/curanest_exe_be/controller/nurse_services/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func (store *nurseStore) RegisterWeeklySchedule(nurse_id string, data *nursemodel.WeeklyWorkSchedule) (err error) {
	//* 1. vào db lấy ra toàn bộ lịch làm việc của điều dưỡng (shift) từ from_date đến to_date trong model weeklyWorkSchedule
	//* 2. lọc ra những date-time mới - không tồn tại trong db - để xoá
	//* 3. lọc ra những date-time cũ - trong db - để add thêm vào db
	//* 4. kiểm ra những date đã có lịch hẹn rồi thì không cho huỷ ( xoá lịch )

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

	// todo 1
	query_get_current_shift := `
		select id, shift_date, shift_from, shift_to, status, appointment_id
		from work_schedules
		where shift_date >= ? and shift_date <= ? and nurse_id=? 
		order by shift_date, shift_from 
	`
	current_shifts := []nursemodel.ShiftCurrent{}
	if err := tx.Select(&current_shifts, query_get_current_shift, data.WeekFrom, data.WeekTo, nurse_id); err != nil {
		return fmt.Errorf("cannot get current shifts <%w>", err)
	}

	// todo 2 vs 3
	toAdd, toRemove := findShiftsToAddOrRemove(data.Shifts, current_shifts)

	// todo 2 -> add new shift
	query_add_shift := `
		insert into work_schedules 
		(id, nurse_id, shift_date, shift_from, shift_to, status)
		values
		(:id, :nurse_id, :shift_date, :shift_from, :shift_to, :status)
	`
	for _, shift := range toAdd {
		shift.ID = uuid.New().String()
		shift.Status = "available"
		shift.NurseID = nurse_id
		if _, err = tx.NamedExec(query_add_shift, shift); err != nil {
			return fmt.Errorf("cannot insert new shift <%w>", err)
		}
	}

	// todo 3 -> remove old shift
	for _, shift_id := range toRemove {
		// todo 4 -> check shift has been available or not -> if it is not-available -> cannot remove because this shift has been booked
		if !checkShiftAvalableToRemove(shift_id, current_shifts) {
			continue
		}
		if _, err = tx.Exec("delete from work_schedules where id=?", shift_id); err != nil {
			return fmt.Errorf("cannot remove old shift <%w>", err)
		}
	}
	return nil
}

func checkShiftAvalableToRemove(shift_id string, current_shifts []nursemodel.ShiftCurrent) bool {
	for _, shift := range current_shifts {
		if shift.ID == shift_id {
			if shift.Status == "not-available" || shift.Status == "pending" {
				return false
			}
			break
		}
	}
	return true
}

func findShiftsToAddOrRemove(request_shifts []nursemodel.Shift, current_shifts []nursemodel.ShiftCurrent) (toAdd []nursemodel.Shift, toRemove []string) {
	// todo: map to quick check shift from db
	currentShiftMap := make(map[string]nursemodel.ShiftCurrent)
	for _, cs := range current_shifts {
		key := fmt.Sprintf("%s-%s-%s", cs.ShiftDate, cs.ShiftFrom, cs.ShiftTo)
		currentShiftMap[key] = cs
	}
	// todo: map to quick check shift from reques
	requestShiftMap := make(map[string]nursemodel.Shift)
	for _, cs := range request_shifts {
		key := fmt.Sprintf("%s-%s-%s", cs.ShiftDate, cs.ShiftFrom, cs.ShiftTo)
		requestShiftMap[key] = cs
	}
	// todo: verify shift need to add to db (toAdd)
	for _, rs := range request_shifts {
		key := fmt.Sprintf("%s-%s-%s", rs.ShiftDate, rs.ShiftFrom, rs.ShiftTo)
		if _, exist := currentShiftMap[key]; !exist {
			toAdd = append(toAdd, rs)
		}
	}
	// todo: verify shift need to remove from db (toRemove)
	for _, rs := range current_shifts {
		key := fmt.Sprintf("%s-%s-%s", rs.ShiftDate, rs.ShiftFrom, rs.ShiftTo)
		if _, exist := requestShiftMap[key]; !exist {
			toRemove = append(toRemove, rs.ID)
		}
	}
	return
}
