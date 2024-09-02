package tables

import "modules/internal/models"

type Car struct {
	Car_id int    `json:"id" example:"12" gorm:"unique;primaryKey;autoIncrement"`
	Mark   string `json:"mark" example:"lada"`
	Year   string `json:"year" example:"1990"`
}

func (instance *Car) RecordCreate(db models.BaseDatabase) models.Response {
	return db.CreateData(instance)
}
func (instance *Car) RecordShow(db models.BaseDatabase) models.Response {
	err := db.ShowData(instance)
	return err
}
func (instance *Car) RecordDelete(db models.BaseDatabase) models.Response {
	return db.DeleteData(instance)
}
func (instance *Car) RecordUpdate(db models.BaseDatabase) models.Response {
	return db.UpdateData(instance)
}
