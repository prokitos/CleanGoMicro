package models

type User struct {
	User_id  int    `json:"id" example:"12" gorm:"unique;primaryKey;autoIncrement"`
	Login    string `json:"login" example:"admin"`
	Password string `json:"password" example:"123456"`
}

func (instance *User) RecordCreate() {

}
func (instance *User) RecordShow() {

}
func (instance *User) RecordDelete() {

}
func (instance *User) RecordUpdate() {

}
