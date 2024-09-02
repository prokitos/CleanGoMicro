package tables

import "modules/internal/models"

type Task struct {
	Task_id     int    `json:"id" example:"12" gorm:"unique;primaryKey;autoIncrement"`
	Description string `json:"description" example:"something"`
	DueDate     string `json:"due_date" example:"01.01.2010"`
}

func (instance *Task) RecordCreate(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	return dao.CreateData(instance, db)
}
func (instance *Task) RecordShow(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	err := dao.ShowData(instance, db)
	return err
}
func (instance *Task) RecordDelete(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	return dao.DeleteData(instance, db)
}
func (instance *Task) RecordUpdate(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	return dao.UpdateData(instance, db)
}
func (instance *Task) GetId() int {
	return instance.Task_id
}
