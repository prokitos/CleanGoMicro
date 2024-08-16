package models

type Response interface {
	GetResponse() Response
	GetError() error
	GoodCreate()
	BadCreate()
	GoodShow([]Table)
	BadShow()
}

type Table interface {
	CreateInstance(interface{})
}
