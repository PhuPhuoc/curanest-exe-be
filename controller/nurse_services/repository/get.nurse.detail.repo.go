package nurserepository

import (
	"fmt"

	nursemodel "github.com/PhuPhuoc/curanest_exe_be/controller/nurse_services/model"
)

func (store *nurseStore) GetNurseDetail(nurse_id, role string) (any, error) {
	var query string
	if role == "user" {
		nurse := nursemodel.DetailNurseUserView{}
		query = `
			select
				u.avatar,
				n.user_id, n.full_name, n.current_workplace,
				n.expertise, n.certificate, n.education_level,
				n.work_experience, n.slogan
			from
				nurses n
			join
				users u on n.user_id=u.id
			where n.user_id=?
		`
		if err := store.db.Get(&nurse, query, nurse_id); err != nil {
			return nil, fmt.Errorf("cannot get nurse's detail <%w>", err)
		}
		var err_get_tech error
		nurse.Techniques, err_get_tech = store.getTechniqueOfNurse(nurse_id)
		if err_get_tech != nil {
			return nil, err_get_tech
		}
		return nurse, nil
	} else if role == "admin" {
		nurse := nursemodel.DetailNurseAdminView{}
		query = `
			select
				u.email, u.name, u.avatar,
				n.user_id, n.citizen_id, n.full_name,
				n.phone_number, n.current_workplace, n.expertise,
				n.certificate, n.education_level, n.work_experience, n.slogan
			from
				nurses n
			join
				users u on n.user_id=u.id
			where n.user_id=?
		`
		if err := store.db.Get(&nurse, query, nurse_id); err != nil {
			return nil, fmt.Errorf("cannot get nurse's detail <%w>", err)
		}
		var err_get_tech error
		nurse.Techniques, err_get_tech = store.getTechniqueOfNurse(nurse_id)
		if err_get_tech != nil {
			return nil, err_get_tech
		}
		return nurse, nil
	} else {
		return nil, fmt.Errorf("%v is not a valid role", role)
	}
}

func (store *nurseStore) getTechniqueOfNurse(nurse_id string) ([]nursemodel.TechniqueModel, error) {
	techs := []nursemodel.TechniqueModel{}
	query := `
		select
			t.id, t.name
		from
			techniques t
		inner join
			nurse_technique nt on t.id=nt.technique_id
		where
			nt.nurse_id=?
	`
	if err := store.db.Select(&techs, query, nurse_id); err != nil {
		return nil, fmt.Errorf("cannot get nurse's technique <%w>", err)
	}
	return techs, nil
}
