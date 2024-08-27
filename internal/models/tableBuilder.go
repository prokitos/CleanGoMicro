package models

type TableBuilder struct {
	Instance Table
}

func (instance TableBuilder) UserCreate(login string, password string) TableBuilder {
	var temp Table
	temp = &User{Login: login, Password: password}
	instance.Instance = temp
	return instance
}
func (instance TableBuilder) ComputerCreate(price int, ram string, cpu string, gpu string) TableBuilder {
	var temp Table
	temp = &Computer{Price: price, Ram: ram, Cpu: cpu, Gpu: gpu}
	instance.Instance = temp
	return instance
}
