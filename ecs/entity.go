package ecs

import (
	"errors"
	"fmt"
)

type IEntity interface {
	GetWorld() *IWorld
	GetId() string
	AddComponent(cmp IComponent) error
	GetComponent(id string) (*IComponent, error)
	GetComponents() ([]*IComponent, error)
}

type Entity struct {
	Id         string
	World      *IWorld
	Components []IComponent
}

func (entity *Entity) GetId() string {
	return entity.Id
}

func CEntity(world *IWorld, id string, components []IComponent) *Entity {
	return &Entity{
		Id:         id,
		World:      world,
		Components: components,
	}
}

func (entity *Entity) AddComponent(cmp IComponent) error {
	// Check if we already have a component with same id
	var foundId int = -1
	for idx, component := range entity.Components {
		if component.GetId() == cmp.GetId() {
			foundId = idx
		}
	}
	if foundId != -1 {
		return errors.New("Component with same id already exist")
	} else {
		entity.Components = append(entity.Components, cmp)
		return nil
	}
}

func (entity *Entity) GetComponent(id string) (cmp *IComponent, err error) {
	var foundId int = -1
	for idx, component := range entity.Components {
		if component.GetId() == id {
			foundId = idx
		}
	}
	if foundId != -1 {
		cmp = &entity.Components[foundId]
	} else {
		err = errors.New(fmt.Sprintf("Component %s not found", id))
	}

	return cmp, err
}

func (entity *Entity) GetComponents() (components []*IComponent, err error) {
	for _, component := range entity.Components {
		components = append(components, &component)
	}
	return components, nil
}

func (entity *Entity) GetWorld() *IWorld {
	return entity.World
}
