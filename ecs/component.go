package ecs

type IComponent interface {
	GetId() string
	SetData(name string, v any)
	GetData(name string) any
}

type Component struct {
	Id   string
	Data []interface{}
}
