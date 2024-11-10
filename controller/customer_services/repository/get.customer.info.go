package customerrepository

import (
	"fmt"

	customermodel "github.com/PhuPhuoc/curanest_exe_be/controller/customer_services/model"
)

func (store *customerStore) GetCustomerInfo(customer_id string) (*customermodel.CustomerInfoModel, error) {
	cus := &customermodel.CustomerInfoModel{}

	query := `
		select
			u.email, u.avatar,
			c.user_id, c.citizen_id, c.full_name,
			c.phone_number, c.dob, c.ward,
			c.district, c.city, c.address,
			u.created_at
		from
			customers c
		join
			users u on u.id=c.user_id
		where
			c.user_id=?
	`

	if err := store.db.Get(cus, query, customer_id); err != nil {
		return nil, fmt.Errorf("cannot get customer info <%w>", err)
	}
	return cus, nil
}
