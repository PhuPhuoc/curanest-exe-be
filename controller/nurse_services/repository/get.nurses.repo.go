package nurserepository

import (
	"fmt"
	"strings"

	nursemodel "github.com/PhuPhuoc/curanest_exe_be/controller/nurse_services/model"
)

func (store *nurseStore) GetNurses(filter *nursemodel.NurseFilter) ([]nursemodel.NurseCardModel, error) {
	list_nurses := []nursemodel.NurseCardModel{}

	query := `
		select
			n.user_id, u.avatar, n.full_name, n.phone_number, n.current_workplace, n.expertise
		from
			nurses n
		join
			users u on u.id=n.user_id
		where
			u.deleted_at is null
	`

	var conditions []string
	var args []any

	if filter.FullName != "" {
		conditions = append(conditions, "n.full_name like ?")
		args = append(args, "%"+filter.FullName+"%")
	}
	if filter.PhoneNumber != "" {
		conditions = append(conditions, "u.phone_number like ?")
		args = append(args, "%"+filter.PhoneNumber+"%")
	}

	if len(conditions) > 0 {
		query += " and " + strings.Join(conditions, " and ")
	}

	if err := store.db.Select(&list_nurses, query, args...); err != nil {
		return nil, fmt.Errorf("cannot get nurses <%w>", err)
	}

	query_get_tech := `
		select
			t.name
		from techniques t
			join nurse_technique nt on nt.technique_id=t.id
		where nt.nurse_id=?
	`
	for i := range list_nurses {
		slice_tech := []string{}
		if err := store.db.Select(&slice_tech, query_get_tech, list_nurses[i].UserID); err != nil {
			return nil, fmt.Errorf("cannot get nurses's technique <%w>", err)
		}
		list_nurses[i].Techniques = slice_tech
	}

	return list_nurses, nil
}
