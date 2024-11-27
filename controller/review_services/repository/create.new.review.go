package reviewrepository

import (
	"fmt"

	reviewmodel "github.com/PhuPhuoc/curanest_exe_be/controller/review_services/model"
	"github.com/PhuPhuoc/curanest_exe_be/utils"
	"github.com/google/uuid"
)

func (store *reviewStore) CreateNewReview(appointment_id string, data *reviewmodel.CreateNewReviewModel) error {
	query := `
		insert into reviews (
			id, appointment_id, patient_name, rate, techniques, content, created_at
		) values (
			:id, :appointment_id, :patient_name, :rate, :techniques, :content, :created_at
		)
	`
	data.ID = uuid.New().String()
	data.AppointmentID = appointment_id
	data.CreatedAt = utils.GetCurrentDateTime()
	if _, err := store.db.NamedExec(query, data); err != nil {
		return fmt.Errorf("cannot insert new reviews <%w>", err)
	}
	return nil
}
