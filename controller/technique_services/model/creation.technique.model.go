package techniquemodel


type TechniqueCreationModel struct {
	ID            string `db:"id" json:"-"`
	Name          string `db:"name" json:"name"`
	EstimatedTime string `db:"estimated_time" json:"estimated_time"`
	Fee           int64  `db:"fee" json:"fee"`
}
