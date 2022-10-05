package ecs

import (
	"errors"
	"fmt"
)

type IEntity interface {
	GetWorld() *IWorld
	GetId() string
	AddComponent(cmp IComponent) error
	HaveComponent(cn string) bool
	GetComponent(id string) (*IComponent, error)
	GetComponents() ([]*IComponent, error)
}

type Entity struct {
	Id         string
	World      *IWorld
	Components []*IComponent
}

func (entity *Entity) GetId() string {
	return entity.Id
}

func CEntity(world *IWorld, id string, components []*IComponent) *Entity {
	return &Entity{
		Id:         id,
		World:      world,
		Components: components,
	}
}

func (entity *Entity) AddComponent(cmp IComponent) error {
	// Check if we already have a component with same id
	var foundId int = -1
	components, err := entity.GetComponents()
	if err != nil {
		return err
	}
	for idx, component := range components {
		componentLocalised := *component
		if componentLocalised.GetId() == cmp.GetId() {
			foundId = idx
		}
	}
	if foundId != -1 {
		return errors.New("Component with same id already exist")
	} else {
		entity.Components = append(entity.Components, &cmp)
		return nil
	}
}

func (entity *Entity) HaveComponent(cn string) bool {
	for _, component := range entity.Components {
		componentLocalised := *component
		if componentLocalised.GetId() == cn {
			return true
		}
	}
	return false
}

func (entity *Entity) GetComponent(id string) (cmp *IComponent, err error) {
	var foundId int = -1

	components, err := entity.GetComponents()
	if err != nil {
		return nil, err
	}
	for idx, component := range components {
		componentLocalised := *component
		if componentLocalised.GetId() == id {
			foundId = idx
		}
	}
	if foundId != -1 {
		cmp = entity.Components[foundId]
	} else {
		err = errors.New(fmt.Sprintf("Component %s not found", id))
	}

	return cmp, err
}

func (entity *Entity) GetComponents() (components []*IComponent, err error) {
	return entity.Components, nil
}

func (entity *Entity) GetWorld() *IWorld {
	return entity.World
}
