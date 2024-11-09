package nurserepository

import (
	"fmt"

	nursemodel "github.com/PhuPhuoc/curanest_exe_be/controller/nurse_services/model"
	"github.com/PhuPhuoc/curanest_exe_be/utils"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func (store *nurseStore) CreateNurse(data *nursemodel.NurseCreationModel) (err error) {
	err = store.checkDuplicateDataForNurse(data)
	if err != nil {
		return err
	}

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

	user_id := uuid.New().String()
	data.ID = user_id
	data.Role = "nurse"
	data.CreatedAt = utils.GetCurrentDateTime()

	query_user := `
	insert into users
		(id, email, password, name, avatar, role, created_at)
	values
		(:id, :email, :password, :name, :avatar, :role, :created_at)
	`
	_, err = tx.NamedExec(query_user, data.NurseAccount)
	if err != nil {
		return fmt.Errorf("failed to create new user <%w>", err)
	}

	data.UserID = user_id
	query_nurse := `
	insert into nurses
		(user_id, citizen_id, full_name, phone_number, current_workplace, expertise, certificate, education_level, work_experience, slogan)
	values
		(:user_id, :citizen_id, :full_name, :phone_number, :current_workplace, :expertise, :certificate, :education_level, :work_experience, :slogan)
	`
	_, err = tx.NamedExec(query_nurse, data.NurseData)
	if err != nil {
		return fmt.Errorf("failed to create new nurse <%w>", err)
	}

	query_cate_fate := `
		insert into nurse_technique
			(nurse_id, technique_id)
		values
			(?,?)
	`
	for _, tid := range data.Techniques {
		if _, err = tx.Exec(query_cate_fate, user_id, tid); err != nil {
			return fmt.Errorf("cannot map technique to nurse <%w>", err)
		}
	}

	return nil
}

func (store *nurseStore) checkDuplicateDataForNurse(data *nursemodel.NurseCreationModel) error {
	flag_duplicate_email := false
	rawsql_checkEmailExist := `select exists(select 1 from users where email=?)`
	if err := store.db.Get(&flag_duplicate_email, rawsql_checkEmailExist, data.Email); err != nil {
		return fmt.Errorf("unable to check for duplicate citizen id: %v", err)
	}
	if flag_duplicate_email {
		return fmt.Errorf("this email already exists")
	}

	flag_duplicate_citizenid := false
	rawsql_checkCitizenIDExist := `select exists(select 1 from nurses where citizen_id=?)`
	if err := store.db.Get(&flag_duplicate_citizenid, rawsql_checkCitizenIDExist, data.CitizenID); err != nil {
		return fmt.Errorf("unable to check for duplicate citizen id: %v", err)
	}
	if flag_duplicate_citizenid {
		return fmt.Errorf("this citizen id already exists")
	}

	flag_duplicate_phone_number := false
	rawsql_checkPhoneNumberExist := `select exists(select 1 from nurses where phone_number=?)`
	if err := store.db.Get(&flag_duplicate_phone_number, rawsql_checkPhoneNumberExist, data.PhoneNumber); err != nil {
		return fmt.Errorf("unable to check for duplicate phone number: %v", err)
	}
	if flag_duplicate_email {
		return fmt.Errorf("this phone number already exists")
	}
	return nil
}
