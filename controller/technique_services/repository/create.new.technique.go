package techniquerepository

import (
	"fmt"

	techniquemodel "github.com/PhuPhuoc/curanest_exe_be/controller/technique_services/model"
	"github.com/google/uuid"
)

func (store *techniqueStore) CreateTechnique(data *techniquemodel.TechniqueCreationModel) error {
	data.ID = uuid.New().String()
	query := `
		insert into
			techniques (id, name, estimated_time, fee)
		values
			(:id, :name, :estimated_time, :fee)
	`
	if _, err := store.db.NamedExec(query, data); err != nil {
		return fmt.Errorf("cannot create new technique <%w>", err)
	}
	return nil
}
