package tables

type Task struct {
	Task_id     int    `json:"id" example:"12" gorm:"unique;primaryKey;autoIncrement"`
	Description string `json:"description" example:"something"`
	DueDate     string `json:"due_date" example:"01.01.2010"`
}
