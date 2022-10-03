package systems

import (
	"hermannvincent/deliveryservice/ecs"
	"log"
)

type CollisionSystem struct {
	ecs.ISystem

	Id    string
	World ecs.IWorld
}

func (c CollisionSystem) GetId() string {
	return c.Id
}

func (c CollisionSystem) Update() {
	for _, firstEntity := range c.World.GetEntities() {
		for _, secondEntity := range c.World.GetEntities() {
			if firstEntity.GetId() != secondEntity.GetId() {
				var firstX int = firstEntity.GetComponent("position").GetData("x").(int)
				var secondX int = secondEntity.GetComponent("position").GetData("x").(int)
				if firstX > secondX {
					log.Println(firstEntity.GetId() + " has x greater than " + secondEntity.GetId())
				}

			}
		}
	}
}
