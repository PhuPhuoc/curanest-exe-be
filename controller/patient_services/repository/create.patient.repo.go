package patientrepository

import (
	"fmt"
	"time"

	patientmodel "github.com/PhuPhuoc/curanest_exe_be/controller/patient_services/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func (store *patientStore) CreatePatient(customer_id string, data *patientmodel.PatientCreationModel) (err error) {

	err = store.checkDuplicateDataForPatient(data)
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

	patient_id := uuid.New().String()
	data.ID = patient_id
	patient_old, err_cal_old := calculatePatientOld(data.DOB)
	if err_cal_old != nil {
		patient_old = 0
	}
	data.Old = patient_old

	query_patient := `
		insert into patients
			(id, avatar, full_name, old, dob, citizen_id, address, phone_number)
		values
			(:id, :avatar, :full_name, :old, :dob, :citizen_id, :address, :phone_number)
	`
	_, err = tx.NamedExec(query_patient, data)
	if err != nil {
		return fmt.Errorf("cannot insert relation of customer vs patient <%w>", err)
	}

	query_customer_patient := `
		insert into customer_patient
			(user_id, patient_id)
		values
			(?,?)
	`
	_, err = tx.Exec(query_customer_patient, customer_id, patient_id)
	if err != nil {
		return fmt.Errorf("cannot insert relation of customer vs patient <%w>", err)
	}

	query_cate_fate := `
		insert into patient_technique
			(patient_id, technique_id)
		values
			(?,?)
	`
	for _, tid := range data.Techniques {
		if _, err = tx.Exec(query_cate_fate, patient_id, tid); err != nil {
			return fmt.Errorf("cannot map technique to patient <%w>", err)
		}
	}
	return nil
}

func calculatePatientOld(ngaySinhStr string) (int, error) {
	ngaySinh, err := time.Parse("02/01/2006", ngaySinhStr)
	if err != nil {
		return 0, err
	}
	homNay := time.Now()
	tuoi := homNay.Year() - ngaySinh.Year()

	if homNay.YearDay() < ngaySinh.YearDay() {
		tuoi--
	}
	return tuoi, nil
}

func (store *patientStore) checkDuplicateDataForPatient(data *patientmodel.PatientCreationModel) error {
	flag_duplicate_citizenid := false
	rawsql_checkCitizenIDExist := `select exists(select 1 from patients where citizen_id=?)`
	if err := store.db.Get(&flag_duplicate_citizenid, rawsql_checkCitizenIDExist, data.CitizenID); err != nil {
		return fmt.Errorf("unable to check for duplicate citizen id: %v", err)
	}
	if flag_duplicate_citizenid {
		return fmt.Errorf("this citizen id already exists")
	}

	flag_duplicate_phone_number := false
	rawsql_checkPhoneNumberExist := `select exists(select 1 from patients where phone_number=?)`
	if err := store.db.Get(&flag_duplicate_phone_number, rawsql_checkPhoneNumberExist, data.PhoneNumber); err != nil {
		return fmt.Errorf("unable to check for duplicate phone number: %v", err)
	}
	if flag_duplicate_phone_number {
		return fmt.Errorf("this phone number already exists")
	}

	return nil
}
