package ecs

type IEntity interface {
	GetId() string
	AddComponent(cmp IComponent)
	GetComponent(id string) IComponent
	GetComponents() []IComponent
}

type Entity struct {
	Id         string
	Components []IComponent
}

func (entity *Entity) GetId() string {
	return entity.Id
}

func (entity *Entity) AddComponent(cmp IComponent) {
	entity.Components = append(entity.Components, cmp)
}

func (entity *Entity) GetComponent(id string) (cmp IComponent) {
	for _, component := range entity.Components {
		if component.GetId() == id {
			cmp = component
		}
	}

	return cmp
}

func (entity *Entity) GetComponents() []IComponent {
	return entity.Components
}
