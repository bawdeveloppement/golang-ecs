package ecs

import (
	"errors"
	"fmt"
)

type IWorld interface {
	AddEntity(*IEntity) error
	GetEntity(string) *IEntity
	GetEntities() []*IEntity
	RemoveEntity(string) error
	AddSystem(ISystem)
	RemoveSystem(string) error
	GetEntitiesByComponentId(string) []*IEntity
	GetEntitiesWithComponents(...string) []*IEntity
	Update()
}

type World struct {
	Entities []*IEntity
	Systems  []ISystem
}

func (world *World) AddEntity(entity *IEntity) (err error) {
	var found bool = false
	for _, ent := range world.Entities {
		entityLocalised := *entity
		entLocalised := *ent
		if entLocalised.GetId() == entityLocalised.GetId() {
			found = true
		}
	}
	if found {
		err = errors.New("Cannot add entity because an entity with same id already exist")
		return err
	} else {
		world.Entities = append(world.Entities, entity)
	}
	return err
}

func (world *World) GetEntity(id string) (ent *IEntity) {
	for _, entity := range world.Entities {
		entityLocalised := *entity
		if entityLocalised.GetId() == id {
			ent = entity
		}
	}
	return ent
}

func (world *World) GetEntities() (entities []*IEntity) {
	return world.Entities
}

func (world *World) GetEntitiesByComponentId(id string) (entities []*IEntity) {
	for _, entity := range world.Entities {
		entityLocalised := *entity
		components, err := entityLocalised.GetComponents()
		if err != nil {
			fmt.Println(err)
		}
		for _, component := range components {
			cmp := *component
			if cmp.GetId() == id {
				entities = append(entities, entity)
			}
		}
	}
	return entities
}

func (world *World) GetEntitiesWithComponents(v ...string) (entities []*IEntity) {
	for _, entity := range world.Entities {
		entityLocalised := *entity
		var checkeds int = 0
		for _, cmpName := range v {
			if entityLocalised.HaveComponent(cmpName) {
				checkeds = checkeds + 1
			}
		}
		if checkeds == len(v) {
			entities = append(entities, entity)
		}
	}
	return entities
}

func (world *World) RemoveEntity(id string) (err error) {
	var newEntities []*IEntity = []*IEntity{}
	var entityFound bool = false

	for _, entity := range world.Entities {
		localisedEntity := *entity
		if localisedEntity.GetId() != id {
			newEntities = append(newEntities, entity)
		} else {
			entityFound = true
		}
	}

	if !entityFound {
		return errors.New("Cannot delete entity because it doesn't exist")
	}

	world.Entities = newEntities

	return nil
}

func (world *World) AddSystem(sys ISystem) {
	world.Systems = append(world.Systems, sys)
}

func (world *World) RemoveSystem(id string) (err error) {
	var newSystems []ISystem = []ISystem{}
	var systemFound bool = false

	for _, system := range world.Systems {
		if system.GetId() != id {
			newSystems = append(newSystems, system)
		} else {
			systemFound = true
		}
	}

	if !systemFound {
		err = errors.New("Cannot delete entity because it doesn't exist")
		return err
	} else {
		world.Systems = newSystems
	}

	return nil
}

func (world *World) Update() {
	for _, system := range world.Systems {
		system.Update()
	}
}
