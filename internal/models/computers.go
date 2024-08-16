package models

type Computer struct {
	Computer_id int    `json:"id" example:"12" gorm:"unique;primaryKey;autoIncrement"`
	Price       int    `json:"price" example:"125000"`
	Ram         string `json:"ram" example:"16gb"`
	Cpu         string `json:"cpu" example:"intel core i7"`
	Gpu         string `json:"gpu" example:"geforce gtx 1080"`
}

func (instance *Computer) CreateInstance(price int, ram string, cpu string, gpu string) {
	instance.Price = price
	instance.Ram = ram
	instance.Cpu = cpu
	instance.Gpu = gpu
}
