package tables

import "modules/internal/models"

type TableBuilder struct {
	Instance models.Table
}

func (instance TableBuilder) UserCreate(login string, password string) TableBuilder {
	var temp models.Table
	temp = &User{Login: login, Password: password}
	instance.Instance = temp
	return instance
}
func (instance TableBuilder) ComputerCreate(price int, ram string, cpu string, gpu string) TableBuilder {
	var temp models.Table
	temp = &Computer{Price: price, Ram: ram, Cpu: cpu, Gpu: gpu}
	instance.Instance = temp
	return instance
}

// var builder models.TableBuilder
// var temp []models.Table
// temp = append(temp, builder.UserCreate("vanya", "123456").Instance)
// temp = append(temp, builder.UserCreate("seg", "wey").Instance)
