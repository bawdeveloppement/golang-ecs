package ecs

type IComponent interface {
	GetId() string
	SetData(name string, v any)
	GetData(name string) any
}

type Component struct {
	Id   string
	Data map[string]interface{}
}

func CreateComponent(id string, data map[string]interface{}) *Component {
	return &Component{Id: id, Data: data}
}

func (p *Component) GetId() string {
	return p.Id
}

func (p *Component) GetData(name string) interface{} {
	return p.Data[name]
}

func (p *Component) SetData(name string, v any) {
	p.Data[name] = v
}
