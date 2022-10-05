package entities

import (
	"errors"
	"fmt"
	"hermannvincent/deliveryservice/components"
	"hermannvincent/deliveryservice/ecs"
	"hermannvincent/deliveryservice/utils"
)

type IPlayerEntity interface {
	ecs.IEntity
	MoveToPoint(utils.Point)
}

type PlayerEntity struct {
	Id         string
	World      *ecs.IWorld
	Components []*ecs.IComponent
}

func CPlayerEntity(world *ecs.IWorld, id string, x int, y int, width int, height int) *PlayerEntity {
	return &PlayerEntity{
		Id:    id,
		World: world,
		Components: []*ecs.IComponent{
			components.PositionComponent(x, y),
			components.DimensionComponent(width, height),
		},
	}
}

//#region IEntity Implementation
func (entity *PlayerEntity) GetId() string {
	return entity.Id
}

func (entity *PlayerEntity) AddComponent(cmp ecs.IComponent) error {
	// Check if we already have a component with same id
	var foundId int = -1
	for idx, component := range entity.Components {
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

func (entity *PlayerEntity) GetComponent(id string) (cmp *ecs.IComponent, err error) {
	var foundId int = -1
	for idx, component := range entity.Components {
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

func (entity *PlayerEntity) GetComponents() (components []*ecs.IComponent, err error) {
	for _, component := range entity.Components {
		components = append(components, component)
	}
	return components, nil
}

func (entity *PlayerEntity) GetWorld() *ecs.IWorld {
	return entity.World
}

//#endregion

//#region IPlayerEntity Implementation
func (entity *PlayerEntity) MoveToPoint(p utils.Point) error {
	positionComponentPointer, err := entity.GetComponent("position")
	if err != nil {
		return err
	}

	positionComponent := *positionComponentPointer
	positionComponent.SetData("x", p.X)
	positionComponent.SetData("y", p.Y)

	return nil
}
