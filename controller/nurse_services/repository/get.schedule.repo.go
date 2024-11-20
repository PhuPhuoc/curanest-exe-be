package nurserepository

import (
	"fmt"

	nursemodel "github.com/PhuPhuoc/curanest_exe_be/controller/nurse_services/model"
)

func (store *nurseStore) GetSchedules(nurse_id, from, to string) ([]nursemodel.ShiftCurrent, error) {
	shifts := []nursemodel.ShiftCurrent{}
	query := `
		select id, shift_date, shift_from, shift_to, status, appointment_id
		from work_schedules
		where shift_date >= ? and shift_date <= ? and nurse_id=? 
		order by shift_date, shift_from 
	`
	if err_query := store.db.Select(&shifts, query, from, to, nurse_id); err_query != nil {
		return nil, fmt.Errorf("cannot select schedule from db <%w>", err_query)
	}
	return shifts, nil
}
