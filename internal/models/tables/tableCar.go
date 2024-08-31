package tables

type Car struct {
	Car_id int    `json:"id" example:"12" gorm:"unique;primaryKey;autoIncrement"`
	Mark   string `json:"mark" example:"lada"`
	Year   string `json:"year" example:"1990"`
}
