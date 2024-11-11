package customerrepository

import (
	"fmt"

	customermodel "github.com/PhuPhuoc/curanest_exe_be/controller/customer_services/model"
	"github.com/jmoiron/sqlx"
)

func (store *customerStore) CreateCustomer(userid string, data *customermodel.CustomerModel) (err error) {
	err = store.checkDuplicateDataForCustomer(data)
	if err != nil {
		return err
	}

	var tx *sqlx.Tx
	tx, err = store.db.Beginx()
	if err != nil {
		return fmt.Errorf("cannot begin transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else if commitErr := tx.Commit(); commitErr != nil {
			err = fmt.Errorf("cannot commit transaction: %w", commitErr)
		}
	}()

	data.UserID = userid
	query := `
	insert into
		customers (user_id, citizen_id, full_name, phone_number, dob, ward, district, city, address)
	values
		(:user_id, :citizen_id, :full_name, :phone_number, :dob, :ward, :district, :city, :address)
	`
	_, err = tx.NamedExec(query, data)
	if err != nil {
		return fmt.Errorf("failed to create new customer: %w", err)
	}

	query_update_user := `
		update
			users
		set
			role=?
		where
			id=?
	`
	if _, err = tx.Exec(query_update_user, "customer", userid); err != nil {
		return fmt.Errorf("failed to create new customer: %w", err)
	}
	return nil
}

func (store *customerStore) checkDuplicateDataForCustomer(data *customermodel.CustomerModel) error {
	flag_duplicate_citizenid := false
	rawsql_checkCitizenIDExist := `select exists(select 1 from customers where citizen_id=?)`
	if err := store.db.Get(&flag_duplicate_citizenid, rawsql_checkCitizenIDExist, data.CitizenID); err != nil {
		return fmt.Errorf("unable to check for duplicate citizen id: %v", err)
	}
	if flag_duplicate_citizenid {
		return fmt.Errorf("this citizen id already exists")
	}

	flag_duplicate_email := false
	rawsql_checkEmailExist := `select exists(select 1 from customers where phone_number=?)`
	if err := store.db.Get(&flag_duplicate_email, rawsql_checkEmailExist, data.PhoneNumber); err != nil {
		return fmt.Errorf("unable to check for duplicate phone number: %v", err)
	}
	if flag_duplicate_email {
		return fmt.Errorf("this phone number already exists")
	}
	return nil
}
