package nurserepository

import (
	"fmt"

	nursemodel "github.com/PhuPhuoc/curanest_exe_be/controller/nurse_services/model"
)

func (store *nurseStore) GetSchedules(nurse_id, from, to string) (shifts []nursemodel.ShiftCurrent, err error) {
	query := `
		select id, shift_date, shift_from, shift_to, status, appointment_id
		from work_schedules
		where nurse_id=? and shift_date between ? and ? 
	`
	if err_query := store.db.Select(&shifts, query, nurse_id, from, to); err_query != nil {
		err = fmt.Errorf("cannot select schedule from db <%w>", err_query)
	}
	return
}
