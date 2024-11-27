package reviewrepository

import (
	"fmt"

	reviewmodel "github.com/PhuPhuoc/curanest_exe_be/controller/review_services/model"
)

func (store *reviewStore) GetAllReviewOrNurse(nurse_id string) ([]reviewmodel.ReviewModel, error) {
	list_review := []reviewmodel.ReviewModel{}
	query_get := `
		select r.id, r.appointment_id, r.patient_name, r.rate, r.techniques, r.content, r.created_at 
		from reviews r
		left join appointments a on a.id = r.appointment_id
		where a.nurse_id=?
	`
	if err := store.db.Select(&list_review, query_get, nurse_id); err != nil {
		return nil, fmt.Errorf("cannot get list review of nurse <%w>", err)
	}

	return list_review, nil
}
