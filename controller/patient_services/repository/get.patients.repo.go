package patientrepository

import (
	"fmt"

	patientmodel "github.com/PhuPhuoc/curanest_exe_be/controller/patient_services/model"
)

func (store *patientStore) GetNurseByCustomerID(customer_id string) ([]patientmodel.PatientModelView, error) {
	list_patient := []patientmodel.PatientModelView{}
	query := `
		select
			p.id, p.avatar, p.full_name, p.old, p.dob, p.citizen_id, p.address, p.phone_number
		from
			patients p
		join
			customer_patient cp on cp.patient_id=p.id
		where
			cp.user_id=?
	`

	if err := store.db.Select(&list_patient, query, customer_id); err != nil {
		return nil, fmt.Errorf("cannot get patient <%w>", err)
	}

	query_get_tech := `
		select
			t.id, t.name
		from techniques t
			join patient_technique pt on pt.technique_id=t.id
		where pt.patient_id=?
	`
	for i := range list_patient {
		slice_tech := []patientmodel.TechniqueModel{}
		if err := store.db.Select(&slice_tech, query_get_tech, list_patient[i].ID); err != nil {
			return nil, fmt.Errorf("cannot get patient's technique <%w>", err)
		}
		list_patient[i].Techniques = slice_tech
	}

	return list_patient, nil
}
