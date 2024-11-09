package userrepository

import (
	"fmt"

	usermodel "github.com/PhuPhuoc/curanest_exe_be/controller/user_services/model"
	"github.com/PhuPhuoc/curanest_exe_be/utils"
	"github.com/google/uuid"
)

func (store *userStore) Register(data *usermodel.RegisterForm) error {
	if data.Password != data.ConfirmPassword {
		return fmt.Errorf("passwords do not match")
	}

	flag_duplicate_email := false
	rawsql_checkEmailExist := `select exists(select 1 from users where email = ?)`

	if err := store.db.Get(&flag_duplicate_email, rawsql_checkEmailExist, data.Email); err != nil {
		return fmt.Errorf("unable to check for duplicate emails: %v", err)
	}
	if flag_duplicate_email {
		return fmt.Errorf("email already exists")
	}

	user_id := uuid.New().String()
	newUser := usermodel.User{
		ID:        user_id,
		Email:     data.Email,
		Password:  data.Password,
		Name:      data.Name,
		Avatar:    "https://cdn.dribbble.com/userupload/13007877/file/original-b46e2284f5d9eabaa14986f1f80d1a62.png?resize=752x",
		Role:      "user",
		CreatedAt: utils.GetCurrentDateTime(),
	}

	query := `
	insert into users (id, email, password, name, avatar, role, created_at)
	values (:id, :email, :password, :name, :avatar, :role, :created_at)
	`
	_, err := store.db.NamedExec(query, newUser)
	if err != nil {
		return fmt.Errorf("failed to create new user: %w", err)
	}
	return nil
}
