package techniquerepository

import (
	"fmt"

	techniquemodel "github.com/PhuPhuoc/curanest_exe_be/controller/technique_services/model"
)

func (store *techniqueStore) GetAllTechnique() ([]techniquemodel.TechniqueModel, error) {
	list_tech := []techniquemodel.TechniqueModel{}
	query := `
		select
			id, name, estimated_time, fee
		from
			techniques
	`
	if err := store.db.Select(&list_tech, query); err != nil {
		return nil, fmt.Errorf("cannot create new technique <%w>", err)
	}
	return list_tech, nil
}
